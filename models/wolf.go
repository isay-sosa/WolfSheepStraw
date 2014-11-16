package models

type Wolf struct{}

func (this *Wolf) Creature() string {
	return "Wolf"
}

func (this *Wolf) Eat() string {
	return "Sheep"
}

func (this *Wolf) CanEat(creature LivingCreature) bool {
	creatureType := creature.Creature()
	return this.Eat() == creatureType
}
