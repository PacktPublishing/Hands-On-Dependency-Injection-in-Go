package disadvantages

import (
	"context"
	"errors"
	"net"
	"sync"
)

type ConnectionPool interface {
	IsReady() <-chan struct{}
	Get() net.Conn
	Release(conn net.Conn)
}

type Sender struct {
	connectionPool ConnectionPool

	initPoolOnce sync.Once
}

func (l *Sender) Send(ctx context.Context, payload []byte) error {
	pool := l.getConnectionPool()

	// ensure pool is ready
	select {
	case <-pool.IsReady():
		// happy path

	case <-ctx.Done():
		// context timed out or was cancelled
		return errors.New("failed to get connection")
	}

	// get connection from pool and return afterwards
	conn := pool.Get()
	defer l.connectionPool.Release(conn)

	// send and return
	_, err := conn.Write(payload)

	return err
}

func (l *Sender) getConnectionPool() ConnectionPool {
	// Inject the connection pool with JIT DI
	if l.connectionPool == nil {
		myPool := &myConnectionPool{}
		go myPool.init()

		l.connectionPool = myPool
	}

	return l.connectionPool
}

// default implementation of the connection pool
type myConnectionPool struct {
}

// IsReady implements ConnectionPool
func (m *myConnectionPool) IsReady() <-chan struct{} {
	// not implemented yet
	return make(chan struct{})
}

// IsReady implements ConnectionPool
func (m *myConnectionPool) Get() net.Conn {
	// not implemented yet
	return nil
}

// IsReady implements ConnectionPool
func (m *myConnectionPool) Release(_ net.Conn) {
	// not implemented yet
}

func (m *myConnectionPool) init() {
	// create connection and populate the pool
}
