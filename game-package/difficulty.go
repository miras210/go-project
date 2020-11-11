package game_package

// DIFFICULTY STRATEGY

type Difficulty interface {
	getNewPlayer(string) CharacterI // Player struct
	getNewEnemies() []CharacterI    // []Enemy struct
	getNewLoots() []interface{}     // []Loot struct
	getNewGameMap() interface{}     // GameMap struct
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

func (e *EasyDifficulty) getNewPlayer(playerName string) CharacterI {
	return &Character{
		health:     100,
		attack:     20,
		defense:    10,
		stamina:    3,
		maxStamina: 3,
	}
}

func (e *EasyDifficulty) getNewEnemies() []CharacterI {
	panic("implement me")
}

func (e *EasyDifficulty) getNewLoots() []interface{} {
	panic("implement me")
}

func (e *EasyDifficulty) getNewGameMap() interface{} {
	panic("implement me")
}

type MediumDifficulty struct{}

func (m *MediumDifficulty) getNewPlayer(playerName string) CharacterI {
	return &Character{
		health:     80,
		attack:     15,
		defense:    7,
		stamina:    3,
		maxStamina: 3,
	}
}

func (m *MediumDifficulty) getNewEnemies() []CharacterI {
	panic("implement me")
}

func (m *MediumDifficulty) getNewLoots() []interface{} {
	panic("implement me")
}

func (m *MediumDifficulty) getNewGameMap() interface{} {
	panic("implement me")
}

type HardDifficulty struct{}

func (h *HardDifficulty) getNewPlayer(playerName string) CharacterI {
	panic("implement me")
}

func (h *HardDifficulty) getNewEnemies() []CharacterI {
	panic("implement me")
}

func (h *HardDifficulty) getNewLoots() []interface{} {
	panic("implement me")
}

func (h *HardDifficulty) getNewGameMap() interface{} {
	panic("implement me")
}
