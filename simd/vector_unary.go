// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sto šestnáctá část
//    Dlouho očekávaná podpora SIMD operací v programovacím jazyku Go
//    https://www.root.cz/clanky/dlouho-ocekavana-podpora-simd-operaci-v-programovacim-jazyku-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze sto šestnácté části:
//    https://github.com/tisnik/go-root/blob/master/simd/README.md
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/simd/vector_unary.html

package main

import (
	"fmt"
	"simd/archsimd"
)

func main() {
	// konstrukce a inicializace pole s odvozením počtu prvků
	a1 := [...]float32{-1.0, -2.0, -3.0, -4.0}
	a2 := [...]float32{1.0, 2.0, 3.0, 4.0}

	// převod pole na vektor se čtyřmi prvky
	v1 := archsimd.LoadFloat32x4Array(&a1)
	v2 := archsimd.LoadFloat32x4Array(&a2)

	// konstrukce řezu
	result := make([]float32, 4)

	// SIMD operace výpočtu absolutní hodnoty
	v3 := v1.Abs()
	v2.Store(result)
	fmt.Println("abs", result)

	// SIMD operace změny znaménka
	v3 = v1.Neg()
	v3.Store(result)
	fmt.Println("neg", result)

	// SIMD operace výpočtu převrácené hodnoty
	v3 = v1.Reciprocal()
	v3.Store(result)
	fmt.Println("reciprocal", result)

	// SIMD operace výpočtu druhé odmocniny
	v3 = v2.Sqrt()
	v3.Store(result)
	fmt.Println("sqrt", result)

	// SIMD operace výpočtu převrácené hodnoty druhé odmocniny
	v3 = v2.ReciprocalSqrt()
	v3.Store(result)
	fmt.Println("reciprocal sqrt", result)
}
