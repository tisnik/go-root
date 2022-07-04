// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá první část
//    Formátovaný tisk obsahu tabulek a dalších datových struktur v Go
//    https://www.root.cz/clanky/formatovany-tisk-obsahu-tabulek-a-dalsich-datovych-struktur-v-go/

package main

import (
	"os"

	"github.com/jedib0t/go-pretty/table"
)

// Role represents user role in some IS
type Role struct {
	name       string
	surname    string
	popularity int
}

func main() {
	roles := []Role{
		Role{"Eliška", "Najbrtová", 4},
		Role{"Jenny", "Suk", 3},
		Role{"Anička", "Šafářová", 1},
		Role{"Sváťa", "Pulec", 3},
		Role{"Blažej", "Motyčka", 8},
		Role{"Eda", "Wasserfall", 3},
		Role{"Přemysl", "Hájek", 10},
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "First Name", "Last Name", "Popularity"})

	for i, role := range roles {
		t.AppendRow(table.Row{i, role.name, role.surname, role.popularity})
	}
	t.Render()
}
