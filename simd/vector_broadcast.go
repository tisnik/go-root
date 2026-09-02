// ----------------------------------------------------------------------
// Podpora nativních SIMD operací v experimentálním balíčku simd/archsimd
// ----------------------------------------------------------------------

package main

import (
	"fmt"
	"simd/archsimd"
)

func AddTwoVectors(v1, v2 archsimd.Float32x4, result []float32) {
	// SIMD operace součtu prvků vektorů
	v3 := v1.Add(v2)

	// zápis prvků vektoru do řezu
	v3.Store(result)
}

func main() {
	// konstrukce a inicializace pole s odvozením počtu prvků
	a1 := [...]float32{1.0, 2.0, 3.0, 4.0}

	// převod pole na vektor se čtyřmi prvky
	v1 := archsimd.LoadFloat32x4Array(&a1)

	// naplnění celého vektoru stejnou hodnotou
	v2 := archsimd.BroadcastFloat32x4(0.5)

	// konstrukce řezu
	result := make([]float32, 4)

	// provedení vybrané operace s vektory
	AddTwoVectors(v1, v2, result)

	// výpis obsahu řezu
	fmt.Println(result)
}
