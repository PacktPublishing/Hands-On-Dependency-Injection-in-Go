package needless_indirection

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildRouter(t *testing.T) {
	router := http.NewServeMux()

	// call function
	buildRouter(router)

	// assertions
	assert.IsType(t, &getEndpoint{}, extractHandler(router, "/get"))
	assert.IsType(t, &listEndpoint{}, extractHandler(router, "/list"))
	assert.IsType(t, &saveEndpoint{}, extractHandler(router, "/save"))
}

func extractHandler(router *http.ServeMux, path string) http.Handler {
	req, _ := http.NewRequest("GET", path, nil)
	handler, _ := router.Handler(req)
	return handler
}
