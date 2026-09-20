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
//	https://tisnik.github.io/go-root/simd/mask_operations.html

package main

import (
	"fmt"
	"simd"
)

func createFourMasks() (simd.Mask32s, simd.Mask32s, simd.Mask32s, simd.Mask32s) {
	// konstrukce a inicializace řezu
	s1 := []float32{1, 2, 3, 4, 5, 6, 7, 8}
	s2 := []float32{4, 4, 4, 4, 4, 4, 4, 4}

	// převod řezu na vektor
	v1 := simd.LoadFloat32s(s1)
	v2 := simd.LoadFloat32s(s2)

	// SIMD operace porovnání prvků vektorů
	mask1 := v1.LessEqual(v2)
	mask2 := v1.GreaterEqual(v2)
	mask3 := v1.Less(v2)
	mask4 := v1.Greater(v2)
	return mask1, mask2, mask3, mask4
}

func main() {
	mask1, mask2, mask3, mask4 := createFourMasks()

	// výpis obsahu masek
	fmt.Println("mask 1        ", mask1.String())
	fmt.Println("mask 2        ", mask2.String())

	// operace s maskami
	mask5 := mask1.Or(mask2)
	fmt.Println("mask 1 | mask2", mask5.String())

	mask6 := mask1.And(mask2)
	fmt.Println("mask 1 & mask2", mask6.String())

	fmt.Println()

	// výpis obsahu masek
	fmt.Println("mask 3        ", mask3.String())
	fmt.Println("mask 4        ", mask4.String())

	// operace s maskami
	mask7 := mask3.Or(mask4)
	fmt.Println("mask 3 | mask4", mask7.String())

	mask8 := mask3.And(mask4)
	fmt.Println("mask 3 & mask4", mask8.String())
}
