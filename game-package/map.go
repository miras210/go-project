package game_package

type Map struct {
	gameMap    [][]interface{}
	terrainMap [][]rune
}

type Coordinate struct {
	x, y int
}
