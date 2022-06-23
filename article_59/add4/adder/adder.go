// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá devátá část
//    Generické datové typy v jazyce Go?
//    https://www.root.cz/clanky/genericke-datove-typy-v-jazyce-go/

package adder

import "github.com/cheekybits/genny/generic"

type NumberType generic.Number

func NumberTypeAdd(x NumberType, y NumberType) NumberType {
	return x + y
}
