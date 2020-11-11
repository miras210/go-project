package game_package

//CharacterI
type CharacterI interface {
	GetAttack() int
	GetDefense() int
	GetStamina() int
	GetHealth() int
	SetHealth(int)
}

//Character
type Character struct {
	health, attack, defense, stamina int
}

func (c *Character) GetAttack() int {
	return c.attack
}
func (c *Character) GetDefense() int {
	return c.defense
}
func (c *Character) GetStamina() int {
	return c.stamina
}
func (c *Character) GetHealth() int {
	return c.health
}
func (c *Character) SetHealth(health int) {
	c.health = health
}
