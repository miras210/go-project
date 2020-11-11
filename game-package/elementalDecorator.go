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
func (e *AbilityDecorator) GetStamina() int {
	return e.character.GetStamina()
}
func (e *AbilityDecorator) SetStamina(stamina int) {
	e.character.SetStamina(stamina)
}
func (e *AbilityDecorator) GetMaxStamina() int {
	return e.character.GetMaxStamina()
}
func (e *AbilityDecorator) SetHealth(health int) {
	e.character.SetHealth(health)
}
func (e *AbilityDecorator) GetPlayerName() string {
	return e.character.GetPlayerName()
}
func (e *AbilityDecorator) SetPlayerName(playerName string) {
	e.character.SetPlayerName(playerName)
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

type EnduranceDecorator struct {
	AbilityDecorator
}

func (e *EnduranceDecorator) GetMaxStamina() int {
	return e.GetStamina() + 3
}
