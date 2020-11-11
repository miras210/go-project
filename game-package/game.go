package game_package

type Game struct {
	difficulty Difficulty
	player     interface{}   //Player
	enemies    []interface{} //[]Enemy
	loots      []interface{} //[]Loot
	display    []interface{} //[]Display
	gameMap    interface{}   //GameMap
}

// The only two Public functions are below
// ***************************************
func NewGame(playerName, difficulty string) *Game {
	// TODO Implement game initialization
	game := &Game{}
	diff := newDifficulty(difficulty)
	game.gameMap = diff.getNewGameMap()
	game.player = diff.getNewPlayer(playerName)
	game.enemies = diff.getNewEnemies()
	game.loots = diff.getNewLoots()
	game.difficulty = diff
	return game
}

func (g *Game) StartGame() {
	for {

	}
}

// ***************************************
