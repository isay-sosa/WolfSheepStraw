package main

import (
	"github.com/iszandro/wolf_sheep_straw/models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRemoveCreatureFromShore(t *testing.T) {
	Convey("Given we have in a shore a wolf, a sheep and straw", t, func() {
		initFirstShore()
		So(len(firstShore), ShouldEqual, 3)

		Convey("We can remove the first creature of that shore", func() {
			removedCreature, err := removeCreatureFromShore(&firstShore, 0)

			So(removedCreature, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(len(firstShore), ShouldEqual, 2)
		})

		Convey("We can remove the last creature of that shore", func() {
			removedCreature, err := removeCreatureFromShore(&firstShore, 2)

			So(removedCreature, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(len(firstShore), ShouldEqual, 2)
		})

		Convey("We can remove a creature from the middle of that shore", func() {
			removedCreature, err := removeCreatureFromShore(&firstShore, 1)

			So(removedCreature, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(len(firstShore), ShouldEqual, 2)
		})

		Convey("We can't remove a creature that doesn't exist in that shore", func() {
			removedCreature, err := removeCreatureFromShore(&firstShore, -1)

			So(removedCreature, ShouldBeNil)
			So(err.Error(), ShouldEqual, "Can't remove that creature from the shore")
			So(len(firstShore), ShouldEqual, 3)
		})

		Reset(reset)
	})
}

func TestMovingCreatureFromShoreToBoat(t *testing.T) {
	Convey("Given we have in a shore a wolf, a sheep and straw", t, func() {
		initFirstShore()

		Convey("When the boat is empty", func() {
			boat.Remove()
			So(boat.IsEmpty(), ShouldBeTrue)

			Convey("We can move the wolf to the boat", func() {
				moved, err := moveCreatureFromShoreToBoat(&firstShore, 0)

				So(moved, ShouldBeTrue)
				So(err, ShouldBeNil)

				Convey("And the shore should not have that creature", func() {
					So(len(firstShore), ShouldEqual, 2)
				})
			})
		})

		Convey("When the boat is not empty", func() {
			boat.Add(firstShore[0]) // Adding the wolf to the boat
			So(boat.IsNotEmpty(), ShouldBeTrue)

			Convey("We can't move the sheep to the boat", func() {
				moved, err := moveCreatureFromShoreToBoat(&firstShore, 1)

				So(moved, ShouldBeFalse)
				So(err.Error(), ShouldEqual, "The boat is full")
			})
		})

		Reset(reset)
	})
}

func TestMovingCreatureFromBoatToShore(t *testing.T) {
	Convey("Given we have a boat with a creature", t, func() {
		initFirstShore()
		moveCreatureFromShoreToBoat(&firstShore, 0)

		So(boat.IsNotEmpty(), ShouldBeTrue)

		Convey("We can move the creature to a shore", func() {
			moved, err := moveCreatureFromBoatToShore(&secondShore)

			So(moved, ShouldBeTrue)
			So(err, ShouldBeNil)
			So(len(secondShore), ShouldEqual, 1)
		})

		Reset(reset)
	})

	Convey("Given we have a boat with no creature", t, func() {
		So(boat.IsEmpty(), ShouldBeTrue)

		Convey("We can't move a creature to a shore", func() {
			moved, err := moveCreatureFromBoatToShore(&firstShore)

			So(moved, ShouldBeFalse)
			So(err.Error(), ShouldEqual, "The boat doesn't have a creature to move")
			So(len(firstShore), ShouldEqual, 0)
		})

		Reset(reset)
	})
}

func TestVerifyCreturesInShore(t *testing.T) {
	wolf := &models.Wolf{}
	sheep := &models.Sheep{}
	straw := &models.Straw{}

	Convey("When we have a sheep and the straw in the shore", t, func() {
		shore := []models.LivingCreature{sheep, straw}

		Convey("The sheep should eat the straw", func() {
			ok, err := verifyLivingCreaturesInShore(shore)
			So(ok, ShouldBeFalse)
			So(err.Error(), ShouldEqual, "The Sheep ate the Straw!")
		})
	})

	Convey("When we have a wolf and a sheep in the shore", t, func() {
		shore := []models.LivingCreature{wolf, sheep}

		Convey("The wolf should eat the sheep", func() {
			ok, err := verifyLivingCreaturesInShore(shore)
			So(ok, ShouldBeFalse)
			So(err.Error(), ShouldEqual, "The Wolf ate the Sheep!")
		})
	})

	Convey("When we have a wolf and the straw in the shore", t, func() {
		shore := []models.LivingCreature{wolf, straw}

		Convey("The wolf shouldn't eat the straw", func() {
			ok, err := verifyLivingCreaturesInShore(shore)
			So(ok, ShouldBeTrue)
			So(err, ShouldBeNil)
		})
	})
}

func TestWolfSheepStraw(t *testing.T) {
	Convey("When we have a wolf, a sheep and straw in a shore", t, func() {
		initFirstShore()

		So(len(firstShore), ShouldEqual, 3)
		So(len(secondShore), ShouldEqual, 0)

		Convey("And we need to move them to the other shore", func() {
			Convey("Then we start moving the wolf first", func() {
				ok, err := run(0)

				So(ok, ShouldBeFalse)
				So(err.Error(), ShouldEqual, "The Sheep ate the Straw!")
			})

			Convey("Then we start moving the sheep first", func() {
				ok, err := run(1)

				So(ok, ShouldBeTrue)
				So(err, ShouldBeNil)
				So(len(firstShore), ShouldEqual, 0)
				So(len(secondShore), ShouldEqual, 3)
			})

			Convey("Then we start moving the straw first", func() {
				ok, err := run(2)

				So(ok, ShouldBeFalse)
				So(err.Error(), ShouldEqual, "The Wolf ate the Sheep!")
			})
		})

		Reset(reset)
	})
}

func reset() {
	boat.Remove()
	firstShore = make([]models.LivingCreature, 0, 3)
	secondShore = make([]models.LivingCreature, 0, 3)
}
