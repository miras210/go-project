package game_package

type Coordinate struct {
	x, y int
}

func (c *Coordinate) getCoordinates() (int, int) {
	return c.x, c.y
}
