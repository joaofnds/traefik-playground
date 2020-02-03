package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set")
	}

	http.HandleFunc("/", rootHandler)

	addr := ":" + port
	server := http.Server{Addr: addr}

	go (func() {
		log.Fatal(server.ListenAndServe())
	})()

	log.Printf("Server up and running at %s\n", addr)
	blockUntilSigTerm()
	log.Printf("So long friend, goodbye...")

	server.Shutdown(context.Background())
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	b, err := httputil.DumpRequest(r, true)
	if err == nil {
		w.WriteHeader(200)
		w.Write(b)
	} else {
		w.WriteHeader(500)
		w.Write([]byte("failed to dump request"))
	}
}

func blockUntilSigTerm() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
