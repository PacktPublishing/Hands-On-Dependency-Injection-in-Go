package advantages

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	// thrown when the supplied order does not exist in the database
	errNotFound = errors.New("order not found")
)

// Loads orders based on supplied owner and order ID
type OrderLoader interface {
	loadOrder(owner Owner, orderID int) (Order, error)
}

// NewLoadOrderHandler creates a new instance of LoadOrderHandler
func NewLoadOrderHandler(loader OrderLoader) *LoadOrderHandler {
	return &LoadOrderHandler{
		loader: loader,
	}
}

// LoadOrderHandler is a HTTP handler that loads orders based on the current user and supplied user ID
type LoadOrderHandler struct {
	loader OrderLoader
}

// ServeHTTP implements http.Handler
func (l *LoadOrderHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// extract user from supplied authentication credentials
	currentUser, err := l.authenticateUser(request)
	if err != nil {
		response.WriteHeader(http.StatusUnauthorized)
		return
	}

	// extract order ID from request
	orderID, err := l.extractOrderID(request)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	// load order using the current user as a request-scoped dependency
	// (with method injection)
	order, err := l.loader.loadOrder(currentUser, orderID)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	// output order
	encoder := json.NewEncoder(response)
	err = encoder.Encode(order)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
}

// AuthenticatedLoader will load orders for based on the supplied owner
type AuthenticatedLoader struct {
	// This pool is expensive to create.  We will want to create it once and then reuse it.
	db *sql.DB
}

// load the order from the database based on owner and order ID
func (a *AuthenticatedLoader) loadByOwner(owner Owner, orderID int) (*Order, error) {
	order, err := a.load(orderID)
	if err != nil {
		return nil, err
	}

	if order.OwnerID != owner.ID() {
		// Return not found so we do not leak information to hackers
		return nil, errNotFound
	}

	// happy path
	return order, nil
}

func (a *AuthenticatedLoader) load(orderID int) (*Order, error) {
	// load order from DB
	return &Order{OwnerID: 1}, nil
}

type Owner interface {
	ID() int
}

type Order struct {
	OwnerID int

	// other order details
}

type User struct {
	id int

	// other attributes
}

func (u *User) ID() int {
	return u.id
}

// Extract the user from the request (e.g. from a JWT token).
func (l *LoadOrderHandler) authenticateUser(request *http.Request) (*User, error) {
	return &User{id: 1}, nil
}

// Extract the order ID from the request (e.g. from the URL or HTTP POST body)
func (l *LoadOrderHandler) extractOrderID(request *http.Request) (int, error) {
	return 2, nil
}
