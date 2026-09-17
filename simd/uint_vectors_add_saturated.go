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
//    https://tisnik.github.io/go-root/simd/uint_vectors_add_saturated.html

package main

import (
	"fmt"
	"simd/archsimd"
)

func main() {
	// konstrukce a inicializace pole s odvozením počtu prvků
	a1 := [...]uint16{65530, 65530, 65530, 65530, 65530, 65530, 65530, 65530}
	a2 := [...]uint16{1, 2, 3, 4, 5, 6, 7, 8}

	// výpis obsahu polí
	fmt.Println("a1   ", a1)
	fmt.Println("a2   ", a2)

	// převod pole na vektor s osmi prvky
	v1 := archsimd.LoadUint16x8Array(&a1)
	v2 := archsimd.LoadUint16x8Array(&a2)

	// SIMD operace součtu prvků vektorů
	v3 := v1.AddSaturated(v2)

	// konstrukce řezu
	result := make([]uint16, 8)

	// zápis prvků vektoru do řezu
	v3.Store(result)

	// výpis obsahu řezu
	fmt.Println("a1+a2", result)
}
