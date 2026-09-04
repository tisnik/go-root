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
	a1 := [...]float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}
	a2 := [...]float32{5.0, 5.0, 5.0, 5.0, 5.0, 5.0, 5.0, 5.0}

	// výpis obsahu polí
	fmt.Println("a1    ", a1)
	fmt.Println("a2    ", a2)

	// převod pole na vektor s osmi prvky
	v1 := archsimd.LoadFloat32x8Array(&a1)
	v2 := archsimd.LoadFloat32x8Array(&a2)

	// SIMD operace porovnání prvků vektoru
	mask := v1.Less(v2)
	fmt.Println("<     ", mask.String())

	v3 := v1.Masked(mask)

	// konstrukce řezu
	result := make([]float32, 8)

	// zápis prvků vektoru do řezu
	v3.Store(result)

	// výpis obsahu řezu
	fmt.Println("masked", v3)
}
