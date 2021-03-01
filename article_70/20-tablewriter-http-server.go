package main

import (
	"net/http"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

// Factorial computes factorial for given n that might be positive integer
func Factorial(n int64) int64 {
	switch {
	case n < 0:
		return 1
	case n == 0:
		return 1
	default:
		return n * Factorial(n-1)
	}
}

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	table := tablewriter.NewWriter(writer)
	table.SetHeader([]string{"n", "n!"})

	for n := 0; n <= 20; n++ {
		f := Factorial(int64(n))
		row := []string{strconv.Itoa(n), strconv.FormatInt(f, 10)}
		table.Append(row)
	}

	table.Render()
}

func main() {
	http.HandleFunc("/", mainEndpoint)
	http.ListenAndServe(":8000", nil)
}
