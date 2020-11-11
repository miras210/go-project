package game_package

type Game struct {
	difficulty Difficulty
	player     CharacterI    //Player
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
		var input string = "" //TODO implement get Input
		//player makes move
		//move in map
		//attack
		//defense
		//skip = g.Player.SetStamina(0)
		//player ends move
		if g.player.GetStamina() > 0 {
			break
		}
		//enemy take turn
	}
}

// ***************************************
