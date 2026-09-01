// ----------------------------------------------------------------------
// Podpora nativních SIMD operací v experimentálním balíčku simd/archsimd
// ----------------------------------------------------------------------

package main

import (
	"fmt"
	"simd/archsimd"
)

func AddTwoVectors(v1, v2 archsimd.Float32x16, result []float32) {
	// SIMD operace součtu prvků vektorů
	v3 := v1.Add(v2)

	// zápis prvků vektoru do řezu
	v3.Store(result)
}

func main() {
	// konstrukce a inicializace pole s odvozením počtu prvků
	a1 := [...]float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}
	a2 := [...]float32{0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5}

	// výpis obsahu polí
	fmt.Println("a1", a1)
	fmt.Println("a2", a2)

	// převod pole na vektor se čtyřmi prvky
	v1 := archsimd.LoadFloat32x16Array(&a1)
	v2 := archsimd.LoadFloat32x16Array(&a2)

	// konstrukce řezu
	result := make([]float32, 16)
	AddTwoVectors(v1, v2, result)

	// výpis obsahu řezu
	fmt.Println("a1+a2", result)
}
