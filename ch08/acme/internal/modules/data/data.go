package data

import (
	// import the MySQL Driver
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/acme/internal/logging"
	_ "github.com/go-sql-driver/mysql"
)

const (
	// default person id (returned on error)
	defaultPersonID = 0

	// SQL statements as constants (to reduce duplication and maintenance in tests)
	sqlAllColumns = "id, fullname, phone, currency, price"
	sqlInsert     = "INSERT INTO person (fullname, phone, currency, price) VALUES (?, ?, ?, ?)"
	sqlLoadAll    = "SELECT " + sqlAllColumns + " FROM person"
	sqlLoadByID   = "SELECT " + sqlAllColumns + " FROM person WHERE id = ? LIMIT 1"
)

var (
	db *sql.DB

	// ErrNotFound is returned when the no records where matched by the query
	ErrNotFound = errors.New("not found")
)

// Config is the configuration for the data package
type Config interface {
	// Logger returns a reference to the logger
	Logger() logging.Logger

	// DataDSN returns the data source name
	DataDSN() string
}

var getDB = func(cfg Config) (*sql.DB, error) {
	if db == nil {
		var err error
		db, err = sql.Open("mysql", cfg.DataDSN())
		if err != nil {
			// if the DB cannot be accessed we are dead
			panic(err.Error())
		}
	}

	return db, nil
}

// Person is the data transfer object (DTO) for this package
type Person struct {
	// ID is the unique ID for this person
	ID int
	// FullName is the name of this person
	FullName string
	// Phone is the phone for this person
	Phone string
	// Currency is the currency this person has paid in
	Currency string
	// Price is the amount (in the above currency) paid by this person
	Price float64
}

// Save will save the supplied person and return the ID of the newly created person or an error.
// Errors returned are caused by the underlying database or our connection to it.
func Save(ctx context.Context, cfg Config, in *Person) (int, error) {
	db, err := getDB(cfg)
	if err != nil {
		cfg.Logger().Error("failed to get DB connection. err: %s", err)
		return defaultPersonID, err
	}

	// set latency budget for the database call
	subCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// perform DB insert
	result, err := db.ExecContext(subCtx, sqlInsert, in.FullName, in.Phone, in.Currency, in.Price)
	if err != nil {
		cfg.Logger().Error("failed to save person into DB. err: %s", err)
		return defaultPersonID, err
	}

	// retrieve and return the ID of the person created
	id, err := result.LastInsertId()
	if err != nil {
		cfg.Logger().Error("failed to retrieve id of last saved person. err: %s", err)
		return defaultPersonID, err
	}

	return int(id), nil
}

// LoadAll will attempt to load all people in the database
// It will return ErrNotFound when there are not people in the database
// Any other errors returned are caused by the underlying database or our connection to it.
func LoadAll(ctx context.Context, cfg Config) ([]*Person, error) {
	db, err := getDB(cfg)
	if err != nil {
		cfg.Logger().Error("failed to get DB connection. err: %s", err)
		return nil, err
	}

	// set latency budget for the database call
	subCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// perform DB select
	rows, err := db.QueryContext(subCtx, sqlLoadAll)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	var out []*Person

	for rows.Next() {
		// retrieve columns and populate the person object
		record, err := populatePerson(rows.Scan)
		if err != nil {
			cfg.Logger().Error("failed to convert query result. err: %s", err)
			return nil, err
		}

		out = append(out, record)
	}

	if len(out) == 0 {
		cfg.Logger().Warn("no people found in the database.")
		return nil, ErrNotFound
	}

	return out, nil
}

// Load will attempt to load and return a person.
// It will return ErrNotFound when the requested person does not exist.
// Any other errors returned are caused by the underlying database or our connection to it.
func Load(ctx context.Context, cfg Config, ID int) (*Person, error) {
	db, err := getDB(cfg)
	if err != nil {
		cfg.Logger().Error("failed to get DB connection. err: %s", err)
		return nil, err
	}

	// set latency budget for the database call
	subCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// perform DB select
	row := db.QueryRowContext(subCtx, sqlLoadByID, ID)

	// retrieve columns and populate the person object
	out, err := populatePerson(row.Scan)
	if err != nil {
		if err == sql.ErrNoRows {
			cfg.Logger().Warn("failed to load requested person '%d'. err: %s", ID, err)
			return nil, ErrNotFound
		}

		cfg.Logger().Error("failed to convert query result. err: %s", err)
		return nil, err
	}
	return out, nil
}

// custom type so we can convert sql results to easily
type scanner func(dest ...interface{}) error

// reduce the duplication (and maintenance) between sql.Row and sql.Rows usage
func populatePerson(scanner scanner) (*Person, error) {
	out := &Person{}
	err := scanner(&out.ID, &out.FullName, &out.Phone, &out.Currency, &out.Price)
	return out, err
}
