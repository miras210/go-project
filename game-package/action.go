package game_package

type ActionVisitor interface {
	visitForPlayer(*Player)
	visitForEnemy(*Enemy)
}

type MoveAction struct{}

func (m *MoveAction) visitForPlayer(p *Player) {
	panic("implement me")
}

func (m *MoveAction) visitForEnemy(e *Enemy) {
	panic("implement me")
}

type AttackAction struct{}

func (a *AttackAction) visitForPlayer(p *Player) {
	panic("implement me")
}

func (a *AttackAction) visitForEnemy(e *Enemy) {
	panic("implement me")
}

type LootAction struct{}

func (l *LootAction) visitForPlayer(p *Player) {
	panic("implement me")
}

// TODO ENEMY CAN NOT LOOT
func (l *LootAction) visitForEnemy(e *Enemy) {
	panic("implement me")
}
