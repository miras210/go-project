package game_package

// DIFFICULTY STRATEGY PATTERN

type Difficulty interface {
	getNewPlayer(string) *Player // Player struct
	getNewGameMap() [][]rune
	getEnemy() CharacterI
	getEventNumber() int
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
		return &EasyDifficulty{}
	}
}

type EasyDifficulty struct{}

func (e *EasyDifficulty) getEventNumber() int {
	return 8
}

func (e *EasyDifficulty) getEnemy() CharacterI {
	return EnemyFactory(EnemyIdentifier(randomGen(0, 2)),
		AbilityIdentifier(-1))
}

func (e *EasyDifficulty) getNewPlayer(playerName string) *Player {
	return &Player{
		Character: Character{
			health:  100,
			attack:  20,
			defense: 10},
		playerName: playerName,
		position: &Coordinate{
			x: 1,
			y: 4,
		},
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

func (e *EasyDifficulty) getNewGameMap() [][]rune {
	myTerrainMap := [][]rune{
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', 'E', '#', '#', ' ', '#', ' ', ' ', ' ', '#'},
		{'#', ' ', ' ', ' ', ' ', '#', ' ', '#', ' ', '#'},
		{'#', 'E', '#', '#', ' ', ' ', ' ', '#', 'E', '#'},
		{'#', '#', '#', '#', ' ', '#', '#', '#', '#', '#'},
		{'#', ' ', 'E', '#', ' ', '#', ' ', 'E', ' ', '#'},
		{'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#'},
		{'#', 'E', ' ', '#', ' ', '#', ' ', ' ', 'E', '#'},
		{'#', '#', '#', '#', '#', '#', 'E', ' ', ' ', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
	}
	return myTerrainMap
}

type MediumDifficulty struct{}

func (m *MediumDifficulty) getEnemy() CharacterI {
	return EnemyFactory(EnemyIdentifier(randomGen(0, 2)),
		AbilityIdentifier(randomGen(1, 2)))
}

func (m *MediumDifficulty) getNewPlayer(playerName string) *Player {
	return &Player{
		Character: Character{
			health:  80,
			attack:  25,
			defense: 10,
		},
		playerName: playerName,
		position:   &Coordinate{1, 7},
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

func (m *MediumDifficulty) getEventNumber() int {
	return 14
}
func (m *MediumDifficulty) getNewGameMap() [][]rune {
	myTerrainMap := [][]rune{
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', ' ', ' ', 'E', '#', ' ', ' ', ' ', ' ', 'E', '#', 'E', ' ', ' ', '#'},
		{'#', 'E', ' ', ' ', '#', ' ', ' ', ' ', ' ', ' ', '#', ' ', ' ', ' ', '#'},
		{'#', ' ', ' ', ' ', '#', ' ', '#', '#', '#', ' ', '#', ' ', ' ', 'E', '#'},
		{'#', '#', ' ', '#', '#', ' ', ' ', '#', 'E', ' ', '#', '#', ' ', '#', '#'},
		{'#', ' ', ' ', ' ', ' ', ' ', ' ', '#', ' ', ' ', ' ', ' ', ' ', ' ', '#'},
		{'#', ' ', ' ', ' ', ' ', ' ', 'E', '#', ' ', ' ', ' ', ' ', ' ', ' ', '#'},
		{'#', ' ', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', 'E', '#'},
		{'#', ' ', ' ', ' ', ' ', ' ', ' ', '#', 'E', ' ', ' ', ' ', ' ', ' ', '#'},
		{'#', ' ', ' ', ' ', ' ', ' ', ' ', '#', ' ', ' ', ' ', ' ', ' ', ' ', '#'},
		{'#', '#', ' ', '#', '#', ' ', 'E', '#', ' ', ' ', '#', '#', ' ', '#', '#'},
		{'#', 'E', ' ', ' ', '#', ' ', '#', '#', '#', ' ', '#', ' ', ' ', ' ', '#'},
		{'#', ' ', ' ', ' ', '#', ' ', ' ', ' ', ' ', ' ', '#', ' ', ' ', ' ', '#'},
		{'#', ' ', ' ', 'E', '#', ' ', ' ', ' ', 'E', ' ', '#', ' ', ' ', 'E', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
	}
	return myTerrainMap
}

type HardDifficulty struct{}

func (m *HardDifficulty) getEventNumber() int {
	return 27
}

func (h *HardDifficulty) getNewPlayer(playerName string) *Player {
	return &Player{
		Character: Character{
			health:  80,
			attack:  10,
			defense: 5,
		},
		playerName: playerName,
		position: &Coordinate{
			x: 3,
			y: 8,
		},
	}
}

func (h *HardDifficulty) getEnemy() CharacterI {
	return EnemyFactory(EnemyIdentifier(randomGen(0, 2)),
		AbilityIdentifier(randomGen(0, 1)))
}

func (h *HardDifficulty) getNewEnemies() []CharacterI {
	var enemies []CharacterI
	for i := 0; i < 8; i++ {
		enemies = append(enemies, EnemyFactory(EnemyIdentifier(randomGen(0, 2)),
			AbilityIdentifier(randomGen(0, 1))))
	}
	return enemies
}

func (h *HardDifficulty) getNewGameMap() [][]rune {
	myTerrainMap := [][]rune{
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '#', 'E', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#', ' ', 'E', ' ', ' ', ' ', 'E', ' ', '#', '#'},
		{'#', ' ', '#', ' ', ' ', ' ', 'E', ' ', ' ', ' ', '#', ' ', ' ', ' ', ' ', ' ', ' ', '#', 'E', '#'},
		{'#', ' ', ' ', '#', '#', ' ', ' ', ' ', ' ', ' ', '#', ' ', ' ', ' ', ' ', '#', '#', ' ', ' ', '#'},
		{'#', ' ', ' ', ' ', 'E', '#', ' ', ' ', ' ', ' ', '#', ' ', ' ', ' ', '#', ' ', ' ', ' ', ' ', '#'},
		{'#', 'E', ' ', ' ', ' ', ' ', '#', ' ', ' ', ' ', '#', ' ', ' ', '#', ' ', ' ', ' ', ' ', ' ', '#'},
		{'#', ' ', ' ', ' ', ' ', ' ', '#', ' ', ' ', ' ', ' ', '#', ' ', '#', ' ', ' ', ' ', ' ', '#', '#'},
		{'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'E', ' ', ' ', ' ', ' ', 'E', ' ', 'E', '#', ' ', '#'},
		{'#', ' ', '#', ' ', ' ', ' ', '#', ' ', ' ', ' ', ' ', ' ', 'E', '#', ' ', ' ', ' ', '#', ' ', '#'},
		{'#', 'E', '#', ' ', ' ', ' ', '#', ' ', ' ', ' ', ' ', ' ', ' ', '#', ' ', ' ', ' ', '#', ' ', '#'},
		{'#', ' ', '#', ' ', 'E', ' ', '#', '#', ' ', ' ', ' ', ' ', ' ', '#', ' ', ' ', ' ', ' ', ' ', '#'},
		{'#', '#', ' ', ' ', '#', '#', 'E', ' ', ' ', '#', '#', ' ', '#', ' ', '#', '#', ' ', '#', ' ', '#'},
		{'#', ' ', ' ', '#', ' ', ' ', ' ', ' ', ' ', '#', ' ', ' ', ' ', ' ', ' ', ' ', '#', ' ', 'E', '#'},
		{'#', 'E', ' ', '#', ' ', ' ', ' ', 'E', ' ', '#', 'E', ' ', ' ', ' ', ' ', 'E', '#', ' ', ' ', '#'},
		{'#', ' ', ' ', '#', ' ', ' ', ' ', ' ', ' ', '#', ' ', ' ', ' ', 'E', ' ', ' ', '#', ' ', ' ', '#'},
		{'#', ' ', '#', ' ', ' ', 'E', ' ', ' ', ' ', 'E', '#', ' ', ' ', ' ', ' ', ' ', ' ', '#', ' ', '#'},
		{'#', '#', ' ', ' ', ' ', ' ', ' ', 'E', ' ', ' ', '#', ' ', 'E', ' ', ' ', 'E', ' ', ' ', '#', '#'},
		{'#', '#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#', '#'},
		{'#', '#', ' ', 'E', ' ', ' ', ' ', ' ', ' ', ' ', '#', ' ', ' ', ' ', ' ', ' ', ' ', 'E', '#', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
	}
	return myTerrainMap
}
