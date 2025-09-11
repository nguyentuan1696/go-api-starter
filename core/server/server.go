package server

import (
	"context"
	"flag"
	"fmt"
	"go-api-starter/core/cache"
	"go-api-starter/core/config"
	"go-api-starter/core/database"
	"go-api-starter/core/logger"
	"go-api-starter/core/middleware"
	storageClient "go-api-starter/core/storage"
	"go-api-starter/core/utils"
	"go-api-starter/modules/auth"
	"go-api-starter/modules/product"
	"go-api-starter/modules/storage"
	"go-api-starter/workers"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

type Server struct {
	echo  *echo.Echo
	addr  string
	cache *cache.Cache
	db    database.Database
}

func initEnvironment() (config.Environment, error) {
	env := flag.String("env", "dev", "Environment (dev/prod)")
	flag.Parse()

	switch *env {
	case "dev":
		return config.DevEnvironment, nil
	case "prod":
		return config.ProdEnvironment, nil
	default:
		return "", fmt.Errorf("invalid environment. Use 'dev' or 'prod'")
	}
}

func initServer() (*Server, error) {
	environment, err := initEnvironment()
	if err != nil {
		return nil, err
	}

	if errInitConfig := config.Init(environment); errInitConfig != nil {
		return nil, fmt.Errorf("failed to initialize config: %w", errInitConfig)
	}

	// Get config safely
	cfg, isInitialized := config.GetSafe()
	if !isInitialized {
		return nil, fmt.Errorf("config was not properly initialized")
	}

	// Initialize logger first, before validation
	if errInitLogger := logger.Init(logger.LogConfig{
		Level:         logger.LogLevelDebug,
		EnableFile:    true,
		JSONFormat:    true, // Thêm dòng này để log dạng JSON
		DailyRotation: true, // Có thể thêm daily rotation nếu muốn
	}); errInitLogger != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", errInitLogger)
	}

	// Validate configuration after logger is initialized
	if err = cfg.Validate(); err != nil {
		logger.Error("Configuration validation failed", "error", err)
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	// Initialize database
	db, err := database.InitDB(database.DatabaseConfig{
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		DBName:   cfg.Database.DBName,
	})
	if err != nil {
		logger.Error("Failed to initialize database", "error", err)
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	// Initialize Redis cache
	redisCache := cache.NewCache(
		cfg.Redis.Address,
		cfg.Redis.Password,
		cfg.Redis.DB,
	)

	// Initialize Email Config
	emailConfig := utils.EmailConfig{
		Host:     cfg.SMTP.Host,
		Port:     cfg.SMTP.Port,
		Username: cfg.SMTP.Username,
		Password: cfg.SMTP.Password,
		From:     cfg.SMTP.From,
		FromName: cfg.SMTP.FromName,
	}
	utils.InitEmailConfig(emailConfig)

	// Initialize R2 client
	r2Client, err := storageClient.NewS3Client(cfg)
	if err != nil {
		logger.Error("Failed to initialize R2 client", "error", err)
		return nil, fmt.Errorf("failed to initialize R2 client: %w", err)
	}

	logger.Info("Server initializing",
		"environment", environment,
		"host", cfg.Server.Host,
		"port", cfg.Server.Port,
		"database_host", cfg.Database.Host,
		"database_port", cfg.Database.Port,
		"database_name", cfg.Database.DBName,
		"redis_address", cfg.Redis.Address,
		"redis_db", cfg.Redis.DB,
		"smtp_host", cfg.SMTP.Host,
		"smtp_port", cfg.SMTP.Port,
	)

	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerMiddleware())
	e.Use(middleware.CORSMiddleware())

	// Initialize modules
	product.Init(e, db, *redisCache)
	storage.Init(e, db, r2Client, *redisCache)
	auth.Init(e, db, *redisCache)

	// Initialize Asynq worker server
	workers.NewServer()

	return &Server{
		echo:  e,
		addr:  fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		cache: redisCache,
		db:    db,
	}, nil
}

func Run() error {
	srv, err := initServer()
	if err != nil {
		// Use fmt.Printf instead of logger since logger might not be initialized
		fmt.Printf("Failed to initialize server: %v\n", err)
		return err
	}
	return srv.start()
}

func (s *Server) start() error {
	logger.Info("Starting HTTP server", "address", s.addr)

	go func() {
		if err := s.echo.Start(s.addr); err != nil {
			logger.Info("Shutting down server", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.echo.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown server gracefully: %w", err)
	}

	// Close Redis connection
	if err := s.cache.Close(); err != nil {
		logger.Error("Failed to close Redis connection", "error", err)
	}

	logger.Info("Server shutdown complete")
	return nil
}
