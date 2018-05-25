package ocp

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

func GetUserHandlerV2(resp http.ResponseWriter, req *http.Request) {
	// validate inputs
	err := req.ParseForm()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	userID, err := extractUserID(req.Form)
	if err != nil {
		resp.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	user := loadUser(userID)
	outputUser(resp, user)
}

func DeleteUserHandlerV2(resp http.ResponseWriter, req *http.Request) {
	// validate inputs
	err := req.ParseForm()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	userID, err := extractUserID(req.Form)
	if err != nil {
		resp.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	deleteUser(userID)
}

func extractUserID(values url.Values) (int64, error) {
	userID, err := strconv.ParseInt(values.Get("UserID"), 10, 64)
	if err != nil {
		return 0, err
	}

	if userID <= 0 {
		return 0, errors.New("userID must be positive")
	}

	return userID, nil
}
