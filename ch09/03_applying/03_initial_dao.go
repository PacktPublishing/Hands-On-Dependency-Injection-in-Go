//+build willNotCompile

package applying

import (
	"context"
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
	return Load(ctx, d.cfg, ID)
}
