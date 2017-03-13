package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cloud.google.com/go/trace"
)

var (
	httpAddr    string
	projectID   string
	databaseDir string
	traceClient *trace.Client
)

func main() {
	flag.StringVar(&httpAddr, "http", "0.0.0.0:8081", "HTTP Service Address")
	flag.StringVar(&projectID, "project-id", "", "App Engine Project ID")
	flag.StringVar(&databaseDir, "database-dir", "", "Application Database")
	flag.Parse()
	log.Println("HTTP service listening on ", httpAddr)

	var err error
	ctx := context.Background()

	traceClient, err = trace.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}

	p, err := trace.NewLimitedSampler(1, 10)
	if err != nil {
		log.Fatal(err)
	}
	traceClient.SetSamplingPolicy(p)

	router := NewRouter()
	server := &http.Server{Addr: httpAddr, Handler: router}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	s := <-signalChan

	log.Println(fmt.Sprintf("Captured %v, exiting...", s))
	shutdownCtx, _ := context.WithTimeout(context.Background(), 120*time.Second)
	server.Shutdown(shutdownCtx)
	<-shutdownCtx.Done()

	log.Println(shutdownCtx.Err())
	os.Exit(0)
}
