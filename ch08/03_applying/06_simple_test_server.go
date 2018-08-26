// +build do-not-build

package applying

import (
	"net/http"
)

type happyExchangeRateService struct{}

// ServeHTTP implements http.Handler
func (*happyExchangeRateService) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	payload := []byte(`
{
  "success":true,
  "timestamp":1535250248,
  "base":"EUR",
  "date":"2018-08-26",
  "rates": {
	"AUD":1.587884
  }
}
`)
	response.Write(payload)
}
