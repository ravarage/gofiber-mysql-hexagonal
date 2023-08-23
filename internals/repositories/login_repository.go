package repositories

import (
	"database/sql"
	"github.com/redis/go-redis/v9"
)

type Database struct {
	*sql.DB
	*redis.Client
}
