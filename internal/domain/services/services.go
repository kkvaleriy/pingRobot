package services

import (
	"sync"
	"time"
)
)

type services struct {
	store map[string]service
	mu    *sync.Mutex
}

func Set(serviceName, status string, timeOfCheck time.Time) {
	once.Do(
		func() {
			s.store = make(map[string]service)
			s.mu = &sync.Mutex{}
		})

	serv := service{
		status:    status,
		lastCheck: timeOfCheck,
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[serviceName] = serv
}
