package applying

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSave_happyPath(t *testing.T) {
	// define a mock db
	testDb, dbMock, err := sqlmock.New()
	require.NoError(t, err)

	// clean up afterwards
	defer testDb.Close()

	// define the query we are expecting as regular expression
	queryRegex := `\QINSERT INTO person (fullname, phone, currency, price) VALUES (?, ?, ?, ?)\E`

	// configure the mock db
	dbMock.ExpectExec(queryRegex).WillReturnResult(sqlmock.NewResult(2, 1))

	// inputs
	person := &Person{
		FullName: "Jake Blues",
		Phone:    "01234567890",
		Currency: "AUD",
		Price:    123.45,
	}

	// call function
	resultID, err := SavePerson(testDb, person)

	// validate result
	require.NoError(t, err)
	assert.Equal(t, 2, resultID)
	assert.NoError(t, dbMock.ExpectationsWereMet())
}

func SavePerson(db *sql.DB, in *Person) (int, error) {
	// perform DB insert
	query := "INSERT INTO person (fullname, phone, currency, price) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, in.FullName, in.Phone, in.Currency, in.Price)
	if err != nil {
		return 0, err
	}

	// retrieve and return the ID of the person created
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
