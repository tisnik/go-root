// Seriál "Programovací jazyk Go"
//	https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sto sedmnáctá část
//	Dlouho očekávaná podpora SIMD operací v programovacím jazyku Go (dokončení)
//	https://www.root.cz/clanky/dlouho-ocekavana-podpora-simd-operaci-v-programovacim-jazyku-go-dokonceni/
//
// Repositář:
//	https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze sto sedmnácté části:
//	https://github.com/tisnik/go-root/blob/master/simd/README.md
//
// Dokumentace ve stylu "literate programming":
//	https://tisnik.github.io/go-root/simd/mask_selection_2.html

package main

import (
	"fmt"
	"simd"
)

func createMask() simd.Mask32s {
	// konstrukce a inicializace řezu
	s1 := []float32{1, 2, 3, 4, 5, 6, 7, 8}
	s2 := []float32{4, 4, 4, 4, 4, 4, 4, 4}

	// převod řezu na vektor
	v1 := simd.LoadFloat32s(s1)
	v2 := simd.LoadFloat32s(s2)

	// SIMD operace porovnání prvků vektorů
	return v1.GreaterEqual(v2)
}

func main() {
	mask := createMask()

	// konstrukce a inicializace řezu
	s1 := []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}
	s2 := []float32{99, 99, 99, 99, 99, 99, 99, 99}

	// výpis obsahu řezu
	fmt.Println("original", s1)

	// převod řezu na vektor
	v1 := simd.LoadFloat32s(s1)
	v2 := simd.LoadFloat32s(s2)

	// výběr prvků na základě masky
	v3 := v1.IfElse(mask, v2)

	// konstrukce řezu
	result := make([]float32, v3.Len())

	// zápis prvků vektoru do řezu
	v3.Store(result)

	// výpis obsahu řezu
	fmt.Println("selected", result)
}
