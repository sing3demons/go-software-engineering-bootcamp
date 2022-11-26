package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`hello`))
	})

	srv := http.Server{
		Addr:    ":2565",
		Handler: mux,
	}

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	fmt.Println("server starting at :2565")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	<-shutdown
	fmt.Println("shutting down...")
	if err := srv.Shutdown(context.Background()); err != nil {
		fmt.Println("shutdown err:", err)
	}
	fmt.Println("server stop")
}
