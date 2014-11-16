package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBoat(t *testing.T) {
	boat := Boat{}

	Convey("When boat doesn't have a creature", t, func() {
		So(boat.IsEmpty(), ShouldBeTrue)

		Convey("We can add a creature", func() {
			added := boat.Add(&Wolf{})
			So(added, ShouldBeTrue)
			So(boat.IsNotEmpty(), ShouldBeTrue)
		})
	})

	Convey("When boat has already a creature", t, func() {
		boat.Add(&Wolf{})

		Convey("We can't add other creature", func() {
			added := boat.Add(&Sheep{})
			So(added, ShouldBeFalse)
		})

		Convey("We can remove the creature from the boat", func() {
			creature := boat.Remove()
			So(boat.IsEmpty(), ShouldBeTrue)
			So(creature, ShouldNotBeNil)
		})
	})
}
