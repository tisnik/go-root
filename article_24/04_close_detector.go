package main

import (
	"net/http"
	"time"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	println("handler started")
	for i := 0; i < 50; i++ {
		n, err := writer.Write([]byte("Hello world!\n"))
		if err != nil {
			println("\nI/O error:", err.Error())
			return
		}
		if n == 0 {
			println("\nnothing written")
			return
		}
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
