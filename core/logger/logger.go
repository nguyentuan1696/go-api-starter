package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"
	"go-api-starter/core/constants"
)

type LogLevel string

const (
	LogLevelDebug LogLevel = "DEBUG"
	LogLevelInfo  LogLevel = "INFO"
	LogLevelWarn  LogLevel = "WARN"
	LogLevelError LogLevel = "ERROR"
)

type LogConfig struct {
	Level         LogLevel
	JSONFormat    bool
	DailyRotation bool // Bật tính năng phân chia log theo ngày
	EnableFile    bool // Bật/tắt ghi log ra file
}

type Logger struct {
	*slog.Logger
	config     LogConfig
	currentDay string
	logFile    *os.File
	level      slog.Level // Lưu trữ level để sử dụng lại
}

var defaultLogger *Logger

func NewLogger(config LogConfig) (*Logger, error) {
	var level slog.Level
	switch config.Level {
	case LogLevelDebug:
		level = slog.LevelDebug
	case LogLevelInfo:
		level = slog.LevelInfo
	case LogLevelWarn:
		level = slog.LevelWarn
	case LogLevelError:
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{
					Key:   "timestamp",
					Value: a.Value,
				}
			}
			return a
		},
	}

	logger := &Logger{
		config: config,
		level:  level, // Lưu trữ level
	}

	var writer io.Writer

	// Mặc định ghi vào file, chỉ ghi stdout khi EnableFile = false
	if config.EnableFile {
		// Tạo thư mục logs nếu chưa tồn tại
		logDir := filepath.Join("logs")
		if err := os.MkdirAll(logDir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create logs directory: %w", err)
		}

		// Dọn dẹp log cũ
		if config.DailyRotation {
			if err := logger.cleanupOldLogs(logDir); err != nil {
				// Log lỗi nhưng không dừng ứng dụng
				fmt.Printf("Warning: failed to cleanup old logs: %v\n", err)
			}
		}

		// Mở file log
		logPath := logger.getLogFilePath(logDir)
		file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("Warning: cannot open log file %s, fallback to stdout: %v\n", logPath, err)
			writer = os.Stdout
		}

		logger.logFile = file
		logger.currentDay = time.Now().Format("2006-01-02")
		writer = file
	} else {
		// Chỉ ghi ra stdout khi EnableFile = false
		writer = os.Stdout
	}

	var handler slog.Handler
	if config.JSONFormat {
		handler = slog.NewJSONHandler(writer, opts)
	} else {
		handler = slog.NewTextHandler(writer, opts)
	}

	logger.Logger = slog.New(handler)
	return logger, nil
}

// getLogFilePath tạo đường dẫn file log với ngày tháng
func (l *Logger) getLogFilePath(logDir string) string {
	if l.config.DailyRotation {
		today := time.Now().Format("2006-01-02")
		filename := constants.LogFileName

		// Tách tên file và extension
		ext := filepath.Ext(filename)
		name := strings.TrimSuffix(filename, ext)

		// Tạo tên file với ngày tháng
		dailyFilename := fmt.Sprintf("%s_%s%s", name, today, ext)
		return filepath.Join(logDir, dailyFilename)
	}
	return filepath.Join(logDir, constants.LogFileName)
}

// cleanupOldLogs xóa các file log cũ hơn constants.LogRetentionDays
func (l *Logger) cleanupOldLogs(logDir string) error {
	cutoffDate := time.Now().AddDate(0, 0, -constants.LogRetentionDays)

	entries, err := os.ReadDir(logDir)
	if err != nil {
		return fmt.Errorf("failed to read logs directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filename := entry.Name()

		// Kiểm tra xem file có phải là log file với format ngày tháng không
		if l.isLogFileWithDate(filename) {
			fileDate, err := l.extractDateFromFilename(filename)
			if err != nil {
				continue // Bỏ qua file không đúng format
			}

			if fileDate.Before(cutoffDate) {
				logPath := filepath.Join(logDir, filename)
				if err := os.Remove(logPath); err != nil {
					fmt.Printf("Warning: failed to remove old log file %s: %v\n", filename, err)
				} else {
					fmt.Printf("Removed old log file: %s\n", filename)
				}
			}
		}
	}

	return nil
}

// isLogFileWithDate kiểm tra xem file có phải là log file với ngày tháng không
func (l *Logger) isLogFileWithDate(filename string) bool {
	baseName := strings.TrimSuffix(constants.LogFileName, filepath.Ext(constants.LogFileName))
	return strings.Contains(filename, baseName+"_") && strings.Contains(filename, "-")
}

// extractDateFromFilename trích xuất ngày từ tên file
func (l *Logger) extractDateFromFilename(filename string) (time.Time, error) {
	// Tìm pattern YYYY-MM-DD trong tên file
	parts := strings.Split(filename, "_")
	if len(parts) < 2 {
		return time.Time{}, fmt.Errorf("invalid filename format")
	}

	// Lấy phần cuối cùng và loại bỏ extension
	datePart := strings.TrimSuffix(parts[len(parts)-1], filepath.Ext(filename))

	return time.Parse("2006-01-02", datePart)
}

// rotateLogIfNeeded kiểm tra và xoay log nếu cần thiết
func (l *Logger) rotateLogIfNeeded() error {
	if !l.config.DailyRotation || l.logFile == nil {
		return nil
	}

	today := time.Now().Format("2006-01-02")
	if l.currentDay == today {
		return nil // Không cần xoay log
	}

	// Đóng file log hiện tại
	if err := l.logFile.Close(); err != nil {
		return fmt.Errorf("failed to close current log file: %w", err)
	}

	// Mở file log mới
	logDir := filepath.Join("logs")
	logPath := l.getLogFilePath(logDir)
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to open new log file: %w", err)
	}

	l.logFile = file
	l.currentDay = today

	// Cập nhật handler với file mới - sử dụng level đã lưu trữ
	opts := &slog.HandlerOptions{
		Level: l.level, // Sử dụng level đã lưu trữ thay vì gọi Enabled()
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{
					Key:   "timestamp",
					Value: a.Value,
				}
			}
			return a
		},
	}

	var handler slog.Handler
	if l.config.JSONFormat {
		handler = slog.NewJSONHandler(file, opts)
	} else {
		handler = slog.NewTextHandler(file, opts)
	}

	l.Logger = slog.New(handler)

	// Dọn dẹp log cũ
	go func() {
		if err := l.cleanupOldLogs(logDir); err != nil {
			fmt.Printf("Warning: failed to cleanup old logs: %v\n", err)
		}
	}()

	return nil
}

// Initialize the default logger
func Init(config LogConfig) error {
	logger, err := NewLogger(config)
	if err != nil {
		return err
	}
	defaultLogger = logger
	return nil
}

// Helper methods for different log levels với log rotation
func (l *Logger) Debug(msg string, args ...any) {
	l.rotateLogIfNeeded()
	l.Logger.Debug(msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.rotateLogIfNeeded()
	l.Logger.Info(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.rotateLogIfNeeded()
	l.Logger.Warn(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.rotateLogIfNeeded()
	l.Logger.Error(msg, args...)
}

func (l *Logger) With(args ...any) *Logger {
	return &Logger{
		Logger:     l.Logger.With(args...),
		config:     l.config,
		currentDay: l.currentDay,
		logFile:    l.logFile,
		level:      l.level, // Sao chép level
	}
}

// Close đóng logger và file log
func (l *Logger) Close() error {
	if l.logFile != nil {
		return l.logFile.Close()
	}
	return nil
}

// Global logger functions with safe fallback
func Error(msg string, args ...any) {
	if defaultLogger == nil {
		fmt.Printf("ERROR: %s %v\n", msg, args)
		return
	}
	defaultLogger.Error(msg, args...)
}

func Debug(msg string, args ...any) {
	if defaultLogger == nil {
		return
	}
	defaultLogger.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	if defaultLogger == nil {
		fmt.Printf("INFO: %s %v\n", msg, args)
		return
	}
	defaultLogger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	if defaultLogger == nil {
		fmt.Printf("WARN: %s %v\n", msg, args)
		return
	}
	defaultLogger.Warn(msg, args...)
}

func With(args ...any) *Logger {
	if defaultLogger == nil {
		tempLogger, _ := NewLogger(LogConfig{
			Level: LogLevelInfo,
		})
		return tempLogger.With(args...)
	}
	return defaultLogger.With(args...)
}

// Close đóng default logger
func Close() error {
	if defaultLogger != nil {
		return defaultLogger.Close()
	}
	return nil
}
