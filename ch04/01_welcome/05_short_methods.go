package welcome

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func shortMethods(resp http.ResponseWriter, req *http.Request) {
	userID, err := extractUserID(req)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	person, err := loadPerson(userID)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	outputPerson(resp, person)
}

func extractUserID(req *http.Request) (int64, error) {
	err := req.ParseForm()
	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(req.Form.Get("UserID"), 10, 64)
}

func loadPerson(userID int64) (*Person, error) {
	row := DB.QueryRow("SELECT * FROM people WHERE ID = ?", userID)

	person := &Person{}
	err := row.Scan(&person.ID, &person.Name, &person.Phone)
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
