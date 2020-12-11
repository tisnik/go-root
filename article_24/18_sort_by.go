// Seriál "Programovací jazyk Go"
//
// Dvacátá čtvrtá část
//     Kontejnery v základní knihovně programovacího jazyka Go
//     https://www.root.cz/clanky/kontejnery-v-zakladni-knihovne-programovaciho-jazyka-go/
//
// Demonstrační příklad číslo 18:
//     	Seřazení sekvence uživatelských datových struktur

package main

import (
	"fmt"
	"sort"
)

// Role represents user role in some IS
type Role struct {
	name    string
	surname string
}

func printRoles(roles []Role) {
	for i, role := range roles {
		fmt.Printf("#%d: %s %s\n", i, role.name, role.surname)
	}
}

func main() {
	roles := []Role{
		Role{"Eliška", "Najbrtová"},
		Role{"Jenny", "Suk"},
		Role{"Anička", "Šafářová"},
		Role{"Sváťa", "Pulec"},
		Role{"Blažej", "Motyčka"},
		Role{"Eda", "Wasserfall"},
		Role{"Přemysl", "Hájek"},
	}

	fmt.Println("Unsorted:")
	printRoles(roles)
	fmt.Println("--------------------")

	sort.Slice(roles, func(i, j int) bool { return roles[i].name < roles[j].name })

	fmt.Println("Sorted by name:")
	printRoles(roles)
	fmt.Println("--------------------")

	sort.Slice(roles, func(i, j int) bool { return roles[i].surname < roles[j].surname })

	fmt.Println("Sorted by surname:")
	printRoles(roles)
	fmt.Println("--------------------")
}
