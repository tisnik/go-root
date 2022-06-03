// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmnáctá část
//     Testování aplikací naprogramovaných v jazyce Go
//     https://www.root.cz/clanky/testovani-aplikaci-naprogramovanych-v-jazyce-go/
//
// Seznam příkladů ze sedmnácté části:
//    https://github.com/tisnik/go-root/blob/master/article_17/README.md
//
// Test práce s řezy (slices).
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_17/slice_test/slice_test.html

package slices

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSliceManipulation(t *testing.T) {
	Convey("Given an empty slice", t, func() {
		s := []int{}

		Convey("The slice should be empty initially", func() {
			So(s, ShouldBeEmpty)
		})

		Convey("When an item is added", func() {
			s = append(s, 1)

			Convey("The slice should not be empty", func() {
				So(s, ShouldNotBeEmpty)
			})
			Convey("The slice length should be one", func() {
				So(len(s), ShouldEqual, 1)
			})
			Convey("And length should NOT be zero, of course", func() {
				So(len(s), ShouldNotEqual, 0)
			})
			Convey("When another item is added", func() {
				s = append(s, 2)

				Convey("The slice length should be two", func() {
					So(len(s), ShouldEqual, 2)
				})
				Convey("And length should NOT be zero, of course", func() {
					So(len(s), ShouldNotEqual, 1)
				})
			})

			Convey("Now the slice length should be one again", func() {
				So(len(s), ShouldEqual, 1)
			})
		})
		Convey("And now the slice should be empty again", func() {
			So(s, ShouldBeEmpty)
		})
	})
}
