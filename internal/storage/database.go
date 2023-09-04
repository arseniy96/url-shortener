package storage

import (
	"context"
	"database/sql"
	"strings"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/arseniy96/url-shortener/internal/logger"
)

const (
	CreateDBTimeOut = 5 * time.Second
)

type Database struct {
	DB *sql.DB
}

func NewDatabase(connectionData string) (*Database, error) {
	// connectionData = "host=localhost user=shortener password= dbname=shortener sslmode=disable"
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
		"SELECT uuid, short_url, origin_url, is_deleted FROM urls WHERE short_url=$1 LIMIT 1", value)

	var rec Record
	err := row.Scan(&rec.UUID, &rec.ShortULR, &rec.OriginalURL, &rec.DeletedFlag)
	if err != nil {
		return rec, err
	}

	return rec, nil
}

func (db *Database) FindRecordByOriginURL(ctx context.Context, value string) (Record, error) {
	row := db.DB.QueryRowContext(ctx,
		"SELECT uuid, short_url, origin_url, is_deleted FROM urls WHERE origin_url=$1 LIMIT 1", value)

	var rec Record
	err := row.Scan(&rec.UUID, &rec.ShortULR, &rec.OriginalURL, &rec.DeletedFlag)
	if err != nil {
		return rec, err
	}

	return rec, nil
}

func (db *Database) SaveRecord(ctx context.Context, rec *Record, userID int) error {
	_, err := db.DB.ExecContext(ctx,
		`INSERT INTO urls(uuid, short_url, origin_url, user_id) VALUES($1, $2, $3, $4)`,
		rec.UUID, rec.ShortULR, rec.OriginalURL, userID)

	return err
}

func (db *Database) SaveRecordsBatch(ctx context.Context, records []Record) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			logger.Log.Error(err)
		}
	}()

	for _, rec := range records {
		_, err := tx.ExecContext(ctx,
			`INSERT INTO urls(uuid, short_url, origin_url) VALUES($1, $2, $3)`,
			rec.UUID, rec.ShortULR, rec.OriginalURL)
		if err != nil {
			if err2 := tx.Rollback(); err2 != nil {
				return err2
			}
			return err
		}
	}
	return tx.Commit()
}

func (db *Database) HealthCheck() error {
	ctx, cancel := context.WithTimeout(context.Background(), TimeOut)
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
	ctx, cancel := context.WithTimeout(context.Background(), CreateDBTimeOut)
	defer cancel()

	_, err := db.DB.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS users(
			"id" SERIAL PRIMARY KEY,
			"cookie" VARCHAR)`)
	if err != nil {
		return err
	}

	_, err = db.DB.ExecContext(ctx,
		`CREATE UNIQUE INDEX IF NOT EXISTS cookie_idx on users(cookie)`)
	if err != nil {
		return err
	}

	_, err = db.DB.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS urls(
			"uuid" VARCHAR,
			"short_url" VARCHAR,
			"origin_url" VARCHAR,
			"user_id" INTEGER,
			"is_deleted" BOOLEAN DEFAULT false)`)
	if err != nil {
		return err
	}

	_, err = db.DB.ExecContext(ctx,
		`CREATE UNIQUE INDEX IF NOT EXISTS origin_url_idx on urls(origin_url)`)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) FindRecordsByUserID(ctx context.Context, userID int) (records []Record, err error) {
	rows, err := db.DB.QueryContext(ctx,
		"SELECT uuid, short_url, origin_url, is_deleted FROM urls WHERE user_id=$1", userID)
	if err != nil {
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			logger.Log.Error(err)
		}
	}()
	for rows.Next() {
		var rec Record
		err = rows.Scan(&rec.UUID, &rec.ShortULR, &rec.OriginalURL, &rec.DeletedFlag)
		if err != nil {
			return
		}

		records = append(records, rec)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

func (db *Database) FindUserByCookie(ctx context.Context, cookie string) (*User, error) {
	row := db.DB.QueryRowContext(ctx,
		"SELECT id, cookie FROM users WHERE cookie=$1 LIMIT 1", cookie)

	var user User
	err := row.Scan(&user.UserID, &user.Cookie)
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (db *Database) FindUserByID(ctx context.Context, userID int) (*User, error) {
	row := db.DB.QueryRowContext(ctx,
		"SELECT id, cookie FROM users WHERE id=$1 LIMIT 1", userID)

	var user User
	err := row.Scan(&user.UserID, &user.Cookie)
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (db *Database) CreateUser(ctx context.Context) (*User, error) {
	_, err := db.DB.ExecContext(ctx, `INSERT INTO users DEFAULT VALUES`)
	if err != nil {
		return nil, err
	}

	row := db.DB.QueryRowContext(ctx,
		"SELECT id FROM users ORDER BY id DESC LIMIT 1")
	var user User
	err = row.Scan(&user.UserID)
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (db *Database) UpdateUser(ctx context.Context, id int, cookie string) error {
	_, err := db.DB.ExecContext(ctx, `UPDATE users SET cookie=$1 WHERE id=$2`, cookie, id)
	return err
}

func (db *Database) FindRecordsBatchByShortURL(ctx context.Context, urls []string) (records []Record, err error) {
	params := paramsBuilder(urls)

	rows, err := db.DB.QueryContext(ctx,
		"SELECT uuid, short_url, origin_url, user_id, is_deleted FROM urls WHERE short_url = ANY($1::text[]);",
		params,
	)
	if err != nil {
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			logger.Log.Error(err)
		}
	}()

	for rows.Next() {
		var rec Record
		err = rows.Scan(&rec.UUID, &rec.ShortULR, &rec.OriginalURL, &rec.UserID, &rec.DeletedFlag)
		if err != nil {
			return
		}

		records = append(records, rec)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

func (db *Database) DeleteBatchRecords(ctx context.Context, records []Record) error {
	urls := make([]string, 0, len(records))

	for _, rec := range records {
		urls = append(urls, rec.ShortULR)
	}

	params := paramsBuilder(urls)

	_, err := db.DB.ExecContext(ctx, "UPDATE urls SET is_deleted=true WHERE short_url = ANY($1::text[]);",
		params)
	return err
}

func paramsBuilder(urls []string) string {
	return "{" + strings.Join(urls, ",") + "}"
}
