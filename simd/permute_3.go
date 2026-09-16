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
//    https://tisnik.github.io/go-root/simd/permute_3.html

package main

import (
	"fmt"
	"simd/archsimd"
)

func main() {
	// konstrukce a inicializace pole s odvozením počtu prvků
	a1 := [...]float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}

	// výpis obsahu pole
	fmt.Println("a1       = ", a1)

	// převod pole na vektor s osmi prvky
	v1 := archsimd.LoadFloat32x8Array(&a1)

	indices := archsimd.LoadUint32x8([]uint32{0, 0, 7, 7, 0, 0, 7, 7})

	// výpočet permutace
	v2 := v1.Permute(indices)

	// konstrukce řezu
	result := make([]float32, 8)

	// zápis prvků vektoru do řezu
	v2.Store(result)

	// výpis obsahu řezu
	fmt.Println("permuted = ", result)
}
