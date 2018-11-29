package disadvantages

func DoJobUpdated(pool WorkerPool, job Job) error {
	worker := pool.GetWorker()
	return worker.Do(job)
}
