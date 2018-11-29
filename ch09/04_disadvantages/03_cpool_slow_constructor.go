package disadvantages

func newConnectionPool() ConnectionPool {
	pool := &myConnectionPool{}

	// initialize the pool
	pool.init()

	// return a "ready to use pool"
	return pool
}
