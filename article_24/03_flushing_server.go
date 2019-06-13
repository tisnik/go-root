// Seriál "Programovací jazyk Go"
//
// Dvacátá čtvrtá část
//
// Demonstrační příklad číslo 3:
//     	Využití metody Flush z rozhraní Flusher

package main

import (
	"net/http"
	"time"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	println("handler started")
	for i := 0; i < 50; i++ {
		writer.Write([]byte("Hello world!\n"))
		print(".")
		if f, ok := writer.(http.Flusher); ok {
			f.Flush()
		}
		time.Sleep(100 * time.Millisecond)
	}
	println("\nhandler finished")
}

func main() {
	http.HandleFunc("/data", handler)
	http.ListenAndServe(":8080", nil)
}
