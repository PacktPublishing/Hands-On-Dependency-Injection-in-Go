package data

import (
	"database/sql"
	"errors"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSave_happyPath(t *testing.T) {
	// define a mock db
	testDb, dbMock, err := sqlmock.New()
	defer testDb.Close()
	require.NoError(t, err)

	// configure the mock db
	queryRegex := convertSQLToRegex(sqlInsert)
	dbMock.ExpectExec(queryRegex).WillReturnResult(sqlmock.NewResult(2, 1))

	// monkey patching starts here
	defer func(original sql.DB) {
		// restore original DB (after test)
		db = &original
	}(*db)

	// replace db for this test
	db = testDb
	// end of monkey patch

	// inputs
	in := &Person{
		FullName: "Jake Blues",
		Phone:    "01234567890",
		Currency: "AUD",
		Price:    123.45,
	}

	// call function
	resultID, err := Save(in)

	// validate result
	require.NoError(t, err)
	assert.Equal(t, 2, resultID)
	assert.NoError(t, dbMock.ExpectationsWereMet())
}

func TestSave_insertError(t *testing.T) {
	// define a mock db
	testDb, dbMock, err := sqlmock.New()
	defer testDb.Close()

	require.NoError(t, err)

	// configure the mock db
	queryRegex := convertSQLToRegex(sqlInsert)
	dbMock.ExpectExec(queryRegex).WillReturnError(errors.New("failed to insert"))

	// monkey patching starts here
	defer func(original sql.DB) {
		// restore original DB (after test)
		db = &original
	}(*db)

	// replace db for this test
	db = testDb
	// end of monkey patch

	// inputs
	in := &Person{
		FullName: "Jake Blues",
		Phone:    "01234567890",
		Currency: "AUD",
		Price:    123.45,
	}

	// call function
	resultID, err := Save(in)

	// validate result
	require.Error(t, err)
	assert.Equal(t, defaultPersonID, resultID)
	assert.NoError(t, dbMock.ExpectationsWereMet())
}

func TestSave_getDBError(t *testing.T) {
	// monkey patching starts here
	defer func(original func() (*sql.DB, error)) {
		// restore original DB (after test)
		getDB = original
	}(getDB)

	// replace getDB() function for this test
	getDB = func() (*sql.DB, error) {
		return nil, errors.New("getDB() failed")
	}
	// end of monkey patch

	// inputs
	in := &Person{
		FullName: "Jake Blues",
		Phone:    "01234567890",
		Currency: "AUD",
		Price:    123.45,
	}

	// call function
	resultID, err := Save(in)
	require.Error(t, err)
	assert.Equal(t, defaultPersonID, resultID)
}

func TestLoadAll_tableDrivenTest(t *testing.T) {
	scenarios := []struct {
		desc            string
		configureMockDB func(sqlmock.Sqlmock)
		expectedResults []*Person
		expectError     bool
	}{
		{
			desc: "happy path",
			configureMockDB: func(dbMock sqlmock.Sqlmock) {
				queryRegex := convertSQLToRegex(sqlLoadAll)
				dbMock.ExpectQuery(queryRegex).WillReturnRows(
					sqlmock.NewRows(strings.Split(sqlAllColumns, ", ")).
						AddRow(1, "John", "0123456789", "AUD", 12.34))
			},
			expectedResults: []*Person{
				{
					ID:       1,
					FullName: "John",
					Phone:    "0123456789",
					Currency: "AUD",
					Price:    12.34,
				},
			},
			expectError: false,
		},
		{
			desc: "load error",
			configureMockDB: func(dbMock sqlmock.Sqlmock) {
				queryRegex := convertSQLToRegex(sqlLoadAll)
				dbMock.ExpectQuery(queryRegex).WillReturnError(errors.New("something failed"))
			},
			expectedResults: nil,
			expectError:     true,
		},
	}

	for _, scenario := range scenarios {
		// define a mock db
		testDb, dbMock, err := sqlmock.New()
		require.NoError(t, err)

		// configure the mock db
		scenario.configureMockDB(dbMock)

		// monkey patch the db for this test
		original := *db
		db = testDb

		// call function
		results, err := LoadAll()

		// validate results
		assert.Equal(t, scenario.expectedResults, results, scenario.desc)
		assert.Equal(t, scenario.expectError, err != nil, scenario.desc)
		assert.NoError(t, dbMock.ExpectationsWereMet())

		// restore original DB (after test)
		db = &original
		testDb.Close()
	}
}

func TestLoad_tableDrivenTest(t *testing.T) {
	scenarios := []struct {
		desc            string
		configureMockDB func(sqlmock.Sqlmock)
		expectedResult  *Person
		expectError     bool
	}{
		{
			desc: "happy path",
			configureMockDB: func(dbMock sqlmock.Sqlmock) {
				queryRegex := convertSQLToRegex(sqlLoadAll)
				dbMock.ExpectQuery(queryRegex).WillReturnRows(
					sqlmock.NewRows(strings.Split(sqlAllColumns, ", ")).
						AddRow(2, "Paul", "0123456789", "CAD", 23.45))
			},
			expectedResult: &Person{
				ID:       2,
				FullName: "Paul",
				Phone:    "0123456789",
				Currency: "CAD",
				Price:    23.45,
			},
			expectError: false,
		},
		{
			desc: "load error",
			configureMockDB: func(dbMock sqlmock.Sqlmock) {
				queryRegex := convertSQLToRegex(sqlLoadAll)
				dbMock.ExpectQuery(queryRegex).WillReturnError(errors.New("something failed"))
			},
			expectedResult: nil,
			expectError:    true,
		},
	}

	for _, scenario := range scenarios {
		// define a mock db
		testDb, dbMock, err := sqlmock.New()
		require.NoError(t, err)

		// configure the mock db
		scenario.configureMockDB(dbMock)

		// monkey db for this test
		original := *db
		db = testDb

		// call function
		result, err := Load(2)

		// validate results
		assert.Equal(t, scenario.expectedResult, result, scenario.desc)
		assert.Equal(t, scenario.expectError, err != nil, scenario.desc)
		assert.NoError(t, dbMock.ExpectationsWereMet())

		// restore original DB (after test)
		db = &original
		testDb.Close()
	}
}

// convert SQL string to regex by treating the entire query as a literal
func convertSQLToRegex(in string) string {
	return `\Q` + in + `\E`
}
