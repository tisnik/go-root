package main

import "bytes"

func BuildStringUsingStringBuffer(n int) string {
	// budeme pouzivat buffer
	var bb bytes.Buffer

	// postupne pridavani prvku do vysledneho retezce
	for i := 0; i < n; i++ {
		bb.WriteString("foo ")
	}

	// prevod objektu na retezec
	return bb.String()
}
