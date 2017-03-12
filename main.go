package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	httpAddr string
)

func main() {
	flag.StringVar(&httpAddr, "http", "0.0.0.0:8081", "HTTP Service Address")
	flag.Parse()
	log.Println("HTTP service listening on ", httpAddr)

	router := NewRouter()
	server := &http.Server{Addr: httpAddr, Handler: router}
	go func() {
		log.Fatal(server.ListenAndServe())
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown signal received, exiting...")
}
