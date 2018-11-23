package needless_indirection

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	router := http.NewServeMux()
	router.HandleFunc("/health", func(resp http.ResponseWriter, req *http.Request) {
		_, _ = resp.Write([]byte(`OK`))
	})

	// start a server
	address := ":8080"
	go func() {
		_ = http.ListenAndServe(address, router)
	}()

	// call the server
	resp, err := http.Get("http://:8080/health")
	require.NoError(t, err)

	// validate the response
	responseBody, err := ioutil.ReadAll(resp.Body)
	assert.Equal(t, []byte(`OK`), responseBody)
}
