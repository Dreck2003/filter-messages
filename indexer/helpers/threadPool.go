package helpers

import (
	"sync"
)

type handleFunc = func(p ...any)

type Job struct {
	Cancel     bool
	HandleFunc handleFunc
	Params     []any
}

func CreateWorker(id uint, c chan Job, wg *sync.WaitGroup) {
	go func(id uint, c chan Job, wg *sync.WaitGroup) {
		for {
			job := <-c
			if job.Cancel {
				wg.Done()
				break
			}
			job.HandleFunc(job.Params...)
		}
	}(id, c, wg)
}

type ThreadPool struct {
	channel chan Job
	wg      *sync.WaitGroup
	maxSize uint
}

func NewThreadPool(size uint) *ThreadPool {

	channel := make(chan Job)
	wg := new(sync.WaitGroup)

	for i := 0; i < int(size); i++ {
		wg.Add(1)
		CreateWorker(uint(i), channel, wg)
	}

	return &ThreadPool{
		channel: channel,
		wg:      wg,
		maxSize: size,
	}
}

func (t *ThreadPool) Execute(f handleFunc, params ...any) {
	t.channel <- Job{
		Cancel:     false,
		HandleFunc: f,
		Params:     params,
	}
}

func (t *ThreadPool) Wait() {
	defer t.wg.Wait()
	for i := 0; i < int(t.maxSize); i++ {
		t.channel <- Job{
			Cancel: true,
			HandleFunc: func(p ...any) {
			},
			Params: []any{},
		}
	}

}
