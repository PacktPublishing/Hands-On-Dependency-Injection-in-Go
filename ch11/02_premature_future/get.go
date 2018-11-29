package rest

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	// default person id (returned on error)
	defaultPersonID = 0

	// key in the mux where the ID is stored
	muxVarID = "id"
)

// GetModel will load a registration
type GetModel interface {
	Do(ID int) (*Person, error)
}

// GetConfig is the config for the Get Handler
type GetConfig interface {
	Logger() Logger
}

// Formatter will convert the supplied object to bytes
type Formatter interface {
	Marshal(interface{}) ([]byte, error)
}

// NewGetHandler is the constructor for GetHandler
func NewGetHandler(cfg GetConfig, model GetModel, formatter Formatter) *GetHandler {
	return &GetHandler{
		cfg:       cfg,
		getter:    model,
		formatter: formatter,
	}
}

// GetHandler is the HTTP handler for the "Get Person" endpoint
type GetHandler struct {
	cfg       GetConfig
	getter    GetModel
	formatter Formatter
}

// ServeHTTP implements http.Handler
func (h *GetHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// extract person id from request
	id, err := h.extractID(request)
	if err != nil {
		// output error
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	// attempt get
	person, err := h.getter.Do(id)
	if err != nil {
		// not need to log here as we can expect other layers to do so
		response.WriteHeader(http.StatusNotFound)
		return
	}

	// happy path
	err = h.buildOutput(response, person)
	if err != nil {
		// this error should not happen but if it does there is nothing we can do to recover
		response.WriteHeader(http.StatusInternalServerError)
	}
}

// extract the person ID from the request
func (h *GetHandler) extractID(request *http.Request) (int, error) {
	// ID is part of the URL, so we extract it from there
	vars := mux.Vars(request)
	idAsString, exists := vars[muxVarID]
	if !exists {
		// log and return error
		err := errors.New("[get] person id missing from request")
		h.cfg.Logger().Warn(err.Error())
		return defaultPersonID, err
	}

	// convert ID to int
	id, err := strconv.Atoi(idAsString)
	if err != nil {
		// log and return error
		err = fmt.Errorf("[get] failed to convert person id into a number. err: %s", err)
		h.cfg.Logger().Error(err.Error())
		return defaultPersonID, err
	}

	return id, nil
}

// output the supplied person
func (h *GetHandler) buildOutput(writer io.Writer, person *Person) error {
	output := &getResponseFormat{
		ID:       person.ID,
		FullName: person.FullName,
		Phone:    person.Phone,
		Currency: person.Currency,
		Price:    person.Price,
	}

	// build output payload
	payload, err := h.formatter.Marshal(output)
	if err != nil {
		return err
	}

	// write payload to response and return
	_, err = writer.Write(payload)
	return err
}

// the JSON response format
type getResponseFormat struct {
	ID       int     `json:"id"`
	FullName string  `json:"name"`
	Phone    string  `json:"phone"`
	Currency string  `json:"currency"`
	Price    float64 `json:"price"`
}

type Person struct {
	ID       int
	FullName string
	Phone    string
	Currency string
	Price    float64
}

type Logger interface {
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
}
