package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"dev11/pkg/delivery/http"
	"dev11/pkg/service"
)

const serverPort = "8080"

func main() {
	s := service.NewService()
	h := http.NewHandler(s)

	srv := new(http.Server)
	go func() {
		if err := srv.Run(serverPort, h.InitRoutes()); err != nil {
			log.Fatal(err)
		}
	}()
	log.Print("Service is successfully started...")

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("Service is shutting down...")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down: %s", err.Error())
	}
}
