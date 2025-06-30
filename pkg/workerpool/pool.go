package workerpool

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

type pool struct {
	worker       worker
	jobs         chan Job
	workersCount atomic.Int32
	maxWorkers   int
	wg           *sync.WaitGroup
}

func New(f func(serviceEndpoint, serviceName string), maxWorkers int) (*pool, chan Job) {

	jc := make(chan Job)

	return &pool{
		worker: worker{f: f},
		jobs: jc,
		maxWorkers: maxWorkers,
		wg: &sync.WaitGroup{},
	}, jc
}

func (w *pool)Run(ctx context.Context) {
	defer w.wg.Wait()

	for {
		select {
		case <- ctx.Done():
			return
		default:
		}

		currentWorkers := w.workersCount.Load()
		if int(currentWorkers) >= w.maxWorkers {
			<-time.After(time.Second)
			continue
		}
		
		go w.startJob()
		
	}
}

func (w *pool) startJob() {
	w.wg.Add(1)
	w.workersCount.Add(1)
	
	j := <- w.jobs
	w.worker.f(j.ServiceEndpoint, j.ServiceName)

	w.wg.Done()
	w.workersCount.Add(-1)
}