package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type FactorialPageDynContent struct {
	Result int
}

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("factorial.html")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprint(writer, "Not found!")
		return
	}

	dynData := FactorialPageDynContent{1}
	err = t.Execute(writer, dynData)
	if err != nil {
		println("Error executing template")
	}
}

func main() {
	http.HandleFunc("/", mainEndpoint)
	http.ListenAndServe(":8000", nil)
}
