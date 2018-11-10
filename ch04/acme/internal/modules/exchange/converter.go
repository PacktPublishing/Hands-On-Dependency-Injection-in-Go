package exchange

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch04/acme/internal/config"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch04/acme/internal/logging"
)

const (
	// request URL for the exchange rate API
	urlFormat = "%s/api/historical?access_key=%s&date=2018-06-20&currencies=%s"

	// default price that is sent when an error occurs
	defaultPrice = 0.0
)

// Converter will convert the base price to the currency supplied
// Note: we are expecting sane inputs and therefore skipping input validation
type Converter struct{}

// Do will perform the conversion
func (c *Converter) Do(basePrice float64, currency string) (float64, error) {
	// load rate from the external API
	response, err := c.loadRateFromServer(currency)
	if err != nil {
		return defaultPrice, err
	}

	// extract rate from response
	rate, err := c.extractRate(response, currency)
	if err != nil {
		return defaultPrice, err
	}

	// apply rate and round to 2 decimal places
	return math.Floor((basePrice/rate)*100) / 100, nil
}

// load rate from the external API
func (c *Converter) loadRateFromServer(currency string) (*http.Response, error) {
	// build the request
	url := fmt.Sprintf(urlFormat,
		config.App.ExchangeRateBaseURL,
		config.App.ExchangeRateAPIKey,
		currency)

	// perform request
	response, err := http.Get(url)
	if err != nil {
		logging.L.Warn("[exchange] failed to load. err: %s", err)
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("request failed with code %d", response.StatusCode)
		logging.L.Warn("[exchange] %s", err)
		return nil, err
	}

	return response, nil
}

func (c *Converter) extractRate(response *http.Response, currency string) (float64, error) {
	defer func() {
		_ = response.Body.Close()
	}()

	// extract data from response
	data, err := c.extractResponse(response)
	if err != nil {
		return defaultPrice, err
	}

	// pull rate from response data
	rate, found := data.Quotes["USD"+currency]
	if !found {
		err = fmt.Errorf("response did not include expected currency '%s'", currency)
		logging.L.Error("[exchange] %s", err)
		return defaultPrice, err
	}

	// happy path
	return rate, nil
}

func (c *Converter) extractResponse(response *http.Response) (*apiResponseFormat, error) {
	payload, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logging.L.Error("[exchange] failed to ready response body. err: %s", err)
		return nil, err
	}

	data := &apiResponseFormat{}
	err = json.Unmarshal(payload, data)
	if err != nil {
		logging.L.Error("[exchange] error converting response. err: %s", err)
		return nil, err
	}

	// happy path
	return data, nil
}

// the response format from the exchange rate API
type apiResponseFormat struct {
	Quotes map[string]float64 `json:"quotes"`
}
