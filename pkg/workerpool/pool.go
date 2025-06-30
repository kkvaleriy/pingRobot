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
