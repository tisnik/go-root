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
//    https://tisnik.github.io/go-root/simd/vector_broadcast.html

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
