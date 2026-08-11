package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const Port = 8080

const HttpRequestTimeout = 5 * time.Second

type Storage interface {
}

type MemoryStorage struct {
}

// NewStorage creates and returns an in-memory storage implementation.
func NewStorage() Storage {
	return MemoryStorage{}
}

type Server interface {
	Serve(port uint)
}

// ServerImpl is a simple HTTP server implementation
type ServerImpl struct {
	storage Storage
}

// NewServer creates a server backed by the provided storage.
func NewServer(storage Storage) Server {
	return ServerImpl{
		storage: storage,
	}
}

func (s ServerImpl) Serve(port uint) {
	log.Printf("Starting server on port %d", port)

	// start the server
	httpServer := &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           http.DefaultServeMux,
		ReadHeaderTimeout: HttpRequestTimeout,
	}
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	storage := NewStorage()
	server := NewServer(storage)
	server.Serve(Port)
}
