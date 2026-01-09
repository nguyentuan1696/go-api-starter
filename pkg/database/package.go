package database

import "github.com/samber/do/v2"

var Package = do.Package(
	do.Lazy(NewPostgresql),
	do.Lazy(NewRedis),
)
