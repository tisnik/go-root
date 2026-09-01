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
	a1 := [...]float32{1.0, 2.0, 3.0, 4.0}

	// převod pole na vektor se čtyřmi prvky
	v1 := archsimd.LoadFloat32x4Array(&a1)

	// naplnění celého vektoru stejnou hodnotou
	v2 := archsimd.BroadcastFloat32x4(0.5)

	// konstrukce řezu
	result := make([]float32, 4)

	// SIMD operace součtu prvků vektorů
	v3 := v1.Add(v2)
	v3.Store(result)
	fmt.Println("add", result)

	// SIMD operace rozdílu prvků vektorů
	v3 = v1.Sub(v2)
	v3.Store(result)
	fmt.Println("sub", result)

	// SIMD operace součinu prvků vektorů
	v3 = v1.Mul(v2)
	v3.Store(result)
	fmt.Println("mul", result)

	// SIMD operace podílu prvků vektorů
	v3 = v1.Div(v2)
	v3.Store(result)
	fmt.Println("div", result)

	// SIMD operace výběru menších prvků vektorů
	v3 = v1.Min(v2)
	v3.Store(result)
	fmt.Println("min", result)

	// SIMD operace výběru větších prvků vektorů
	v3 = v1.Max(v2)
	v3.Store(result)
	fmt.Println("max", result)
}
