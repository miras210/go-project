package game_package

type Game struct {
	difficulty  Difficulty
	player      *Player                     //Player
	enemies     []CharacterI                //[]Enemy
	loots       []Loot                      //[]Loot
	display     []interface{}               //[]Display
	gameMap     *Map                        //GameMap
	coordinates map[interface{}]*Coordinate //directAccess to the coordinates of object (for movement, attack)
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
	//TODO insertion of the objects to the gameMap and coordinates
	game.difficulty = diff
	return game
}

func (g *Game) StartGame() string {
	move, attack, loot := &MoveAction{}, &AttackAction{}, &LootAction{}
	var result string
	for {
		var input string = "" //TODO implement get Input
		var isValid bool = true
		switch input {
		case "move":
			move.visitForPlayer(g.player)
		case "attack":
			attack.visitForPlayer(g.player)
		case "loot":
			loot.visitForPlayer(g.player)
		case "skip":
			g.player.SetStamina(0)
		default:
			isValid = false
		}
		if !isValid {
			break
		}
		if g.player.GetStamina() > 0 {
			break
		}
		//enemy take turn
		for _, enemy := range g.enemies {
			attack.visitForEnemy(enemy)
			move.visitForEnemy(enemy)
		}
		g.player.SetStamina(g.player.GetMaxStamina())
		if g.player.GetHealth() <= 0 {
			result = "lose"
			break
		}
		if len(g.enemies) == 0 {
			result = "win"
			break
		}
	}
	return result
}

// ***************************************
