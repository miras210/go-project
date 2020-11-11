package game_package

type Zombie struct {
	Character
}

func NewZombie() *Zombie {
	return &Zombie{Character{
		health:     30,
		attack:     15,
		defense:    5,
		stamina:    5,
		maxStamina: 5,
	}}
}

type Swordsman struct {
	Character
}

func NewSwordsman() *Swordsman {
	return &Swordsman{Character{
		health:     40,
		attack:     20,
		defense:    7,
		stamina:    2,
		maxStamina: 2,
	}}
}

type Thief struct {
	Character
}

func NewThief() *Thief {
	return &Thief{Character{
		health:     20,
		attack:     15,
		defense:    5,
		stamina:    7,
		maxStamina: 7,
	}}
}
