package main

import "strings"

func BuildStringUsingStringBuilder(n int) string {
	// budeme pouzivat String Builder
	var sb strings.Builder

	// postupne pridavani prvku do vysledneho retezce
	for i := 0; i < n; i++ {
		sb.WriteRune('*')
	}

	// prevod objektu na retezec
	return sb.String()
}
