package main

import (
	"fmt"
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

	totalPopularity := 0
	for i, role := range roles {
		t.AppendRow(table.Row{i, role.name, role.surname, role.popularity})
		totalPopularity += role.popularity
	}
	t.AppendFooter(table.Row{"", "", "Total", totalPopularity})

	styles := []table.Style{
		table.StyleColoredBright,
		table.StyleColoredDark,
		table.StyleColoredBlackOnBlueWhite,
		table.StyleColoredBlackOnCyanWhite,
		table.StyleColoredBlackOnGreenWhite,
		table.StyleColoredBlackOnMagentaWhite,
		table.StyleColoredBlackOnYellowWhite,
		table.StyleColoredBlackOnRedWhite,
		table.StyleColoredBlueWhiteOnBlack,
		table.StyleColoredCyanWhiteOnBlack,
		table.StyleColoredGreenWhiteOnBlack,
		table.StyleColoredMagentaWhiteOnBlack,
		table.StyleColoredRedWhiteOnBlack,
		table.StyleColoredYellowWhiteOnBlack,
	}

	for _, style := range styles {
		fmt.Println(style.Name)
		t.SetStyle(style)
		t.Render()
		fmt.Println()
		fmt.Println()
	}
}
