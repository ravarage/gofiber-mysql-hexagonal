package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Database struct {
	*sql.DB
	*redis.Client
	context.Context
}

// Database settings
type DatabaseConfig struct {
	Driver   string
	Host     string
	Username string
	Password string
	Port     int
	Database string
}

func New(config *DatabaseConfig) (*Database, error) {
	var err error
	// Use DSN string to open
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", config.Username, config.Password, config.Database))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &Database{db, rdb, ctx}, err
}
