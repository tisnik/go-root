package main

func BuildStringUsingConcatenation(n int) string {
	// budeme pouzivat primo retezec
	s := ""

	// postupne pridavani prvku do vysledneho retezce
	for i := 0; i < n; i++ {
		s += "foo "
	}

	// vysledny retezec
	return s
}
