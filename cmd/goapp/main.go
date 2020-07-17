package main

import (
	"goapp/internal/app/goapp/handlers"
	"goapp/internal/pkg/server"
	"goapp/pkg/env"
	"goapp/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

const listenAddress env.EnvironmentVariable = "LISTEN_ADDRESS"

func main() {
	gracefulInterrupt()

	address := listenAddress.GetOr(":8080")

	log := logger.New(logger.Info)

	srv := server.NewServer(log)

	handlers.Register(srv, log)

	if err := srv.Run(address); err != nil {
		log.Fatal("server stopped", err)
		exit()
	}
}

func gracefulInterrupt() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		// Run cleanup functions here
		exit()
	}()
}

func exit() {
	os.Exit(0)
}
