package disadvantages

func (l *Sender) SendWithoutReadyCheck(payload []byte) error {
	pool := l.getConnectionPool()

	// get connection from pool and return afterwards
	conn := pool.Get()
	defer l.connectionPool.Release(conn)

	// send and return
	_, err := conn.Write(payload)

	return err
}
