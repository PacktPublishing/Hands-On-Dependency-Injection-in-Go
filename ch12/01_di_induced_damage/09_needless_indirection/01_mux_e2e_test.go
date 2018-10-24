package needless_indirection

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildRouter(t *testing.T) {
	router := http.NewServeMux()

	// call function
	buildRouter(router)

	// start a server
	address := ":8080"
	go func() {
		_ = http.ListenAndServe(address, router)
	}()

	// call endpoints
	responseBody := doGet(t, address+"/get")
	assert.Equal(t, `Hi from Get!`, responseBody)

	responseBody = doGet(t, address+"/list")
	assert.Equal(t, `Hi from List!`, responseBody)

	responseBody = doGet(t, address+"/save")
	assert.Equal(t, `Hi from Save!`, responseBody)

}

func doGet(t *testing.T, address string) string {
	resp, err := http.Get("http://" + address)
	require.NoError(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	defer resp.Body.Close()
	return string(body)
}
