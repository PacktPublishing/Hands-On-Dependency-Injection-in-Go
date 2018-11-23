package welcome

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLongMethod_happyPath(t *testing.T) {
	// build request
	request := &http.Request{}
	request.PostForm = url.Values{}
	request.PostForm.Add("UserID", "123")

	// mock the database
	var mockDB sqlmock.Sqlmock
	var err error

	DB, mockDB, err = sqlmock.New()
	require.NoError(t, err)
	mockDB.ExpectQuery("SELECT .* FROM people WHERE ID = ?").
		WithArgs(123).
		WillReturnRows(sqlmock.NewRows([]string{"ID", "Name", "Phone"}).AddRow(123, "May", "0123456789"))

	// build response
	response := httptest.NewRecorder()

	// call method
	longMethod(response, request)

	// validate response
	require.Equal(t, http.StatusOK, response.Code)

	// validate the JSON
	responseBytes, err := ioutil.ReadAll(response.Body)
	require.NoError(t, err)

	expectedJSON := `{"ID":123,"Name":"May","Phone":"0123456789"}` + "\n"
	assert.Equal(t, expectedJSON, string(responseBytes))
}
