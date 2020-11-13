package game_package

// DIFFICULTY STRATEGY

type Difficulty interface {
	getNewPlayer(string) *Player // Player struct
	getNewEnemies() []CharacterI // []Enemy struct
	// []Loot struct
	getNewGameMap() [][]rune
	getEnemy() CharacterI
	getEventNumber() int // GameMap struct
}

func newDifficulty(difficulty string) Difficulty {
	switch difficulty {
	case "easy":
		return &EasyDifficulty{}
	/*case "medium":
		return &MediumDifficulty{}
	case "hard":
		return &HardDifficulty{}*/
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

func (e *EasyDifficulty) getNewLoots() []Loot {
	var loots []Loot
	loots = append(loots, NewLoot())
	return loots
}

func (e *EasyDifficulty) getNewGameMap() [][]rune {
	myTerrainMap := [][]rune{
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', 'E', '#', '#', '*', '#', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '#', '*', '#', '*', '#'},
		{'#', 'E', '#', '#', '*', '*', '*', '#', 'E', '#'},
		{'#', '#', '#', '#', '*', '#', '#', '#', '#', '#'},
		{'#', '*', 'E', '#', '*', '#', '*', 'E', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', 'E', '*', '#', '*', '#', '*', '*', 'E', '#'},
		{'#', '#', '#', '#', '#', '#', 'E', '*', '*', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
	}
	return myTerrainMap
}

/*type MediumDifficulty struct{}

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


func (m *MediumDifficulty) getNewGameMap() *Map {
	myTerrainMap := [][]rune{
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
	}
	myGameMap := make([][]interface{}, 10)
	for i := range myGameMap {
		myGameMap[i] = make([]interface{}, 10)
	}
	myMap := Map{
		gameMap:    myGameMap,
		terrainMap: myTerrainMap,
	}
	return &myMap
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
		playerName: playerName,
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



func (h *HardDifficulty) getNewGameMap() *Map {
	myTerrainMap := [][]rune{
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
	}

	myGameMap := make([][]interface{}, 10)
	for i := range myGameMap {
		myGameMap[i] = make([]interface{}, 10)
	}
	myMap := Map{
		gameMap:    myGameMap,
		terrainMap: myTerrainMap,
	}
	return &myMap
}
*/
