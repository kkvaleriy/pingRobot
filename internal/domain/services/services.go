package services

import (
	"sync"
)

type services struct {
	store map[string]service
	wg    *sync.Mutex
	once  *sync.Once
}
