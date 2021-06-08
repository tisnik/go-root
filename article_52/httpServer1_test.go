// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá druhá část
//    Pomůcky při tvorbě jednotkových testů v jazyce Go
//    https://www.root.cz/clanky/pomucky-pri-tvorbe-jednotkovych-testu-v-jazyce-go/

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
