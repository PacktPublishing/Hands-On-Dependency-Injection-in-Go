package ocp

import (
	"net/http"
	"strconv"
)

func GetUserHandlerV1(resp http.ResponseWriter, req *http.Request) {
	// validate inputs
	err := req.ParseForm()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	userID, err := strconv.ParseInt(req.Form.Get("UserID"), 10, 64)
	if err != nil {
		resp.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	user := loadUser(userID)
	outputUser(resp, user)
}

func DeleteUserHandlerV1(resp http.ResponseWriter, req *http.Request) {
	// validate inputs
	err := req.ParseForm()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	userID, err := strconv.ParseInt(req.Form.Get("UserID"), 10, 64)
	if err != nil {
		resp.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	deleteUser(userID)
}

func loadUser(userID int64) interface{} {
	// TODO: implement
	return nil
}

func deleteUser(userID int64) {
	// TODO: implement
}

func outputUser(resp http.ResponseWriter, user interface{}) {
	// TODO: implement
}
