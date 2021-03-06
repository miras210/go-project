package game_package

import "fmt"

//CharacterI
type CharacterI interface {
	GetAttack() int
	GetDefense() int
	GetHealth() int
	SetHealth(int)
	Attack(i CharacterI)
	isAlive() bool
}

//Character
type Character struct {
	health, attack, defense int
}

func (c *Character) String() string {
	return fmt.Sprintf("[HP: %v ATK: %v DEF: %v]", c.health, c.attack, c.defense)
}

func (c *Character) GetAttack() int {
	return c.attack
}
func (c *Character) GetDefense() int {
	return c.defense
}
func (c *Character) GetHealth() int {
	return c.health
}
func (c *Character) SetHealth(health int) {
	c.health = health
}
func (c *Character) Attack(character CharacterI) {
	coeff := float64(c.GetAttack()) / float64(character.GetDefense())
	if coeff > 1 {
		coeff = 1
	}
	resultingDamage := int(float64(c.GetAttack()) * coeff)
	fmt.Printf("Dealt %v damage to %v!\n", resultingDamage, character)
	character.SetHealth(character.GetHealth() - resultingDamage)
}

func (c *Character) isAlive() bool {
	if c.GetHealth() > 0 {
		return true
	} else {
		return false
	}
}
