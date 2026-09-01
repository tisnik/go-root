// ----------------------------------------------------------------------
// Podpora nativních SIMD operací v experimentálním balíčku simd/archsimd
// ----------------------------------------------------------------------

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
}
