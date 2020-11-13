package game_package

import "fmt"

type Enemy struct {
	Character
}

type Zombie struct {
	Enemy
}

func (z *Zombie) Attack(i CharacterI) {
	fmt.Printf("%v attacks! ", z.String())
	z.Character.Attack(i)
}

func NewZombie() *Zombie {
	return &Zombie{Enemy{Character{
		health:  30,
		attack:  15,
		defense: 5,
	}}}
}

func (z *Zombie) String() string {
	return fmt.Sprintf("Zombie %v", z.Character.String())
}

type Swordsman struct {
	Enemy
}

func (s *Swordsman) Attack(i CharacterI) {
	fmt.Printf("%v attacks! ", s.String())
	s.Character.Attack(i)
}

func (s *Swordsman) String() string {
	return fmt.Sprintf("Swordsman %v", s.Character.String())
}

func NewSwordsman() *Swordsman {
	return &Swordsman{Enemy{Character{
		health:  40,
		attack:  20,
		defense: 7,
	}}}
}

type Thief struct {
	Enemy
}

func (t *Thief) Attack(i CharacterI) {
	fmt.Printf("%v attacks! ", t.String())
	t.Character.Attack(i)
}
func (t *Thief) String() string {
	return fmt.Sprintf("Thief %v", t.Character.String())
}
func NewThief() *Thief {
	return &Thief{Enemy{Character{
		health:  20,
		attack:  15,
		defense: 5,
	}}}
}
