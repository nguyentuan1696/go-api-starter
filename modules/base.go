package modules

import (
	"go-api-starter/modules/auth"

	"github.com/samber/do/v2"
)

var BasePackage = do.Package(
	auth.Package,
)
