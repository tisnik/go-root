// Seriál "Programovací jazyk Go"
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
//
// Příkaz zapsaný za klíčovým slovem if, hodnota není viditelná za if.

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
	println(value)
	return ""
}

func main() {
	println(x())
}
