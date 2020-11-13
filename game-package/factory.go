package game_package

type EnemyIdentifier int

const (
	ZombieGen EnemyIdentifier = iota
	SwordsmanGen
	ThiefGen
)

type AbilityIdentifier int

const (
	SharpGen AbilityIdentifier = iota
	StoneGen
	EnduranceGen
)

func EnemyFactory(enemyType EnemyIdentifier, indexDecorator AbilityIdentifier) CharacterI {
	var enemy CharacterI
	switch enemyType {
	case ZombieGen:
		enemy = NewZombie()
	case SwordsmanGen:
		enemy = NewSwordsman()
	case ThiefGen:
		enemy = NewThief()
	default:
		enemy = NewZombie()
	}
	switch indexDecorator {
	case SharpGen:
		enemy = &SharpDecorator{AbilityDecorator{character: enemy}}
	case StoneGen:
		enemy = &StoneDecorator{AbilityDecorator{character: enemy}}
	}
	return enemy
}
