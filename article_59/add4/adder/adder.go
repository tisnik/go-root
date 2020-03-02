package adder

import "github.com/cheekybits/genny/generic"

type NumberType generic.Number

func NumberTypeAdd(x NumberType, y NumberType) NumberType {
	return x + y
}
