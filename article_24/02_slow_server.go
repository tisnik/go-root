// Seriál "Programovací jazyk Go"
//
// Dvacátá čtvrtá část
//
// Demonstrační příklad číslo 2:
//     	Zpomalení generování jednotlivých bloků generovaného obsahu

package main

import (
	"net/http"
	"time"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	println("handler started")
	for i := 0; i < 50; i++ {
		writer.Write([]byte("Hello world!\n"))
		time.Sleep(100 * time.Millisecond)
		print(".")
	}
	println("\nhandler finished")
}

func main() {
	http.HandleFunc("/data", handler)
	http.ListenAndServe(":8080", nil)
}
