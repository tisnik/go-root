// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá pátá část
//     Testování Go aplikací s využitím knihovny GΩmega a frameworku Ginkgo
//     https://www.root.cz/clanky/testovani-go-aplikaci-s-vyuzitim-knihovny-gomega-mega-a-frameworku-ginkgo/

package factorial_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "factorial"
)

var _ = Describe("Factorial", func() {
	Context("For negative input", func() {
		It("should be one", func() {
			Expect(Factorial(-1)).To(Equal(int64(1)))
			Expect(Factorial(-10)).To(Equal(int64(1)))
		})
	})
	Context("For zero input", func() {
		It("should be one", func() {
			Expect(Factorial(0)).To(Equal(int64(1)))
		})
	})
	Context("For positive input", func() {
		It("should be n!", func() {
			Expect(Factorial(1)).To(Equal(int64(1)))
			Expect(Factorial(2)).To(Equal(int64(2)))
			Expect(Factorial(9)).To(Equal(int64(362880)))
			Expect(Factorial(10)).To(Equal(int64(3628800)))
		})
	})
})
