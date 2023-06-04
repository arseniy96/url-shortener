package database

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"time"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase(connectionData string) (*Database, error) {
	//connectionData = "host=localhost user=shortener password= dbname=shortener sslmode=disable"
	db, err := sql.Open("pgx", connectionData)
	if err != nil {
		return nil, err
	}
	database := &Database{
		DB: db,
	}

	return database, nil
}

func (db *Database) HealthCheck() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := db.DB.PingContext(ctx); err != nil {
		return err
	}

	return nil
}

func (db *Database) Close() error {
	return db.DB.Close()
}
