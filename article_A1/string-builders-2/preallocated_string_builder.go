package main

import "strings"

func BuildStringUsingPreallocatedStringBuilder(n int) string {
	// budeme pouzivat String Builder
	var sb strings.Builder

	// alokace pameti pro pozdeji vkladane prvky
	sb.Grow(n * len("foo "))

	// postupne pridavani prvku do vysledneho retezce
	for i := 0; i < n; i++ {
		sb.WriteString("foo ")
	}

	// prevod objektu na retezec
	return sb.String()
}
