package game_package

type Loot interface {
	getType() string
	getValue() int
}

type MedicalPills struct{}

type EnergyDrink struct{}

//TODO Implement Loot that buffs Attack and Defense, that can sum with Chain of Responsibility
