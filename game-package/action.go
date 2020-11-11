package game_package

type ActionVisitor interface {
	visitForPlayer(i *CharacterI)
	visitForEnemy(i *CharacterI)
}

type MoveAction struct{}

func (m *MoveAction) visitForPlayer(p *CharacterI) {
	panic("implement me")
}

func (m *MoveAction) visitForEnemy(e *CharacterI) {
	panic("implement me")
}

type AttackAction struct{}

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
