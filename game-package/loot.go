package game_package

type Loot interface {
	getType() string
	getValue() int
	setNext(Loot)
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

func (a *Armor) getType() string {
	return a.DefenseArtifact.getType()
}

func (a *Armor) getValue() int {
	return 10 + a.BaseLoot.getValue()
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
func NewLoot() Loot {
	i := randomGen(0, 1)
	switch i {
	case 0:
		return &Armor{}
	case 1:
		return &FlameRing{}
	default:
		return nil
	}
}

//TODO Implement Loot that buffs Attack and Defense, that can sum with Chain of Responsibility
