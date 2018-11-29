package method_injection

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	// validate method
	m, err := validateMethod(method)
	if err != nil {
		return nil, err
	}

	// validate URL
	u, err := validateURL(url)
	if err != nil {
		return nil, err
	}

	// process body (if exists)
	var b io.ReadCloser
	if body != nil {
		// read body
		b = ioutil.NopCloser(body)
	}

	// build Request and return
	req := &http.Request{
		URL:    u,
		Method: m,
		Body:   b,
	}

	return req, nil
}

func validateMethod(method string) (string, error) {
	return "", nil
}

func validateURL(url string) (*url.URL, error) {
	return nil, nil
}
