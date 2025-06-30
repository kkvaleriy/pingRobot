package server

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/kkvaleriy/pingRobot/internal/usecase/status"
)

var (
	once sync.Once = sync.Once{}
	srv  *http.Server
)

func metrics(w http.ResponseWriter, r *http.Request) {
	b, err := CreateJson(status.Check())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Write(b)
}
