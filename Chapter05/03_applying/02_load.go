package applying

import (
	"database/sql"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	// SQL statements as constants (to reduce duplication and maintenance in tests)
	sqlAllColumns = "id, fullname, phone, currency, price"
	sqlLoadByID   = "SELECT " + sqlAllColumns + " FROM person WHERE id = ? LIMIT 1"
)

func TestLoad_happyPath(t *testing.T) {
	expectedResult := &Person{
		ID:       2,
		FullName: "Paul",
		Phone:    "0123456789",
		Currency: "CAD",
		Price:    23.45,
	}

	// define a mock db
	testDb, dbMock, err := sqlmock.New()
	require.NoError(t, err)

	// configure the mock db
	queryRegex := convertSQLToRegex(sqlLoadByID)
	dbMock.ExpectQuery(queryRegex).WillReturnRows(
		sqlmock.NewRows(strings.Split(sqlAllColumns, ", ")).
			AddRow(2, "Paul", "0123456789", "CAD", 23.45))

	// monkey patching the database
	defer func(original sql.DB) {
		// restore original DB (after test)
		db = &original
	}(*db)

	db = testDb
	// end of monkey patch

	// call function
	result, err := Load(2)

	// validate results
	assert.Equal(t, expectedResult, result)
	assert.NoError(t, err)
	assert.NoError(t, dbMock.ExpectationsWereMet())
}

// convert SQL string to regex by treating the entire query as a literal
func convertSQLToRegex(in string) string {
	return `\Q` + in + `\E`
}

// Load will attempt to load and return a person.
// It will return ErrNotFound when the requested person does not exist.
// Any other errors returned are caused by the underlying database or our connection to it.
func Load(ID int) (*Person, error) {
	// code removed/faked for brevity
	return &Person{
		ID:       2,
		FullName: "Paul",
		Phone:    "0123456789",
		Currency: "CAD",
		Price:    23.45,
	}, nil
}

// code removed for brevity
var db = &sql.DB{}

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
