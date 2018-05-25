package srp

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func loadUserHandlerSRP(resp http.ResponseWriter, req *http.Request) {
	userID, err := extractIDFromRequest(req)
	if err != nil {
		resp.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	person, err := loadPersonByID(userID)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	outputPerson(resp, person)
}

func extractIDFromRequest(req *http.Request) (int64, error) {
	err := req.ParseForm()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(req.Form.Get("UserID"), 10, 64)
}

func loadPersonByID(userID int64) (*Person, error) {
	row := DB.QueryRow("SELECT * FROM Users WHERE userID = ?", userID)

	person := &Person{}
	err := row.Scan(person.ID, person.Name, person.Phone)
	if err != nil {
		return nil, err
	}
	return person, nil
}

func outputPerson(resp http.ResponseWriter, person *Person) {
	encoder := json.NewEncoder(resp)
	err := encoder.Encode(person)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
}
