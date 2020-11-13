package game_package

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type Game struct {
	difficulty      Difficulty
	player          *Player //Player 	//[]Display
	gameMap         [][]rune
	eventCompletion int
	running         bool //GameMap
	//directAccess to the coordinates of object (for movement, attack)
}

// The only two Public functions are below
// ***************************************
func NewGame(playerName, difficulty string) *Game {
	// TODO Implement game initialization
	game := &Game{}
	diff := newDifficulty(difficulty)
	game.gameMap = diff.getNewGameMap()
	game.player = diff.getNewPlayer(playerName)
	//TODO insertion of the objects to the gameMap and coordinates
	game.difficulty = diff
	game.running = true
	game.eventCompletion = 0
	return game
}

func (g *Game) isRunning() bool {
	return g.running
}

func (g *Game) StartGame() string {
	for g.player.isAlive() || g.isRunning() {
		if g.eventCompletion >= g.difficulty.getEventNumber() {
			g.running = false
			continue
		}
		g.display()
		g.player.move(g.gameMap)
		if g.gameMap[g.player.position.x][g.player.position.y] == 'E' {
			eve := randomGen(0, 1)
			if eve == 0 {
				if !eventBattle(g.player, g.difficulty) {
					g.running = false
				}
			} else {
				eventLoot(g.player, g.difficulty)
			}
			g.gameMap[g.player.position.x][g.player.position.y] = '*'
			g.eventCompletion++
		}
	}
	if g.player.isAlive() {
		return "win"
	} else {
		return "lose"
	}
}
func (g *Game) display() {
	//TODO clear the previous screen

	// WORKS ONLY IN TERMINALS BASH / CMD / .exe file
	// Clearing the console
	clear := make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	value, ok := clear[runtime.GOOS]

	if ok {
		value()
	}
	x, y := g.player.position.getCoordinates()
	//TODO Miras, fix rendering here, please
	for a, row := range g.gameMap {
		for b, cell := range row {
			if a == x && b == y {
				fmt.Print(string('P'))
			} else {
				fmt.Print(string(cell))
			}
		}
		fmt.Println()
	}
	fmt.Println(g.player)

}

// ***************************************
