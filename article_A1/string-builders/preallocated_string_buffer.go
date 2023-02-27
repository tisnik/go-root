package main

import "bytes"

func BuildStringUsingPreallocatedStringBuffer(n int) string {
	// budeme pouzivat buffer
	var bb bytes.Buffer

	// alokace pameti pro pozdeji vkladane prvky
	bb.Grow(n)

	// postupne pridavani prvku do vysledneho retezce
	for i := 0; i < n; i++ {
		bb.WriteRune('*')
	}

	// prevod objektu na retezec
	return bb.String()
}
