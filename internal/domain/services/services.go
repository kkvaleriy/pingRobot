package services

import (
	"sync"
	"time"
)

var (
	once sync.Once
	s    *services
)

type services struct {
	store map[string]Service
	mu    *sync.Mutex
}

func Set(serviceName, status string, timeOfCheck time.Time) {
	once.Do(
		func() {
			s.store = make(map[string]Service)
			s.mu = &sync.Mutex{}
		})

	serv := Service{
		status:    status,
		lastCheck: timeOfCheck,
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[serviceName] = serv
}

func Status() map[string]Service {
	r := make(map[string]Service)
	s.mu.Lock()
	defer s.mu.Unlock()

	for name, serv := range s.store {
		r[name] = Service{
			status:    serv.status,
			lastCheck: serv.lastCheck,
		}
	}

	return r
}
