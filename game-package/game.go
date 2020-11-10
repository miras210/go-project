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
	return nil
}

func (g *Game) StartGame() {

}

// ***************************************
