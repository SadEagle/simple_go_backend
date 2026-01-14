package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  string
}

func (c Config) configString() string {
	return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s", c.Host, c.Port, c.Database, c.User, c.Password, c.SSLMode)
}

// WARN: need close() db in outer function
func InitDB(c Config) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(context.Background(), c.configString())
	if err != nil {
		log.Fatalln("Unable to create connection pool: %w", err)
	}
	return dbPool, nil
}
