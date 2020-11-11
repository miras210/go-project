package game_package

//ElementalDecorator
type ElementalDecorator struct {
	character CharacterI
}

func (e *ElementalDecorator) GetAttack() int {
	return e.character.GetAttack()
}
func (e *ElementalDecorator) GetDefense() int {
	return e.character.GetDefense()
}
func (e *ElementalDecorator) GetHealth() int {
	return e.character.GetHealth()
}
func (e *ElementalDecorator) GetStamina() int {
	return e.character.GetStamina()
}
func (e *ElementalDecorator) SetHealth(health int) {
	e.character.SetHealth(health)
}

//Decorators
type FireElemental struct {
	ElementalDecorator
}

//TODO implement FireElementalDecorator
type ElectroElemental struct {
	ElementalDecorator
}

//TODO implement ElectroElementalDecorator
type AnemoElemental struct {
	ElementalDecorator
}

//TODO implement AnemoElementalDecorator
