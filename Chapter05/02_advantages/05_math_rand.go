package advantages

// A Rand is a source of random numbers.
type Rand struct {
	src Source

	// code removed
}

// Int returns a non-negative pseudo-random int.
func (r *Rand) Int() int {
	// code changed for brevity
	value := r.src.Int63()
	return int(value)
}

/*
 * Top-level convenience functions
 */

var globalRand = New(&lockedSource{})

// Int returns a non-negative pseudo-random int from the default Source.
func Int() int { return globalRand.Int() }

/*
 * Code below here has been modified so that it compiles but does nothing.
 * The original code is: https://golang.org/src/math/rand/rand.go
 */

// New returns a new Rand that uses random values from src
// to generate other random values.
func New(src Source) *Rand {
	// code changed for brevity
	return &Rand{
		src: src,
	}
}

type lockedSource struct {
	// code removed
}

func (l *lockedSource) Int63() int64 {
	// code removed
	return 0
}

// A Source represents a source of uniformly-distributed
// pseudo-random int64 values in the range [0, 1<<63).
type Source interface {
	Int63() int64

	// code removed
}
