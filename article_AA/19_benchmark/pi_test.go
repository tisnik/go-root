package main

import (
	"math"
	"math/big"

	"testing"

	"github.com/shopspring/decimal"
)

func piBig(n int) float64 {
	var result big.Rat

	result.SetFrac64(2, 1)

	one := big.NewInt(1)
	limit := big.NewInt(1000)

	for n := big.NewInt(1); n.Cmp(limit) <= 0; n.Add(n, one) {
		m := big.NewInt(4)
		m.Mul(m, n)
		m.Mul(m, n)
		mn := big.NewInt(0)
		mn.Sub(m, one)

		var item big.Rat
		item.SetFrac(m, mn)
		result.Mul(&result, &item)

	}
	f, _ := result.Float64()
	return f
}

func piDecimal(n int) float64 {
	result := decimal.NewFromInt32(2)

	one := decimal.NewFromInt32(1)
	limit := decimal.NewFromInt32(200)

	for n := decimal.NewFromInt32(1); n.Cmp(limit) <= 0; n = n.Add(one) {
		m := decimal.NewFromInt32(4)
		m = m.Mul(n)
		m = m.Mul(n)
		mn := m.Sub(one)

		m = m.Div(mn)

		result = result.Mul(m)
	}
	f, _ := result.Float64()
	return f
}

func BenchmarkPiBig(b *testing.B) {
	for n := 1; n < b.N; n++ {
		f := piBig(n)
		if math.Abs(f-math.Pi) >= 0.01 {
			b.Fatal("Incorrect result", f)
		}
	}
}

func BenchmarkPiDecimal(b *testing.B) {
	for n := 1; n < b.N; n++ {
		f := piDecimal(n)
		if math.Abs(f-math.Pi) >= 0.01 {
			b.Fatal("Incorrect result", f)
		}
	}
}
