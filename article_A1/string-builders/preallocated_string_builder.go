package main

import "strings"

func BuildStringUsingPreallocatedStringBuilder(n int) string {
	// budeme pouzivat String Builder
	var sb strings.Builder

	// alokace pameti pro pozdeji vkladane prvky
	sb.Grow(n)

	// postupne pridavani prvku do vysledneho retezce
	for i := 0; i < n; i++ {
		sb.WriteRune('*')
	}

	// prevod objektu na retezec
	return sb.String()
}
