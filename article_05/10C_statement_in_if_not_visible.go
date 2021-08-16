// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
//
// Příkaz zapsaný za klíčovým slovem if, hodnota není viditelná za if.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_05/10C_statement_in_if_not_visible.html

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
