package gatewaysrv

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/koen-or-nant/go-notification-service/internal/config"
)

type Server struct {
	srv *http.Server
}

func NewServer() Server {
	r := mux.NewRouter()
	r.HandleFunc("/email", email).Methods("POST")
	r.HandleFunc("/sms", sms).Methods("POST")
	if port, exist := config.GetConfig("PORT"); exist {
		address := fmt.Sprintf("0.0.0.0:%s", port)
		log.Println("listening on", address)
		return Server{
			srv: &http.Server{
				Addr: address,
				// Good practice to set timeouts to avoid Slowloris attacks.
				WriteTimeout: time.Second * 15,
				ReadTimeout:  time.Second * 15,
				IdleTimeout:  time.Second * 60,
				Handler:      r,
			},
		}
	}
	return Server{}
}

func (s Server) Run() {
	if s.srv.Addr == "" {
		log.Println("ERROR: unable to run server")
		return
	}
	go func() {
		if err := s.srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	s.srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
}
