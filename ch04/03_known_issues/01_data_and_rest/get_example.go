//+build ignore

package data_and_rest

import (
	"encoding/json"
	"io"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch06/03_applying/01/data"
)

// output the supplied person as JSON
func (h *GetHandler) writeJSON(writer io.Writer, person *data.Person) error {
	// call to http.ResponseWriter.Write() will cause HTTP OK (200) to be output as well
	return json.NewEncoder(writer).Encode(person)
}
