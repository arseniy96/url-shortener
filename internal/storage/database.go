package storage

import (
	"context"
	"database/sql"
	"github.com/arseniy96/url-shortener/internal/logger"
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
	logger.Log.Info("Database connection was created")

	return database, nil
}

func (db *Database) FindRecord(ctx context.Context, value string) (Record, error) {
	row := db.DB.QueryRowContext(ctx,
		"SELECT uuid, short_url, origin_url FROM urls WHERE short_url=$1 LIMIT 1", value)

	var rec Record
	err := row.Scan(&rec.UUID, &rec.ShortULR, &rec.OriginalURL)
	if err != nil {
		return rec, err
	}

	return rec, nil
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
