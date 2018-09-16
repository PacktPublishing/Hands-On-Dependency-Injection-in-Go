package disadvantages

func (l *Sender) getConnectionPoolOnce() ConnectionPool {
	l.initPoolOnce.Do(func() {
		l.connectionPool = newConnectionPool()
	})

	return l.connectionPool
}
