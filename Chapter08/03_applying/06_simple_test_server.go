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
   "historical":true,
   "date":"2010-11-09",
   "timestamp":1289347199,
   "source":"USD",
   "quotes":{
      "USDAUD":0.989981
   }
}`)
	response.Write(payload)
}
