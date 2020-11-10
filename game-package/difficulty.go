package game_package

// DIFFICULTY STRATEGY

type Difficulty interface {
	getNewPlayer(string) interface{} // Player struct
	getNewEnemies() []interface{}    // []Enemy struct
	getNewLoots() []interface{}      // []Loot struct
	getNewGameMap() interface{}      // GameMap struct
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

func (e *EasyDifficulty) getNewPlayer(playerName string) interface{} {
	panic("implement me")
}

func (e *EasyDifficulty) getNewEnemies() []interface{} {
	panic("implement me")
}

func (e *EasyDifficulty) getNewLoots() []interface{} {
	panic("implement me")
}

func (e *EasyDifficulty) getNewGameMap() interface{} {
	panic("implement me")
}

type MediumDifficulty struct{}

func (m MediumDifficulty) getNewPlayer(playerName string) interface{} {
	panic("implement me")
}

func (m MediumDifficulty) getNewEnemies() []interface{} {
	panic("implement me")
}

func (m MediumDifficulty) getNewLoots() []interface{} {
	panic("implement me")
}

func (m MediumDifficulty) getNewGameMap() interface{} {
	panic("implement me")
}

type HardDifficulty struct{}

func (h HardDifficulty) getNewPlayer(playerName string) interface{} {
	panic("implement me")
}

func (h HardDifficulty) getNewEnemies() []interface{} {
	panic("implement me")
}

func (h HardDifficulty) getNewLoots() []interface{} {
	panic("implement me")
}

func (h HardDifficulty) getNewGameMap() interface{} {
	panic("implement me")
}
