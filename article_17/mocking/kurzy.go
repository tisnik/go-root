// Seriál "Programovací jazyk Go"
//
// Sedmnáctá část
//     Testování aplikací naprogramovaných v jazyce Go
//     https://www.root.cz/clanky/testovani-aplikaci-naprogramovanych-v-jazyce-go/
//
// Přístup k tabulce kurzů a použití kurzů při výpočtu převodu měn.

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type ExchangeDataGetter func(code string) float64

type ExchangeGetter struct {
	get_exchange_rate ExchangeDataGetter
}

func NewExchangeGetter(g ExchangeDataGetter) *ExchangeGetter {
	return &ExchangeGetter{get_exchange_rate: g}
}

func get_exchange_rate_from_url(code string) float64 {
	const URL = "https://www.cnb.cz/cs/financni_trhy/devizovy_trh/kurzy_devizoveho_trhu/denni_kurz.txt"

	response, err := http.Get(URL)
	if err != nil {
		panic("Connection refused")
	}
	defer response.Body.Close()

	fmt.Printf("Status: %s\n", response.Status)
	fmt.Printf("Content length: %d\n", response.ContentLength)

	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "|")
		if len(s) == 5 {
			code_str := s[3]
			rate_str := strings.Replace(s[4], ",", ".", 1)
			if code == code_str {
				rate_f, err := strconv.ParseFloat(rate_str, 64)
				if err != nil {
					panic(err)
				}
				return rate_f
			}
		}
	}

	return 0
}

func get_exchange_rate_from_file(code string) float64 {
	const FILENAME = "kurzy.txt"

	file, err := os.Open(FILENAME)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "|")
		if len(s) == 5 {
			code_str := s[3]
			rate_str := strings.Replace(s[4], ",", ".", 1)
			if code == code_str {
				rate_f, err := strconv.ParseFloat(rate_str, 64)
				if err != nil {
					panic(err)
				}
				return rate_f
			}
		}
	}

	return 0
}

func mocked_get_exchange_rate(code string) float64 {
	return 21.5
}

func (g *ExchangeGetter) exchange(amount float64, code string) float64 {
	rate := g.get_exchange_rate(code)
	return rate * amount
}

func main() {
	g := NewExchangeGetter(get_exchange_rate_from_file)
	fmt.Printf("%5.2f\n", g.exchange(10, "USD"))
	g2 := NewExchangeGetter(mocked_get_exchange_rate)
	fmt.Printf("%5.2f\n", g2.exchange(10, "USD"))
}
