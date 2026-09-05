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
	a1 := [...]int16{1<<15 - 1, 1<<15 - 1, 1<<15 - 1, 1<<15 - 1, 1<<15 - 1, 1<<15 - 1, 1<<15 - 1, 1<<15 - 1}
	a2 := [...]int16{-3, -2, -1, 0, 1, 2, 3, 4}

	// výpis obsahu polí
	fmt.Println("a1   ", a1)
	fmt.Println("a2   ", a2)

	// převod pole na vektor s osmi prvky
	v1 := archsimd.LoadInt16x8Array(&a1)
	v2 := archsimd.LoadInt16x8Array(&a2)

	// SIMD operace součtu prvků vektorů
	v3 := v1.Add(v2)

	// konstrukce řezu
	result := make([]int16, 8)

	// zápis prvků vektoru do řezu
	v3.Store(result)

	// výpis obsahu řezu
	fmt.Println("a1+a2", result)
}
