package storage

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/arseniy96/url-shortener/internal/logger"
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

func (db *Database) SaveRecord(ctx context.Context, rec *Record) error {
	_, err := db.DB.ExecContext(ctx,
		`INSERT INTO urls(uuid, short_url, origin_url) VALUES($1, $2, $3)`,
		rec.UUID, rec.ShortULR, rec.OriginalURL)

	return err
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

func (db *Database) CreateDatabase() error {
	ctx, close := context.WithTimeout(context.Background(), 5*time.Second)
	defer close()

	_, err := db.DB.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS urls(
			"uuid" VARCHAR,
			"short_url" VARCHAR,
			"origin_url" VARCHAR)`)
	if err != nil {
		return err
	}

	return nil
}
