package disadvantages

import (
	"errors"
	"time"
)

func DoJob(pool WorkerPool, job Job) error {
	// wait for pool
	ready := pool.IsReady()

	select {
	case <-ready:
		// happy path

	case <-time.After(1 * time.Second):
		return errors.New("timeout waiting for worker pool")
	}

	worker := pool.GetWorker()
	return worker.Do(job)
}

// Pool of workers
type WorkerPool interface {
	GetWorker() Worker
	IsReady() chan struct{}
}

// Executes/processes a unit of work and returns
type Worker interface {
	Do(job Job) error
}

// A unit of work to be executed against the pool
type Job interface {
	// implementation omitted
}
