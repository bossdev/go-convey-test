package handle

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTestGetNumber(t *testing.T) {
	Convey("Given number", t, func() {
		Convey("It must be return a correct value", func() {
			So(GetNumber(1), ShouldEqual, "1")
			So(GetNumber(2), ShouldEqual, "2")
			So(GetNumber(3), ShouldEqual, "FIZZ")
			So(GetNumber(4), ShouldEqual, "4")
			So(GetNumber(5), ShouldEqual, "BUZZ")
			So(GetNumber(6), ShouldEqual, "FIZZ")
			So(GetNumber(7), ShouldEqual, "7")
			So(GetNumber(8), ShouldEqual, "8")
			So(GetNumber(9), ShouldEqual, "FIZZ")
			So(GetNumber(10), ShouldEqual, "BUZZ")
			So(GetNumber(11), ShouldEqual, "11")
			So(GetNumber(12), ShouldEqual, "FIZZ")
			So(GetNumber(13), ShouldEqual, "13")
			So(GetNumber(14), ShouldEqual, "14")
			So(GetNumber(15), ShouldEqual, "FIZZBUZZ")
		})
	})
}
