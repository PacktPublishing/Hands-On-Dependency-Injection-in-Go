package method_injection

import (
	"io"
)

// Fprint formats using the default formats for its operands and writes to w.
// It returns the number of bytes written and any write error encountered.
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	return
}
