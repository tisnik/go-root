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
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDataHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "/data", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(dataHandler)

	handler.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("improper status code: got %v instead of %v",
			status, http.StatusOK)
	}

	body := recorder.Body.String()
	if body != `"x": [1, 2, 3, 4, 5]` {
		t.Errorf("wrong response body: %s", body)
	}

	if ctype := recorder.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("content type header does not match: got %s want %s",
			ctype, "application/json")
	}
}
