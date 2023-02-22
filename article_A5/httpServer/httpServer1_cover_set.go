//line httpServer1.go:1
// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Část číslo 105:
//    Nová funkcionalita v Go 1.20: detekce skutečně volaných řádků v programovém kódu
//
// Seznam příkladů z části číslo 105:
//    https://github.com/tisnik/go-root/blob/master/article_A5/README.md

package main

import (
	"fmt"
	"log"
	"net/http"
)

func dataHandler(writer http.ResponseWriter, request *http.Request) {coverageVariable.Count[0] = 1;
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, `"x": [1, 2, 3, 4, 5]`)
}

func otherHandler(writer http.ResponseWriter, request *http.Request) {coverageVariable.Count[1] = 1;
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, `foobar`)
}

func startHttpServer(address string) {coverageVariable.Count[2] = 1;
	log.Printf("Starting server on address %s", address)
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/data", dataHandler)
	http.HandleFunc("/other", otherHandler)
	http.ListenAndServe(address, nil)
}

func main() {coverageVariable.Count[3] = 1;
	startHttpServer(":8080")
}

var coverageVariable = struct {
	Count     [4]uint32
	Pos       [3 * 4]uint32
	NumStmt   [4]uint16
} {
	Pos: [3 * 4]uint32{
		18, 22, 0x20045, // [0]
		24, 28, 0x20046, // [1]
		30, 36, 0x20026, // [2]
		38, 40, 0x2000d, // [3]
	},
	NumStmt: [4]uint16{
		3, // 0
		3, // 1
		5, // 2
		1, // 3
	},
}
