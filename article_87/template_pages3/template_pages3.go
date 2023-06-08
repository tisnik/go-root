package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fsnotify/fsnotify"
)

const templateFilename = "./html_template06.htm"

// datový typ, jehož prvky budou vypisovány v šabloně
type Role struct {
	Name       string
	Surname    string
	Popularity int
}

func sendStaticPage(writer http.ResponseWriter, filename string) {
	log.Printf("Sending static file %s", filename)

	body, err := ioutil.ReadFile(filename)
	if err == nil {
		fmt.Fprint(writer, string(body))
	} else {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprint(writer, "Not found!")
	}
}

func staticPage(filename string) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		sendStaticPage(writer, filename)
	}
}

func rolesHandler(tmpl *template.Template, roles []Role, changed chan string) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		select {
		case filename := <-changed:
			log.Printf("Have to reload template from file %s", filename)
			tmpl = loadTemplate(filename)
		default:
			log.Print("Using old template")
		}

		log.Printf("Application template for %d data records", len(roles))
		// aplikace šablony - přepis hodnot
		err := tmpl.Execute(writer, roles)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error executing template: %v", err)
			return
		}
	}
}

func loadTemplate(templateFilename string) *template.Template {
	log.Printf("Constructing template from file %s", templateFilename)

	// vytvoření nové šablony
	tmpl, err := template.ParseFiles(templateFilename)
	if err != nil {
		log.Fatalf("Template can't be constructed: %v", err)
		return nil
	}

	return tmpl
}

func startWatcher(directory string, filename string, changed chan string) {
	log.Print("Starting watcher")

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	log.Printf("Watching directory %s", directory)

	err = watcher.Add(directory)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			log.Println("event:", event)
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Println("modified file:", event.Name)
				if filename == event.Name {
					log.Println("Template change detected")
					changed <- filename
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}

func main() {
	// tyto hodnoty budou použity při aplikaci šablony
	roles := []Role{
		{"Eliška", "Najbrtová", 4},
		{"Jenny", "Suk", 3},
		{"Anička", "Šafářová", 0},
		{"Sváťa", "Pulec", 3},
		{"Blažej", "Motyčka", 8},
		{"Eda", "Wasserfall", 0},
		{"Přemysl", "Hájek", 10},
	}

	changed := make(chan string)
	go startWatcher(".", templateFilename, changed)

	tmpl := loadTemplate(templateFilename)

	const address = ":8080"

	log.Printf("Starting server on address %s", address)

	http.HandleFunc("/", staticPage("index.html"))
	http.HandleFunc("/missing", staticPage("missing.html"))
	http.HandleFunc("/roles", rolesHandler(tmpl, roles, changed))
	http.ListenAndServe(address, nil)
}
