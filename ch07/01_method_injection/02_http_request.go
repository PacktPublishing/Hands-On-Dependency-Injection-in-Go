package method_injection

import (
	"bytes"
	"fmt"
	"net/http"
)

func ExampleB() {
	// added to make the compiler happy
	body := &bytes.Buffer{}

	// example is here
	req, err := http.NewRequest("POST", "/login", body)

	// added to make the compiler happy
	fmt.Printf("req: %#v / err: %s", req, err)
}
