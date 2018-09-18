package _4_tight_coupling

import (
	"errors"
	"time"
)

type searchRequest struct {
	query string
	start time.Time
	end   time.Time
}

func (request searchRequest) validate() error {
	if request.query == "" {
		return errors.New("search term is missing")
	}
	if request.start.IsZero() || request.start.After(time.Now()) {
		return errors.New("start time is missing or invalid")
	}
	if request.end.IsZero() || request.end.Before(request.start) {
		return errors.New("end time is missing or invalid")
	}

	return nil
}

type searchResults struct {
	result string
}

func doSearchWithEnvy(request searchRequest) ([]searchResults, error) {
	// validate request
	if request.query == "" {
		return nil, errors.New("search term is missing")
	}
	if request.start.IsZero() || request.start.After(time.Now()) {
		return nil, errors.New("start time is missing or invalid")
	}
	if request.end.IsZero() || request.end.Before(request.start) {
		return nil, errors.New("end time is missing or invalid")
	}

	return performSearch(request)
}

func doSearchWithoutEnvy(request searchRequest) ([]searchResults, error) {
	err := request.validate()
	if err != nil {
		return nil, err
	}

	return performSearch(request)
}

func performSearch(request searchRequest) ([]searchResults, error) {
	// TODO: implement
	return nil, errors.New("not implemented")
}
