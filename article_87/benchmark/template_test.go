package main

import (
	"testing"
)

func BenchmarkReadTemplate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmpl := readTemplate("./html_template05.htm")
		if tmpl == nil {
			b.Fatal("Template was not created")
		}
	}
}

func BenchmarkApplyTemplate(b *testing.B) {
	// tyto hodnoty budou použity při aplikaci šablony
	roles := []Role{
		Role{"Eliška", "Najbrtová", 4},
		Role{"Jenny", "Suk", 3},
		Role{"Anička", "Šafářová", 0},
		Role{"Sváťa", "Pulec", 3},
		Role{"Blažej", "Motyčka", 8},
		Role{"Eda", "Wasserfall", 0},
		Role{"Přemysl", "Hájek", 10},
	}

	tmpl := readTemplate("./html_template05.htm")

	for i := 0; i < b.N; i++ {
		length := applyTemplate(tmpl, roles)
		if length <= 0 {
			b.Fatal("Don't work")
		}
	}
}

func BenchmarkReadAndApplyTemplate(b *testing.B) {
	// tyto hodnoty budou použity při aplikaci šablony
	roles := []Role{
		Role{"Eliška", "Najbrtová", 4},
		Role{"Jenny", "Suk", 3},
		Role{"Anička", "Šafářová", 0},
		Role{"Sváťa", "Pulec", 3},
		Role{"Blažej", "Motyčka", 8},
		Role{"Eda", "Wasserfall", 0},
		Role{"Přemysl", "Hájek", 10},
	}

	for i := 0; i < b.N; i++ {
		tmpl := readTemplate("./html_template05.htm")
		if tmpl == nil {
			b.Fatal("Template was not created")
		}

		length := applyTemplate(tmpl, roles)
		if length <= 0 {
			b.Fatal("Don't work")
		}
	}
}
