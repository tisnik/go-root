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
//	https://tisnik.github.io/go-root/simd/float32_comparison.html

package main

import (
	"fmt"
	"simd"
)

func main() {
	// konstrukce a inicializace řezu
	s1 := []float32{1, 2, 3, 4, 5, 6, 7, 8}
	s2 := []float32{4, 4, 4, 4, 4, 4, 4, 4}

	// výpis obsahu řezů
	fmt.Println("s1    ", s1)
	fmt.Println("s2    ", s2)

	// převod řezu na vektor
	v1 := simd.LoadFloat32s(s1)
	v2 := simd.LoadFloat32s(s2)

	// SIMD operace porovnání prvků vektorů
	v3 := v1.Equal(v2)

	// výpis obsahu řezu
	fmt.Println("s1==s2", v3.String())

	// SIMD operace porovnání prvků vektorů
	v4 := v1.Less(v2)

	// výpis obsahu řezu
	fmt.Println("s1<s2 ", v4.String())

	// SIMD operace porovnání prvků vektorů
	v5 := v1.LessEqual(v2)

	// výpis obsahu řezu
	fmt.Println("s1<=s2", v5.String())

	// SIMD operace porovnání prvků vektorů
	v6 := v1.Greater(v2)

	// výpis obsahu řezu
	fmt.Println("s1>s2 ", v6.String())

	// SIMD operace porovnání prvků vektorů
	v7 := v1.GreaterEqual(v2)

	// výpis obsahu řezu
	fmt.Println("s1>=s2", v7.String())

	// SIMD operace porovnání prvků vektorů
	v8 := v1.NotEqual(v2)

	// výpis obsahu řezu
	fmt.Println("s1<>s2", v8.String())
}
