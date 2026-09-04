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
	a1 := [...]float32{0.4, 0.5, 0.50001, 0.6}

	fmt.Println("a1", a1)

	// převod pole na vektor se čtyřmi prvky
	v1 := archsimd.LoadFloat32x4Array(&a1)

	// konstrukce řezu
	result := make([]float32, 4)

	v2 := v1.Floor()
	v2.Store(result)
	fmt.Println("floor", result)

	v2 = v1.Ceil()
	v2.Store(result)
	fmt.Println("ceil ", result)

	v2 = v1.Trunc()
	v2.Store(result)
	fmt.Println("trunc", result)

	v2 = v1.Round()
	v2.Store(result)
	fmt.Println("round", result)
}
