package models

type LivingCreature interface {
	Creature() string
	Eat() string
	CanEat(creature LivingCreature) bool
}
