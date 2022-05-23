// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Druhá část:
//    Datové typy v programovacím jazyku Go
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go/
//
// Seznam příkladů z druhé části:
//    https://github.com/tisnik/go-root/blob/master/article_02/README.md
//
// Demonstrační příklad číslo 12
//    Datové typy reprezentující hodnoty s plovoucí řádovou čárkou, kontrola rozsahu
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_02/12_fp_types_checks.html

package main

func main() {
	var c float32 = 1e300
	var d float32 = -1e300

	var g float64 = 1e3000
	var h float64 = -1e3000
}
