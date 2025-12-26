package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/google/uuid"
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

var schema = `
	CREATE TABLE IF NOT EXISTS user_data(
	id UUID PRIMARY KEY,
	name VARCHAR NOT NULL,
	login VARCHAR NOT NULL,
	PASSWORD VARCHAR NOT NULL
	);

	CREATE TABLE IF NOT EXISTS movie(
	id UUID PRIMARY KEY,
	title VARCHAR NOT NULL,
	amount_marks INTEGER NOT NULL DEFAULT 0,
	total_mark INTEGER NOT NULL DEFAULT 0,
	rating REAL GENERATED ALWAYS AS (total_mark::REAL / NULLIF(amount_marks, 0)) STORED
	);

	CREATE TABLE IF NOT EXISTS favorite_movie(
	id UUID PRIMARY KEY,
	user_id UUID REFERENCES user_data,
	movie_id UUID REFERENCES movie
	);
	`

type UserData struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
}
type Movie struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	AmountMarks int       `json:"amount_marks"`
	TotalMark   int       `json:"total_mark"`
	Rating      float32   `json:"rating"`
}

type FavoriteMovie struct {
	ID      uuid.UUID `json:"id"`
	UserID  uuid.UUID `json:"user_id"`
	MovieID uuid.UUID `json:"movie_id"`
}

// WARN: need close() db in outer function
func InitDB(c Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", c.configString())
	if err != nil {
		return nil, fmt.Errorf("Can't open db: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Can't ping db: %w", err)
	}

	// TODO: parametrize parameters?
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(2 * time.Hour)
	db.SetConnMaxIdleTime(10 * time.Minute)

	_, err = db.Exec(schema)
	if err != nil {
		return nil, fmt.Errorf("Can't create non exist tables: %w", err)
	}
	return db, nil
}
