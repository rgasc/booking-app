package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// DB holds the database connection pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const MAX_OPEN = 10
const MAX_IDLE = 5
const MAX_LIFETIME = 5 * time.Minute

func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(MAX_OPEN)
	d.SetConnMaxIdleTime(MAX_IDLE)
	d.SetConnMaxLifetime(MAX_LIFETIME)

	dbConn.SQL = d

	err = testDB(d)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func testDB(d *sql.DB) error {
	err := d.Ping()

	if err != nil {
		return err
	}

	return nil
}

func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
