package main

import (
	"log"
	"os"

	"github.com/koen-or-nant/go-notification-service/internal/config"
	"github.com/koen-or-nant/go-notification-service/internal/creatorsrv"
	"github.com/koen-or-nant/go-notification-service/pkg/api"
	"github.com/koen-or-nant/go-notification-service/pkg/creator"
	"github.com/koen-or-nant/go-notification-service/pkg/dispatcher"
	"github.com/koen-or-nant/go-notification-service/pkg/types"
)

func main() {
	parallelRequestsNo, exist := config.GetConfigAsInt("PARALLEL_REQUESTS_NO")
	if !exist {
		log.Println("PARALLEL_REQUESTS_NO env var must be set")
		os.Exit(1)
	}
	in := make(chan api.Notification, parallelRequestsNo)
	defer close(in)
	out := make(chan types.Sendable, parallelRequestsNo)
	defer close(out)
	service := creator.NewNotificationCreator(in, out)
	go func() {
		service.Run()
	}()
	dispatchService := dispatcher.NewDispatcher(out)
	go func() {
		dispatchService.Run()
	}()
	server := creatorsrv.NewServer(
		creatorsrv.NewNotificationServer(in, out))
	server.Run()
	log.Println("service is closed")
}
