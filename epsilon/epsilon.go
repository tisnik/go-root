// Backend for the Epsilon project
package main

import (
	"embed"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// default server port
const Port = 8080

const HttpRequestTimeout = 5 * time.Second
const MaxRows = 99
const MaxColumns = 'Z'

// ---------------------------------------------------------------------------
// New data types
// ---------------------------------------------------------------------------

type Storage interface {
}

type MemoryStorage struct {
}

type Configuration struct {
	port        uint `json:"port"`
	showHelp    bool
	showVersion bool
}

// application holds the dependencies for our web application.
type Application struct {
	config Configuration
	logger *log.Logger
}

// Server interface
type Server interface {
	Serve()
}

// ServerImpl is a simple HTTP server implementation
type ServerImpl struct {
	app *Application
}

// ---------------------------------------------------------------------------
// Resources embedded into the final binary file
// ---------------------------------------------------------------------------

// HTML pages, templates, and part of pages
//
//go:embed html/page_header.htm
var PageHeader string

//
//go:embed html/page_footer.htm
var PageFooter string

// Static images
//
//go:embed img/*.png
var StaticImages embed.FS

//go:emded img/favicon.ico
var Favicon []byte

// Stylesheets
//
//go:embed css/*.css
var StaticStylesheets embed.FS

// Scripts
//
//go:embed scripts/fengari-web.js
var FengariWebJS string

//go:embed scripts/*.lua
var Scripts embed.FS

// ---------------------------------------------------------------------------
// Interface implementations
// ---------------------------------------------------------------------------

// NewStorage creates and returns an in-memory storage implementation.
func NewStorage() Storage {
	return MemoryStorage{}
}

// ---------------------------------------------------------------------------
// HTTP server implementations
// ---------------------------------------------------------------------------

// NewServer creates a server for the provided application.
func NewServer(app *Application) Server {
	return ServerImpl{
		app: app,
	}
}

// serveStaticImage serves a PNG image from the embedded static image bundle.
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

// serveStaticStylesheet serves a CSS stylesheet from the embedded static
// assets.
//
// It maps the request URL path to a file under the embedded css/
// directory, reads the file contents, and returns them with a text/css
// content type.
//
// If the stylesheet cannot be found, it logs the error and
// responds with 404 Not Found.
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

// serveLuaInterpreter serves the embedded JavaScript runtime needed to run Lua
// interpreter in the web browser.
//
// It responds with the Fengari web bundle and sets the content type to
// application/javascript.
func (s ServerImpl) serveLuaInterpreter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	io.WriteString(w, FengariWebJS)
}

// serveFavicon serves the application's favicon as an ICO image.
//
// It sets the response content type to image/x-icon and writes the embedded
// icon bytes.
func (s ServerImpl) serveFavicon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/x-icon")
	w.Write(Favicon)
}

// serveScripts serves a JavaScript file from the embedded scripts directory.
//
// It maps the request path to an embedded file, returns the file contents,
// and responds with 404 if the script does not exist.
func (s ServerImpl) serveScripts(w http.ResponseWriter, r *http.Request) {
	path := r.URL.String()
	scriptFileName := "scripts/" + strings.TrimPrefix(path, "/scripts/")
	data, err := Scripts.ReadFile(scriptFileName)
	if err != nil {
		log.Print(err)
		http.NotFound(w, r)
		return
	}
	w.Write(data)
}

func (s ServerImpl) renderTable(writer http.ResponseWriter) {
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
}

// mainEndpoint serves the main HTML page at the root path.
//
// It returns 404 for any non-root request, writes the page header and footer,
// and renders the main table content in between.
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

func (s ServerImpl) Serve() {
	port := s.app.config.port
	log.Printf("Starting server on port %d", port)

	http.HandleFunc("/", s.mainEndpoint)

	// static content
	http.HandleFunc("/favicon.ico", s.serveFavicon)
	http.HandleFunc("/fengari-web.js", s.serveLuaInterpreter)

	// scripts
	http.HandleFunc("/scripts/{path}", s.serveScripts)

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
		ReadTimeout:       HttpRequestTimeout,
		WriteTimeout:      HttpRequestTimeout,
		IdleTimeout:       HttpRequestTimeout,
	}
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

// ---------------------------------------------------------------------------
// Import/export
// ---------------------------------------------------------------------------

// main parses command-line options, configures the application, and starts the HTTP server.
// It prints help or version information when requested and reports invalid port values.
func main() {
	var cfg Configuration

	// Read env vars (if any)
	defaultPort := os.Getenv("PORT")
	if defaultPort == "" {
		defaultPort = strconv.Itoa(Port)
	}
	defaultPortNumber, err := strconv.Atoi(defaultPort)
	if err != nil {
		fmt.Printf("Improper port value: %s\n", defaultPort)
		return
	}

	flag.BoolVar(&cfg.showHelp, "h", false, "display help")
	flag.BoolVar(&cfg.showVersion, "v", false, "display version")
	flag.UintVar(&cfg.port, "p", uint(defaultPortNumber), "port for the server (shorthand)")
	flag.UintVar(&cfg.port, "port", uint(defaultPortNumber), "port for the server")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &Application{
		config: cfg,
		logger: logger,
	}

	switch {
	case cfg.showHelp:
		flag.PrintDefaults()
	case cfg.showVersion:
		fmt.Println("version")
	default:
		// default operation: start the HTTP server
		//storage := NewStorage()
		server := NewServer(app)
		server.Serve()
	}

}
