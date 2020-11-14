package game_package

import "fmt"

//AbilityDecorator WRAPPER/DECORATOR PATTERN
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
func (s *SharpDecorator) Attack(i CharacterI) {
	fmt.Print("Sharp ")
	s.character.Attack(i)
}

func (s *SharpDecorator) String() string {
	return fmt.Sprintf("Sharp %v", s.character)
}

type StoneDecorator struct {
	AbilityDecorator
}

func (s *StoneDecorator) String() string {
	return fmt.Sprintf("Stone %v", s.character)
}

func (s *StoneDecorator) Attack(i CharacterI) {
	fmt.Print("Stone ")
	s.character.Attack(i)
}

func (s *StoneDecorator) GetDefense() int {
	return s.character.GetDefense() * 2
}
