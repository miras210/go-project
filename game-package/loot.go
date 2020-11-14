package game_package

type Loot interface {
	getType() string
	getValue() int
}

type Pills struct {
	name      string
	addHealth int
}

func (p *Pills) getType() string {
	return p.name
}

func (p *Pills) getValue() int {
	return p.addHealth
}

type Redbull struct {
	name       string
	addStamina int
}

func (r *Redbull) getType() string {
	return r.name
}

func (r *Redbull) getValue() int {
	return r.addStamina
}
func NewLoot() (*Pills, *Redbull) {
	return &Pills{addHealth: 5, name: "pills"}, &Redbull{addStamina: 5, name: "redbull"}
}

//TODO Implement Loot that buffs Attack and Defense, that can sum with Chain of Responsibility
