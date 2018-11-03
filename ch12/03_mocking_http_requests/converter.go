package mocking_http_reques

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"time"
)

const (
	// request URL for the exchange rate API
	urlFormat = "%s/api/latest?access_key=%s&format=1&symbols=%s"

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
	Logger() Logger
	ExchangeBaseURL() string
	ExchangeAPIKey() string
}

// Converter will convert the base price to the currency supplied
// Note: we are expecting sane inputs and therefore skipping input validation
type Converter struct {
	cfg       Config
	requester requester
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
	return math.Floor(rate*basePrice*100) / 100, nil
}

// load rate from the external API
func (c *Converter) loadRateFromServer(ctx context.Context, currency string) (*http.Response, error) {
	// build the request
	url := fmt.Sprintf(urlFormat,
		c.cfg.ExchangeBaseURL(),
		c.cfg.ExchangeAPIKey(),
		currency)

	// perform request
	response, err := c.getRequester().doRequest(ctx, url)
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
	rate, found := data.Rates[currency]
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

func (c *Converter) logger() Logger {
	return c.cfg.Logger()
}

func (c *Converter) getRequester() requester {
	if c.requester == nil {
		c.requester = &requesterer{}
	}

	return c.requester
}

// the response format from the exchange rate API
type apiResponseFormat struct {
	Rates map[string]float64 `json:"rates"`
}

//go:generate mockery -name=requester -case underscore -testonly -inpkg -note @generated
type requester interface {
	doRequest(ctx context.Context, url string) (*http.Response, error)
}

type requesterer struct {
}

func (r *requesterer) doRequest(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// set latency budget for the upstream call
	subCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// replace the default context with our custom one
	req = req.WithContext(subCtx)

	// perform the HTTP request
	return http.DefaultClient.Do(req)
}

type Logger interface {
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
}

type stubLogger struct{}

func (l *stubLogger) Warn(message string, args ...interface{}) {
	// do nothing
}

func (l *stubLogger) Error(message string, args ...interface{}) {
	// do nothing
}
