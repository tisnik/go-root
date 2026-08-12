package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const Port = 8080

const HttpRequestTimeout = 5 * time.Second
const PageHeader = `
<!doctype html>
<html>
    <head>
        <title>ε</title>
        <script src="fengari-web.js"></script>
    </head>
    <body>
`

const PageFooter = `
        <canvas id="canvas" width="400" height="400"></canvas>
        <script src="canvas.lua" type="text/lua">
        </script>
    </body>
</html>
`

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

	http.HandleFunc("/", s.mainEndpoint)

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

func (s ServerImpl) mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	}
	io.WriteString(writer, PageHeader)
	io.WriteString(writer, "        <table>\n")
	io.WriteString(writer, "            <tr><th>&nbsp;</th>")
	for column := 'A'; column <= MaxColumns; column++ {
		io.WriteString(writer, fmt.Sprintf("<th>%c</th>", column))

	}
	io.WriteString(writer, "</tr>\n")
	for row := 1; row <= MaxRows; row++ {
		io.WriteString(writer, fmt.Sprintf("            <tr><th>%d</th>", row))
		for column := 'A'; column <= MaxColumns; column++ {
			cellID := fmt.Sprintf("%c%d", column, row)
			io.WriteString(writer, fmt.Sprintf("<td id='%s'></td>", cellID))
		}
		io.WriteString(writer, "</tr>\n")
	}
	io.WriteString(writer, "        </table>\n")
	io.WriteString(writer, PageFooter)
}

func main() {
	storage := NewStorage()
	server := NewServer(storage)
	server.Serve(Port)
}
