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
//    https://tisnik.github.io/go-root/simd/int_vectors_add_1.html

package main

import (
	"fmt"
	"simd/archsimd"
)

func main() {
	// konstrukce a inicializace pole s odvozením počtu prvků
	a1 := [...]int32{1, 2, 3, 4}
	a2 := [...]int32{5, 6, 7, 8}

	// výpis obsahu polí
	fmt.Println("a1   ", a1)
	fmt.Println("a2   ", a2)

	// převod pole na vektor se čtyřmi prvky
	v1 := archsimd.LoadInt32x4Array(&a1)
	v2 := archsimd.LoadInt32x4Array(&a2)

	// SIMD operace součtu prvků vektorů
	v3 := v1.Add(v2)

	// konstrukce řezu
	result := make([]int32, 4)

	// zápis prvků vektoru do řezu
	v3.Store(result)

	// výpis obsahu řezu
	fmt.Println("a1+a2", result)
}
