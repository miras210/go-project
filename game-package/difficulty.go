package game_package

// DIFFICULTY STRATEGY

type Difficulty interface {
	getNewPlayer(string) *Player // Player struct
	getNewEnemies() []CharacterI // []Enemy struct
	getNewLoots() []Loot         // []Loot struct
	getNewGameMap() *Map         // GameMap struct
}

func newDifficulty(difficulty string) Difficulty {
	switch difficulty {
	case "easy":
		return &EasyDifficulty{}
	case "medium":
		return &MediumDifficulty{}
	case "hard":
		return &HardDifficulty{}
	default:
		return &MediumDifficulty{}
	}
}

type EasyDifficulty struct{}

func (e *EasyDifficulty) getNewPlayer(playerName string) *Player {
	return &Player{
		Character: Character{
			health:     100,
			attack:     20,
			defense:    10,
			stamina:    4,
			maxStamina: 4},
		playerName: playerName,
	}
}

func (e *EasyDifficulty) getNewEnemies() []CharacterI {
	var enemies []CharacterI
	for i := 0; i < 3; i++ {
		enemies = append(enemies, EnemyFactory(EnemyIdentifier(randomGen(0, 2)),
			AbilityIdentifier(-1)))
	}
	return enemies
}

func (e *EasyDifficulty) getNewLoots() []Loot {
	panic("implement me")
}

func (e *EasyDifficulty) getNewGameMap() *Map {
	panic("implement me")
}

type MediumDifficulty struct{}

func (m *MediumDifficulty) getNewPlayer(playerName string) *Player {
	return &Player{
		Character: Character{
			health:     80,
			attack:     15,
			defense:    7,
			stamina:    3,
			maxStamina: 3,
		},
		playerName: playerName,
	}
}

func (m *MediumDifficulty) getNewEnemies() []CharacterI {
	var enemies []CharacterI
	for i := 0; i < 5; i++ {
		enemies = append(enemies, EnemyFactory(EnemyIdentifier(randomGen(0, 2)),
			AbilityIdentifier(randomGen(1, 2))))
	}
	return enemies
}

func (m *MediumDifficulty) getNewLoots() []Loot {
	panic("implement me")
}

func (m *MediumDifficulty) getNewGameMap() *Map {
	panic("implement me")
}

type HardDifficulty struct{}

func (h *HardDifficulty) getNewPlayer(playerName string) *Player {
	return &Player{
		Character: Character{
			health:     80,
			attack:     10,
			defense:    5,
			stamina:    3,
			maxStamina: 3,
		},
		playerName: "",
	}
}

func (h *HardDifficulty) getNewEnemies() []CharacterI {
	var enemies []CharacterI
	for i := 0; i < 8; i++ {
		enemies = append(enemies, EnemyFactory(EnemyIdentifier(randomGen(0, 2)),
			AbilityIdentifier(randomGen(0, 2))))
	}
	return enemies
}

func (h *HardDifficulty) getNewLoots() []Loot {
	panic("implement me")
}

func (h *HardDifficulty) getNewGameMap() *Map {
	panic("implement me")
}
