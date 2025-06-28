package services

import (
	"sync"
)

type services struct {
	store map[string]service
	mu    *sync.Mutex
}
}
