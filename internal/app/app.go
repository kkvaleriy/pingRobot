package app

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/kkvaleriy/pingRobot/internal/config"
	"github.com/kkvaleriy/pingRobot/internal/transport/http/ping"
	"github.com/kkvaleriy/pingRobot/internal/transport/http/server"
	"github.com/kkvaleriy/pingRobot/pkg/workerpool"
)

func Run() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	gracefulShutdownWG := &sync.WaitGroup{}
	gracefulShutdownWG.Add(1)

	cfg := config.FromFile()
	rootCtx, stopRootCtx := context.WithCancel(context.Background())
	
	server.Run(rootCtx, cfg)
	
	wp, workerChannel := workerpool.New(ping.Do, 5)
	go wp.Run(rootCtx)
	
	go SendServiceToWorker(rootCtx, workerChannel, *cfg)

	go func(){
		<- sigChan
		stopRootCtx()
		<- time.After(time.Second * 5)
		gracefulShutdownWG.Done()
	}()
	
	gracefulShutdownWG.Wait()

}

func SendServiceToWorker(ctx context.Context,workerChannel chan workerpool.Job, cfg config.Config) {
	servicesFromCfg := cfg.ServicesForCheck()

	servicesForWorkerpool := []workerpool.Job{}

	for _, v := range servicesFromCfg {
		servicesForWorkerpool = append(servicesForWorkerpool,
			 workerpool.Job{ServiceName: v.Name, ServiceEndpoint: v.Url})
	}


	for {
		for _, v := range servicesForWorkerpool {
			workerChannel <- v
		}
		select {
		case <- ctx.Done():
			return
		case <-time.After(time.Second * 10):
		}
		
	}
}