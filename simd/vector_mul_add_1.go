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
	a2 := [...]float32{2.0, 2.0, 2.0, 2.0}
	a3 := [...]float32{100.0, 100.0, 100.0, 100.0}

	// výpis obsahu polí
	fmt.Println("a1 =       ", a1)
	fmt.Println("a2 =       ", a2)
	fmt.Println("a3 =       ", a3)

	// převod pole na vektor se čtyřmi prvky
	v1 := archsimd.LoadFloat32x4Array(&a1)
	v2 := archsimd.LoadFloat32x4Array(&a2)
	v3 := archsimd.LoadFloat32x4Array(&a3)

	// výpočet
	v4 := v1.MulAdd(v2, v3)

	// konstrukce řezu
	result := make([]float32, 4)

	// zápis prvků vektoru do řezu
	v4.Store(result)

	// výpis obsahu řezu
	fmt.Println("a1*a2+a3 = ", result)
}
