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
	return "Char"
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
	resultingDamage := c.attack - character.GetDefense()
	if resultingDamage < 0 {
		resultingDamage = 0
	}
	fmt.Printf("Dealt %v damage!\n", resultingDamage)
	character.SetHealth(character.GetHealth() - resultingDamage)
}

func (c *Character) isAlive() bool {
	if c.GetHealth() > 0 {
		return true
	} else {
		return false
	}
}
