package helpers

import (
	"sync"
)

type handleFuncWorker = func(...any)

type WorkerJob struct {
	Cancel     bool
	HandleFunc handleFuncWorker
	Params     []any
}

type WorkerPool struct {
	jobs       []WorkerJob
	maxWorkers uint
	wg         *sync.WaitGroup
	mut        *sync.Mutex
}

func NewWorker(workerPool *WorkerPool, wg *sync.WaitGroup, id uint) {
	go func() {
		for {
			workerPool.mut.Lock()
			if len(workerPool.jobs) > 0 {
				first := workerPool.jobs[0]
				workerPool.jobs = workerPool.jobs[1:]
				if first.Cancel {
					wg.Done()
					workerPool.mut.Unlock()
					break
				}
				workerPool.mut.Unlock()
				first.HandleFunc(first.Params...)
				continue //This continue statement avoid the execute the second Unlock :!
			}
			workerPool.mut.Unlock()

		}
	}()
}

func NewWorkerPool(size uint) *WorkerPool {

	mut := new(sync.Mutex)
	wg := new(sync.WaitGroup)

	workerPool := &WorkerPool{
		wg:         wg,
		mut:        mut,
		jobs:       []WorkerJob{},
		maxWorkers: size,
	}

	wg.Add(int(size))
	for i := 0; i < int(size); i++ {
		NewWorker(workerPool, wg, uint(i))
	}
	return workerPool
}

func (w *WorkerPool) Execute(f handleFuncWorker, p ...any) {
	w.mut.Lock()
	w.jobs = append(w.jobs, WorkerJob{
		Cancel:     false,
		HandleFunc: f,
		Params:     p,
	})
	w.mut.Unlock()
}

func (w *WorkerPool) Wait() {
	defer w.wg.Wait()
	jobs := []WorkerJob{}
	for i := 0; i < int(w.maxWorkers); i++ {
		jobs = append(jobs, WorkerJob{
			Cancel: true,
			HandleFunc: func(p ...any) {
			},
			Params: []any{},
		})
	}
	w.mut.Lock()
	w.jobs = append(w.jobs, jobs...)
	w.mut.Unlock()

}
