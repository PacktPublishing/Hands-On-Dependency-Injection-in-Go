package exchange

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"time"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/acme/internal/logging"
)

const (
	// request URL for the exchange rate API
	urlFormat = "%s/api/historical?access_key=%s&date=2018-06-20&currencies=%s"

	// default price that is sent when an error occurs
	defaultPrice = 0.0
)

// NewConverter creates and initializes the converter
func NewConverter(cfg Config) *Converter {
	return &Converter{
		cfg: cfg,
	}
}

// Config is the config for Converter
type Config interface {
	Logger() logging.Logger
	ExchangeBaseURL() string
	ExchangeAPIKey() string
}

// Converter will convert the base price to the currency supplied
// Note: we are expecting sane inputs and therefore skipping input validation
type Converter struct {
	cfg Config
}

// Exchange will perform the conversion
func (c *Converter) Exchange(ctx context.Context, basePrice float64, currency string) (float64, error) {
	// load rate from the external API
	response, err := c.loadRateFromServer(ctx, currency)
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
func (c *Converter) loadRateFromServer(ctx context.Context, currency string) (*http.Response, error) {
	// build the request
	url := fmt.Sprintf(urlFormat,
		c.cfg.ExchangeBaseURL(),
		c.cfg.ExchangeAPIKey(),
		currency)

	// perform request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.logger().Warn("[exchange] failed to create request. err: %s", err)
		return nil, err
	}

	// set latency budget for the upstream call
	subCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// replace the default context with our custom one
	req = req.WithContext(subCtx)

	// perform the HTTP request
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		c.logger().Warn("[exchange] failed to load. err: %s", err)
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("request failed with code %d", response.StatusCode)
		c.logger().Warn("[exchange] %s", err)
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
		c.logger().Error("[exchange] %s", err)
		return defaultPrice, err
	}

	// happy path
	return rate, nil
}

func (c *Converter) extractResponse(response *http.Response) (*apiResponseFormat, error) {
	payload, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.logger().Error("[exchange] failed to ready response body. err: %s", err)
		return nil, err
	}

	data := &apiResponseFormat{}
	err = json.Unmarshal(payload, data)
	if err != nil {
		c.logger().Error("[exchange] error converting response. err: %s", err)
		return nil, err
	}

	// happy path
	return data, nil
}

func (c *Converter) logger() logging.Logger {
	return c.cfg.Logger()
}

// the response format from the exchange rate API
type apiResponseFormat struct {
	Quotes map[string]float64 `json:"quotes"`
}
