// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sto sedmnáctá část
//    Dlouho očekávaná podpora SIMD operací v programovacím jazyku Go (2. část)
//    https://www.root.cz/clanky/dlouho-ocekavana-podpora-simd-operaci-v-programovacim-jazyku-go-2-cast/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze sto sedmnácté části:
//    https://github.com/tisnik/go-root/blob/master/simd/README.md
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/simd/dot_product_1.html

package main

import (
	"fmt"
	"simd/archsimd"
)

func main() {
	// konstrukce a inicializace pole s odvozením počtu prvků
	a1 := [...]int16{0, 1, 2, 3, 4, 5, 6, 7}
	a2 := [...]int16{0, 1, 2, 3, 4, 5, 6, 7}

	// výpis obsahu polí
	fmt.Println("a1 ", a1)
	fmt.Println("a2 ", a2)

	// převod pole na vektor s osmi prvky
	v1 := archsimd.LoadInt16x8Array(&a1)
	v2 := archsimd.LoadInt16x8Array(&a2)

	// SIMD operace součtu prvků vektorů
	v3 := v1.DotProductPairs(v2)

	// konstrukce řezu
	result := make([]int32, 4)

	// zápis prvků vektoru do řezu
	v3.Store(result)

	// výpis obsahu řezu
	fmt.Println("dot", result)
}
