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
