package main

import "bytes"

func BuildStringUsingPreallocatedStringBuffer(n int) string {
	// budeme pouzivat buffer
	var bb bytes.Buffer

	// alokace pameti pro pozdeji vkladane prvky
	bb.Grow(n * len("foo "))

	// postupne pridavani prvku do vysledneho retezce
	for i := 0; i < n; i++ {
		bb.WriteString("foo ")
	}

	// prevod objektu na retezec
	return bb.String()
}
