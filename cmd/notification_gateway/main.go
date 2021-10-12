package main

import (
	"log"

	"github.com/koen-or-nant/go-notification-service/internal/gatewaysrv"
)

func main() {
	server := gatewaysrv.NewServer()
	server.Run()
	log.Println("service is closed")
}
