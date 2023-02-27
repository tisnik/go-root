package main

import (
	"bytes"
	"strings"

	"testing"
)

func checkBuiltString(b *testing.B, s *string) {
	// kontrola delky vysledneho retezce
	if len(*s) != b.N {
		b.Fatal("Wrong string length", len(*s), b.N)
	}

	// kontrola obsahu vysledneho retezce
	for i, r := range *s {
		if r != '*' {
			b.Fatal("Wrong rune", r, "at index", i)
		}
	}
}

func BenchmarkBuildStringUsingConcatenation(b *testing.B) {
	// budeme pouzivat primo retezec
	s := ""

	// postupne pridavani prvku do vysledneho retezce
	for n := 0; n < b.N; n++ {
		s += "*"
	}

	// zkontrolovat obsah vytvoreneho retezce po zastaveni casovace
	b.StopTimer()
	checkBuiltString(b, &s)
}

func BenchmarkBuildStringUsingStringBuffer(b *testing.B) {
	// budeme pouzivat buffer
	var bb bytes.Buffer

	// postupne pridavani prvku do vysledneho retezce
	for n := 0; n < b.N; n++ {
		bb.WriteRune('*')
	}

	// prevod objektu na retezec
	s := bb.String()

	// zkontrolovat obsah vytvoreneho retezce po zastaveni casovace
	b.StopTimer()
	checkBuiltString(b, &s)
}

func BenchmarkBuildStringUsingPreallocatedStringBuffer(b *testing.B) {
	// budeme pouzivat buffer
	var bb bytes.Buffer

	// alokace pameti pro pozdeji vkladane prvky
	bb.Grow(b.N)

	// postupne pridavani prvku do vysledneho retezce
	for n := 0; n < b.N; n++ {
		bb.WriteRune('*')
	}

	// prevod objektu na retezec
	s := bb.String()

	// zkontrolovat obsah vytvoreneho retezce po zastaveni casovace
	b.StopTimer()
	checkBuiltString(b, &s)
}

func BenchmarkBuildStringUsingStringBuilder(b *testing.B) {
	// budeme pouzivat String Builder
	var sb strings.Builder

	// postupne pridavani prvku do vysledneho retezce
	for n := 0; n < b.N; n++ {
		sb.WriteRune('*')
	}

	// prevod objektu na retezec
	s := sb.String()

	// zkontrolovat obsah vytvoreneho retezce po zastaveni casovace
	b.StopTimer()
	checkBuiltString(b, &s)
}

func BenchmarkBuildStringUsingPreallocatedStringBuilder(b *testing.B) {
	// budeme pouzivat String Builder
	var sb strings.Builder

	// alokace pameti pro pozdeji vkladane prvky
	sb.Grow(b.N)

	// postupne pridavani prvku do vysledneho retezce
	for n := 0; n < b.N; n++ {
		sb.WriteRune('*')
	}

	// prevod objektu na retezec
	s := sb.String()

	// zkontrolovat obsah vytvoreneho retezce po zastaveni casovace
	b.StopTimer()
	checkBuiltString(b, &s)
}
