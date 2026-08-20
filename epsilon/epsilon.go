package main

import (
	"embed"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const Port = 8080

const HttpRequestTimeout = 5 * time.Second
const MaxRows = 99
const MaxColumns = 'Z'

const PageHeader = `
<!doctype html>
<html>
    <head>
        <title>ε</title>
	<link rel="stylesheet" href="css/main.css">
        <script src="fengari-web.js"></script>
    </head>
    <body>
        <div id="toolbar">
            <button><img src="/image/edit-copy.png" /></button>
            <button><img src='/image/edit-paste.png' /></button>
            <button><img src='/image/edit-delete.png' /></button>
            &nbsp;&nbsp;&nbsp;
            <button><img src='/image/help-about.png' /></button>
            <button><img src='/image/help-faq.png' /></button>
        </div>
	<div id="inputbar">
	    <input type="text" id="input" name="input" />
        </div>
`

const PageFooter = `
        <canvas id="canvas" width="400" height="400"></canvas>
        <script src="canvas.lua" type="text/lua">
        </script>
    </body>
</html>
`

// ---------------------------------------------------------------------------
// Resources embedded into the final binary file
// ---------------------------------------------------------------------------

// Static images
//
//go:embed img/*.png
var StaticImages embed.FS

// Stylesheets
//
//go:embed css/*.css
var StaticStylesheets embed.FS

// ---------------------------------------------------------------------------
// New data types
// ---------------------------------------------------------------------------

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

// serveStaticImage serves a PNG image from the embedded static image bundle
//
// It maps requests under /image/ to files in the img/ directory, reads the
// requested file from StaticImages, and writes it to the response with a
// Content-Type of image/png.
//
// If the image cannot be found or read, it logs the error and returns a 404
// Not Found.
func (s ServerImpl) serveStaticImage(w http.ResponseWriter, r *http.Request) {
	path := r.URL.String()

	// construct proper image name from provided path
	imageName := "img/" + strings.TrimPrefix(path, "/image/")

	// read binary data bundled together with the application
	binaryData, err := StaticImages.ReadFile(imageName)
	if err != nil {
		log.Print(err)
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image/png")
	w.Write(binaryData)
}

func (s ServerImpl) serveStaticStylesheet(w http.ResponseWriter, r *http.Request) {
	path := r.URL.String()
	styleSheetName := "css/" + strings.TrimPrefix(path, "/css/")
	data, err := StaticStylesheets.ReadFile(styleSheetName)
	if err != nil {
		log.Print(err)
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/css")
	w.Write(data)
}

func (s ServerImpl) renderTable(writer http.ResponseWriter) {
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
}

func (s ServerImpl) mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	}
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(writer, PageHeader)
	s.renderTable(writer)
	io.WriteString(writer, PageFooter)
}

func (s ServerImpl) serveLuaInterpreter(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "fengari-web.js")
}

func (s ServerImpl) serveCanvas(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "canvas.lua")
}

func (s ServerImpl) returnCell(writer http.ResponseWriter, r *http.Request) {
	ID := r.PathValue("id")
	log.Printf("Cell ID provided: %s", ID)

	/*
		ID, err := strconv.Atoi(IDs)
		if err != nil {
			writer.Header().Set("Content-Type", "text/plain")
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

			user, found := s.storage.ReadUser(ID)
			if !found {
				writer.Header().Set("Content-Type", "text/plain")
				writer.WriteHeader(http.StatusNotFound)
				return
			}*/
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(ID)
}

func (s ServerImpl) Serve(port uint) {
	log.Printf("Starting server on port %d", port)

	http.HandleFunc("/", s.mainEndpoint)

	// static content
	http.HandleFunc("/fengari-web.js", s.serveLuaInterpreter)
	http.HandleFunc("/canvas.lua", s.serveCanvas)

	// images
	http.HandleFunc("/image/{path}", s.serveStaticImage)
	// stylesheets
	http.HandleFunc("/css/{path}", s.serveStaticStylesheet)

	// REST API endpoints
	http.HandleFunc("GET /cell/{id}", s.returnCell)

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
