// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá devátá část
//    Generické datové typy v jazyce Go?
//    https://www.root.cz/clanky/genericke-datove-typy-v-jazyce-go/
//
// Seznam příkladů z padesáté deváté části:
//    https://github.com/tisnik/go-root/blob/master/article_59/README.md

package adder

import "github.com/cheekybits/genny/generic"

type NumberType generic.Number

func Add(x, y NumberType) NumberType {
	return x + y
}
