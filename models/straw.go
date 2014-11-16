package models

type Straw struct{}

func (this *Straw) Creature() string {
	return "Straw"
}

func (this *Straw) Eat() string {
	return ""
}

func (this *Straw) CanEat(creature LivingCreature) bool {
	creatureType := creature.Creature()
	return this.Eat() == creatureType
}
