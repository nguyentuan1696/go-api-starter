package pkg

import (
	"go-api-starter/pkg/cache"
	"go-api-starter/pkg/cli"
	"go-api-starter/pkg/config"
	"go-api-starter/pkg/database"
	"go-api-starter/pkg/logger"

	"github.com/samber/do/v2"
)

var BasePackage = do.Package(
	do.Lazy(config.NewConfig),
	do.Lazy(cli.NewCLI),
	do.Lazy(logger.NewLogger),
	do.Lazy(database.NewPostgresql),
	do.Lazy(cache.NewRedis),
)
