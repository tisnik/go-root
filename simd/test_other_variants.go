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
//    https://tisnik.github.io/go-root/simd/test_other_variants.html

package main

import (
	"fmt"
	"simd/archsimd"
)

func main() {
	// struktura zpřístupňující informace o vlastnostech CPU
	features := archsimd.X86

	// samotná struktura může být prázdná, resp. může obsahovat
	// nedostupné prvky (mimo svůj balíček)
	fmt.Println("Features struct:", features)

	// metoda vracející informaci o podpoře
	// rozšíření instrukční sady AVX
	fmt.Println("AVX support:", features.AVX())

	// metoda vracející informaci o podpoře
	// rozšíření instrukční sady AVX2
	fmt.Println("AVX2 support:", features.AVX2())

	// metoda vracející informaci o podpoře
	// rozšíření instrukční sady AVX512
	fmt.Println("AVX512 support:", features.AVX512())
}
