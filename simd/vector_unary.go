// ----------------------------------------------------------------------
// Podpora nativních SIMD operací v experimentálním balíčku simd/archsimd
// ----------------------------------------------------------------------

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
