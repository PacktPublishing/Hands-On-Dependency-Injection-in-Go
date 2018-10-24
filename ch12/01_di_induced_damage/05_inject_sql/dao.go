package data

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"
)

// NewDAO will initialize the database connection pool (if not already done) and return a data access object which
// can be used to interact with the database
func NewDAO(cfg Config) *DAO {
	// initialize the db connection pool
	_, _ = getDB(cfg)

	return &DAO{
		cfg: cfg,
	}
}

// DAO is a data access object that provides an abstraction over our database interactions.
type DAO struct {
	cfg Config
}

// Load will attempt to load and return a person.
// It will return ErrNotFound when the requested person does not exist.
// Any other errors returned are caused by the underlying database or our connection to it.
func (d *DAO) Load(ctx context.Context, ID int) (*Person, error) {
	db, err := getDB(d.cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get DB connection. err: %s", err)
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
			fmt.Fprintf(os.Stderr, "failed to load requested person '%d'. err: %s", ID, err)
			return nil, ErrNotFound
		}

		fmt.Fprintf(os.Stderr, "failed to convert query result. err: %s", err)
		return nil, err
	}
	return out, nil
}

// LoadAll will attempt to load all people in the database
// It will return ErrNotFound when there are not people in the database
// Any other errors returned are caused by the underlying database or our connection to it.
func (d *DAO) LoadAll(ctx context.Context) ([]*Person, error) {
	db, err := getDB(d.cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get DB connection. err: %s", err)
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
			fmt.Fprintf(os.Stderr, "failed to convert query result. err: %s", err)
			return nil, err
		}

		out = append(out, record)
	}

	if len(out) == 0 {
		fmt.Fprintf(os.Stderr, "no people found in the database.")
		return nil, ErrNotFound
	}

	return out, nil
}

// Save will save the supplied person and return the ID of the newly created person or an error.
// Errors returned are caused by the underlying database or our connection to it.
func (d *DAO) Save(ctx context.Context, in *Person) (int, error) {
	db, err := getDB(d.cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get DB connection. err: %s", err)
		return defaultPersonID, err
	}

	// set latency budget for the database call
	subCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// perform DB insert
	result, err := db.ExecContext(subCtx, sqlInsert, in.FullName, in.Phone, in.Currency, in.Price)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to save person into DB. err: %s", err)
		return defaultPersonID, err
	}

	// retrieve and return the ID of the person created
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to retrieve id of last saved person. err: %s", err)
		return defaultPersonID, err
	}

	return int(id), nil
}
