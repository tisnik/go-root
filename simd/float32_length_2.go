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
//	https://tisnik.github.io/go-root/simd/float32_length_2.html

package main

import (
	"fmt"
	"simd"
)

func main() {
	// konstrukce a inicializace řezu
	s1 := []float32{5.0, 6.0, 7.0, 8.0}
	s2 := []float32{0.5, 0.5, 0.5, 0.5}

	// výpis obsahu řezů
	fmt.Println("s1   ", s1)
	fmt.Println("s2   ", s2)

	// převod řezu na vektor
	v1 := simd.LoadFloat32s(s1)
	v2 := simd.LoadFloat32s(s2)

	// výpis délky vektorů
	fmt.Println("len(v1)=", v1.Len())
	fmt.Println("len(v2)=", v2.Len())

	// SIMD operace součtu prvků vektorů
	v3 := v1.Add(v2)

	// konstrukce řezu
	result := make([]float32, v3.Len())

	// zápis prvků vektoru do řezu
	v3.Store(result)

	// výpis obsahu řezu
	fmt.Println("s1+s2", result)
}
