package game_package

type Enemy struct {
	Character
}

type Zombie struct {
	Enemy
}

func NewZombie() *Zombie {
	return &Zombie{Enemy{Character{
		health:     30,
		attack:     15,
		defense:    5,
		stamina:    5,
		maxStamina: 5,
	}}}
}

type Swordsman struct {
	Enemy
}

func NewSwordsman() *Swordsman {
	return &Swordsman{Enemy{Character{
		health:     40,
		attack:     20,
		defense:    7,
		stamina:    2,
		maxStamina: 2,
	}}}
}

type Thief struct {
	Enemy
}

func NewThief() *Thief {
	return &Thief{Enemy{Character{
		health:     20,
		attack:     15,
		defense:    5,
		stamina:    7,
		maxStamina: 7,
	}}}
}
