// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
//
// Seznam příkladů z páté části:
//    https://github.com/tisnik/go-root/blob/master/article_05/README.md
//
// Demonstrační příklad číslo 26:
//    For s range prováděným nad mapou.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_05/26_for_range_map.html

package main

func main() {
	var m1 map[int]string = make(map[int]string)
	m1[0] = "nula"
	m1[1] = "jedna"
	m1[2] = "dva"
	m1[3] = "tri"
	m1[4] = "ctyri"
	m1[5] = "pet"
	m1[6] = "sest"

	for key, val := range m1 {
		println(key, val)
	}
}
