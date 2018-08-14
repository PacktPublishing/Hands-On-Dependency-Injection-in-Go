package method_injection

import (
	"fmt"
	"os"
)

func ExampleA() {
	fmt.Fprint(os.Stdout, "Hello World")
}
