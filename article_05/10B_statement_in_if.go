// Seriál "Programovací jazyk Go"
//
// Pátá část
//
// Příkaz zapsaný za klíčovým slovem if.

package main

func funkce() int {
	return -1
}

func x() string {
	if value := funkce(); value < 0 {
		return "záporná hodnota"
	} else if value > 0 {
		return "kladná hodnota"
	} else {
		return "nula"
	}
}

func main() {
	println(x())
}
