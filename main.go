package main

import (
	"errors"
	"github.com/iszandro/wolf_sheep_straw/models"
)

var (
	boat        = models.Boat{}
	firstShore  = make([]models.LivingCreature, 0, 3)
	secondShore = make([]models.LivingCreature, 0, 3)
)

func initFirstShore() {
	addCreatureToShore(&firstShore, &models.Wolf{})
	addCreatureToShore(&firstShore, &models.Sheep{})
	addCreatureToShore(&firstShore, &models.Straw{})
}

func verifyLivingCreaturesInShore(shore []models.LivingCreature) (bool, error) {
	if len(shore) == 2 {
		firstCreature, secondCreature := shore[0], shore[1]
		if firstCreature.CanEat(secondCreature) {
			return false, errors.New("The " + firstCreature.Creature() + " ate the " + secondCreature.Creature() + "!")
		} else if secondCreature.CanEat(firstCreature) {
			return false, errors.New("The " + secondCreature.Creature() + " ate the " + firstCreature.Creature() + "!")
		}
	}

	return true, nil
}

func moveCreatureFromShoreToBoat(shore *[]models.LivingCreature, position int) (bool, error) {
	creature, err := removeCreatureFromShore(shore, position)

	if err != nil {
		return false, err
	}

	if !boat.Add(creature) {
		return false, errors.New("The boat is full")
	}

	return true, nil
}

func moveCreatureFromBoatToShore(shore *[]models.LivingCreature) (bool, error) {
	if boat.IsEmpty() {
		return false, errors.New("The boat doesn't have a creature to move")
	}

	addCreatureToShore(shore, boat.Remove())
	return true, nil
}

func addCreatureToShore(shore *[]models.LivingCreature, creature models.LivingCreature) {
	(*shore) = append((*shore), creature)
}

func removeCreatureFromShore(shore *[]models.LivingCreature, position int) (models.LivingCreature, error) {
	if position < 0 || position >= len(*shore) {
		return nil, errors.New("Can't remove that creature from the shore")
	}

	creature := (*shore)[position]

	if position == 0 {
		(*shore) = (*shore)[1:len(*shore)]
		return creature, nil
	} else if position == len(*shore)-1 {
		(*shore) = (*shore)[0 : len(*shore)-1]
		return creature, nil
	} else {
		tempShore := make([]models.LivingCreature, 0, 3)
		for i, c := range *shore {
			if i != position {
				addCreatureToShore(&tempShore, c)
			}
		}
		(*shore) = tempShore
		return creature, nil
	}
}

func crossRiver(positionOfCreatureToMove int) (bool, error) {
	var ok bool
	var err error

	// 1st step
	if ok, err = moveCreatureFromShoreToBoat(&firstShore, positionOfCreatureToMove); ok {

		// 2nd step
		if ok, err = verifyLivingCreaturesInShore(firstShore); ok {

			// 3rd step
			if ok, err = moveCreatureFromBoatToShore(&secondShore); ok {

				// 4th step
				if ok, err = verifyLivingCreaturesInShore(secondShore); ok {
					// 5th step
					crossRiver(0)
				} else {
					moveCreatureFromShoreToBoat(&secondShore, 0)
					moveCreatureFromBoatToShore(&firstShore)

					// 5th step
					crossRiver(0)
				}
			}
		}
	}

	return ok, err
}
