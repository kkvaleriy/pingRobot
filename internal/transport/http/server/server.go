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

func initServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/metrics", handleMetrics)

	srv = &http.Server{
		Addr:    ":80",
		Handler: mux,
	}

	err := srv.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatalf("server failed to start: %s", err.Error())
	}
}

func Run(ctx context.Context) {
	once.Do(
		func() {
			go initServer()

			<-ctx.Done()

			c, _ := context.WithTimeout(context.Background(), time.Minute)
			err := srv.Shutdown(c)
			if err != nil {
				log.Printf("server failed to shutdown: %s", err.Error())
				return
			}
			log.Print("server shutdown successfuly")
		})
}

func handleMetrics(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		metrics(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func metrics(w http.ResponseWriter, r *http.Request) {
	b, err := CreateJson(status.Check())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Write(b)
}
