package models

type Sheep struct{}

func (this *Sheep) Creature() string {
	return "Sheep"
}

func (this *Sheep) Eat() string {
	return "Straw"
}

func (this *Sheep) CanEat(creature LivingCreature) bool {
	creatureType := creature.Creature()
	return this.Eat() == creatureType
}
