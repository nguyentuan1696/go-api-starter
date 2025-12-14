package di

import (
	"github.com/samber/do/v2"
	"go-api-starter/core/config"
)

var BasePackage = do.Package(
	do.Lazy(config.NewConfig),
)
