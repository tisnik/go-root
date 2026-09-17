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
//    https://tisnik.github.io/go-root/simd/vector_add_1.html

package main

import (
	"fmt"
	"simd/archsimd"
)

func main() {
	// konstrukce a inicializace pole s odvozením počtu prvků
	a1 := [...]float32{1.0, 2.0, 3.0, 4.0}
	a2 := [...]float32{0.5, 0.5, 0.5, 0.5}

	// výpis obsahu polí
	fmt.Println("a1", a1)
	fmt.Println("a2", a2)

	// převod pole na vektor se čtyřmi prvky
	v1 := archsimd.LoadFloat32x4Array(&a1)
	v2 := archsimd.LoadFloat32x4Array(&a2)

	// SIMD operace součtu prvků vektorů
	v3 := v1.Add(v2)

	// konstrukce řezu
	result := make([]float32, 4)

	// zápis prvků vektoru do řezu
	v3.Store(result)

	// výpis obsahu řezu
	fmt.Println("a1+a2", result)
}
