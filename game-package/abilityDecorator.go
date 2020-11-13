package game_package

//AbilityDecorator
type AbilityDecorator struct {
	character CharacterI
}

func (e *AbilityDecorator) GetAttack() int {
	return e.character.GetAttack()
}
func (e *AbilityDecorator) GetDefense() int {
	return e.character.GetDefense()
}
func (e *AbilityDecorator) GetHealth() int {
	return e.character.GetHealth()
}
func (e *AbilityDecorator) SetHealth(health int) {
	e.character.SetHealth(health)
}
func (e *AbilityDecorator) Attack(i CharacterI) {
	e.character.Attack(i)
}
func (e *AbilityDecorator) isAlive() bool {
	return e.character.isAlive()
}

//Decorators
type SharpDecorator struct {
	AbilityDecorator
}

func (s *SharpDecorator) GetAttack() int {
	return s.character.GetAttack() * 2
}

type StoneDecorator struct {
	AbilityDecorator
}

func (s *StoneDecorator) GetDefense() int {
	return s.character.GetDefense() * 2
}
