package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

const templateFilename = "./html_template06.htm"

func startWatcher(directory string, filename string) {
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
	startWatcher(".", templateFilename)
}
