package models

type Boat struct {
	Creature LivingCreature
}

func (this *Boat) Add(creature LivingCreature) bool {
	if this.IsNotEmpty() {
		return false
	}

	this.Creature = creature
	return true
}

func (this *Boat) Remove() LivingCreature {
	creature := this.Creature
	this.Creature = nil
	return creature
}

func (this *Boat) IsEmpty() bool {
	return this.Creature == nil
}

func (this *Boat) IsNotEmpty() bool {
	return !this.IsEmpty()
}
