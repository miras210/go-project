package game_package

//CharacterI
type CharacterI interface {
	GetPlayerName() string
	SetPlayerName(string)
	GetAttack() int
	GetDefense() int
	GetStamina() int
	SetStamina(int)
	GetMaxStamina() int
	GetHealth() int
	SetHealth(int)
}

//Character
type Character struct {
	health, attack, defense, stamina, maxStamina int
	playerName                                   string
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
func (c *Character) SetStamina(stamina int) {
	c.stamina = stamina
}
func (c *Character) GetMaxStamina() int {
	return c.maxStamina
}
func (c *Character) GetHealth() int {
	return c.health
}
func (c *Character) SetHealth(health int) {
	c.health = health
}
func (c *Character) GetPlayerName() string {
	return c.playerName
}
func (c *Character) SetPlayerName(playerName string) {
	c.playerName = playerName
}
