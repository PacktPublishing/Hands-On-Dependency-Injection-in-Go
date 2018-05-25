package srp

import (
	"fmt"
	"io"
)

// CalculatorV1 calculates the test coverage for a directory and it's sub-directories
type CalculatorV1 struct {
	// coverage data populated by `Calculate()` method
	data map[string]float64
}

// Calculate will calculate the coverage
func (c *CalculatorV1) Calculate(path string) error {
	// run `go test -cover ./[path]/...` and store the results
	return nil
}

// Output will print the coverage data to the supplied writer
func (c *CalculatorV1) Output(writer io.Writer) {
	for path, result := range c.data {
		fmt.Fprintf(writer, "%s -> %.1f\n", path, result)
	}
}
