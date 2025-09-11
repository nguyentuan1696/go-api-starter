package workers

import (
	"context"
	"fmt"
	"go-api-starter/core/config"
	"go-api-starter/core/constants"
	"go-api-starter/core/logger"
	"sync"
	"time"

	"github.com/hibiken/asynq"
)

// Global client instance
var (
	clientInstance *asynq.Client
	clientOnce     sync.Once
)

// NewServer creates and returns a new asynq server instance
func NewServer() *asynq.Server {
	cfg := config.Get()
	redisOpt := asynq.RedisClientOpt{
		Addr:     cfg.Redis.Address,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	}

	server := asynq.NewServer(redisOpt, asynq.Config{
		Concurrency: 10,
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
	})

	go func() {
		if err := server.Run(asynq.HandlerFunc(ProcessTask)); err != nil {
			logger.Error("NewServer: Asynq server failed to run", "error", err)
		}
	}()

	logger.Info("NewServer: Asynq server initialized successfully", "redis_addr", cfg.Redis.Address, "concurrency", 10)
	return server
}

// GetClient returns a global asynq client instance (singleton)
func GetClient() *asynq.Client {
	clientOnce.Do(func() {
		cfg := config.Get()
		redisOpt := asynq.RedisClientOpt{
			Addr:     cfg.Redis.Address,
			Password: cfg.Redis.Password,
			DB:       cfg.Redis.DB,
		}
		clientInstance = asynq.NewClient(redisOpt)
	})
	return clientInstance
}

// CloseClient closes the global client instance
func CloseClient() error {
	if clientInstance != nil {
		return clientInstance.Close()
	}
	return nil
}

// ProcessTask is the handler function for Asynq tasks
// It processes tasks based on their type and payload
func ProcessTask(ctx context.Context, task *asynq.Task) error {
	logger.Info("ProcessTask:Handling task", "type", task.Type(), "payload", string(task.Payload()))

	switch task.Type() {
	case constants.TopicQueueEmailDelivery:
		if err := SendEmail(ctx, task.Payload()); err != nil {
			logger.Error("ProcessTask:SendEmail failed", "err", err)
			return fmt.Errorf("send email failed: %w", err)
		}
	}

	return nil
}

// Enqueue enqueues a task using the global client instance
// This is a convenience function for services to easily enqueue tasks
func Enqueue(typeName string, payload []byte, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	client := GetClient()
	task := asynq.NewTask(typeName, payload)
	return client.Enqueue(task, opts...)
}

// EnqueueAt enqueues a task to be processed at the specified time
// This is useful for scheduling tasks to run at a specific time
func EnqueueAt(typeName string, payload []byte, processAt time.Time, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	client := GetClient()
	task := asynq.NewTask(typeName, payload)
	return client.Enqueue(task, asynq.ProcessAt(processAt))
}

// EnqueueIn enqueues a task to be processed after the specified delay
// This is useful for scheduling tasks to run after a certain duration
func EnqueueIn(typeName string, payload []byte, delay time.Duration, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	client := GetClient()
	task := asynq.NewTask(typeName, payload)
	return client.Enqueue(task, asynq.ProcessIn(delay))
}

// ScheduleTaskForSpecificDate schedules a task to run at a specific time
// Parameters:
//   - taskType: type of task to schedule (e.g., constants.TopicQueueEmailDelivery)
//   - payload: task payload data
//   - targetTime: when to execute the task (e.g., targetTime := time.Date(2025, 6, 11, 9, 0, 0, 0, time.UTC))
func ScheduleTaskForSpecificDate(taskType string, payload []byte, targetTime time.Time) (*asynq.TaskInfo, error) {
	// Chuyển đổi thời gian sang UTC
	utcTime := targetTime.UTC()
	taskInfo, err := EnqueueAt(taskType, payload, utcTime)
	if err != nil {
		logger.Error("ScheduleTaskForSpecificDate: Failed to schedule task", "error", err, "task_type", taskType, "target_time", targetTime)
		return nil, fmt.Errorf("failed to schedule task: %w", err)
	}

	logger.Info("ScheduleTaskForSpecificDate: Task scheduled successfully", "task_id", taskInfo.ID, "task_type", taskType, "original_time", targetTime.Format("2006-01-02 15:04:05 MST"), "utc_time", utcTime.Format("2006-01-02 15:04:05 MST"), "queue", taskInfo.Queue)
	return taskInfo, nil
}
