package game_package

type Game struct {
	difficulty Difficulty
	player     CharacterI    //Player
	enemies    []CharacterI  //[]Enemy
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
	move, attack, loot := &MoveAction{}, &AttackAction{}, &LootAction{}
	var result string
	for {
		var input string = "" //TODO implement get Input
		//player makes move
		//move in map
		//attack
		//defense
		//status (maybe shows stats of everything on the map?)
		//skip = g.Player.SetStamina(0)
		//player ends move
		if g.player.GetStamina() > 0 {
			break
		}
		//enemy take turn
		for _, enemy := range g.enemies {
			attack.visitForEnemy(&enemy)
			move.visitForEnemy(&enemy)
		}
		if g.player.GetHealth() <= 0 {
			result = "lose"
			break
		}
		if len(g.enemies) == 0 {
			result = "win"
			break
		}
	}
}

// ***************************************
