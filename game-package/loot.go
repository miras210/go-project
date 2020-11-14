package game_package

//CHAIN OF RESPONSIBILITY PATTERN
type Loot interface {
	getType() string
	getValue() int
	setNext(Loot)
	String() string
}
type BaseLoot struct {
	next Loot
}

func (b *BaseLoot) getType() string {
	return "base"
}

func (b *BaseLoot) getValue() int {
	if b.next != nil {
		return b.next.getValue()
	} else {
		return 0
	}
}

func (b *BaseLoot) setNext(l Loot) {
	if b.next != nil {
		b.next.setNext(l)
	} else {
		b.next = l
	}
}

type AttackArtifact struct {
	BaseLoot
}

func (a *AttackArtifact) getType() string {
	return "attack"
}

type DefenseArtifact struct {
	BaseLoot
}

func (d *DefenseArtifact) getType() string {
	return "defense"
}

type Armor struct {
	DefenseArtifact
}

func (a *Armor) String() string {
	return "Armor"
}

func (a *Armor) getType() string {
	return a.DefenseArtifact.getType()
}

func (a *Armor) getValue() int {
	return 15 + a.BaseLoot.getValue()
}

type IronShield struct {
	DefenseArtifact
}

func (a *IronShield) getType() string {
	return a.DefenseArtifact.getType()
}

func (a *IronShield) getValue() int {
	return 10 + a.BaseLoot.getValue()
}
func (a *IronShield) String() string {
	return "Iron Shield"
}

type FlameRing struct {
	AttackArtifact
}

func (f *FlameRing) getType() string {
	return f.AttackArtifact.getType()
}

func (f *FlameRing) getValue() int {
	return 10 + f.BaseLoot.getValue()
}

func (a *FlameRing) String() string {
	return "Flame Ring"
}

type RingOfDarkMagic struct {
	AttackArtifact
}

func (f *RingOfDarkMagic) getType() string {
	return f.AttackArtifact.getType()
}

func (f *RingOfDarkMagic) getValue() int {
	return 15 + f.BaseLoot.getValue()
}

func (a *RingOfDarkMagic) String() string {
	return "Ring of Dark Magic"
}
func NewLoot() Loot {
	i := randomGen(0, 3)
	switch i {
	case 0:
		return &Armor{}
	case 1:
		return &FlameRing{}
	case 2:
		return &RingOfDarkMagic{}
	case 3:
		return &IronShield{}
	default:
		return nil
	}
}

//TODO Implement Loot that buffs Attack and Defense, that can sum with Chain of Responsibility
