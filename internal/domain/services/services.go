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
}
