package game_package

type ActionVisitor interface {
	visitForPlayer(i *CharacterI)
	visitForEnemy(i *CharacterI)
}

//TODO all of the actions will check for stamina! If not the action cannot be performed

type MoveAction struct{}

//TODO it will check, is it possible to move in that direction, WAIT what, how do we specify direction?
//Do we need polymorphism here?
func (m *MoveAction) visitForPlayer(p *CharacterI) {
	panic("implement me")
}

func (m *MoveAction) visitForEnemy(e *CharacterI) {
	panic("implement me")
}

type AttackAction struct{}

//TODO if something is out of range, Action won't be performed and no stamina will be used
func (a *AttackAction) visitForPlayer(p *CharacterI) {
	panic("implement me")
}

func (a *AttackAction) visitForEnemy(e *CharacterI) {
	panic("implement me")
}

type LootAction struct{}

func (l *LootAction) visitForPlayer(p *CharacterI) {
	panic("implement me")
}

// TODO ENEMY CAN NOT LOOT
func (l *LootAction) visitForEnemy(e *CharacterI) {
	panic("implement me")
}
