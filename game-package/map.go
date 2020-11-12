package game_package

type Map struct {
	gameMap    [][]interface{}
	terrainMap [][]rune
}

func (m *Map) setPlayer(c Player) {
	x, y := randomGen(0, len(m.gameMap)), randomGen(0, len(m.gameMap))
	var pCoord *PlayerCoordinate
	for {
		if m.terrainMap[x][y] == '#' || m.gameMap[x][y] != nil {
			x, y = randomGen(0, len(m.gameMap)), randomGen(0, len(m.gameMap))
			continue
		}
		pCoord = &PlayerCoordinate{
			Player: c,
			Coordinate: Coordinate{
				x: x,
				y: y,
			},
		}
		break
	}
	m.gameMap[x][y] = pCoord
}

func (m *Map) setEnemies(c []CharacterI) {
	for _, enemy := range c {
		x, y := randomGen(0, len(m.gameMap)), randomGen(0, len(m.gameMap))
		var eCoord *EnemyCoordinate
		for {
			if m.terrainMap[x][y] == '#' || m.gameMap[x][y] != nil {
				x, y = randomGen(0, len(m.gameMap)), randomGen(0, len(m.gameMap))
				continue
			}
			eCoord = &EnemyCoordinate{
				CharacterI: enemy,
				Coordinate: Coordinate{
					x: x,
					y: y,
				},
			}
			break
		}
		m.gameMap[x][y] = eCoord
	}
}

func (m *Map) setLoots(c []Loot) {
	for _, loot := range c {
		x, y := randomGen(0, len(m.gameMap)), randomGen(0, len(m.gameMap))
		var lCoord *LootCoordinate
		for {
			if m.terrainMap[x][y] == '#' || m.gameMap[x][y] != nil {
				x, y = randomGen(0, len(m.gameMap)), randomGen(0, len(m.gameMap))
				continue
			}
			lCoord = &LootCoordinate{
				Loot: loot,
				Coordinate: Coordinate{
					x: x,
					y: y,
				},
			}
			break
		}
		m.gameMap[x][y] = lCoord
	}
}

type Coordinate struct {
	x, y int
}

func (c *Coordinate) getCoordinates() (int, int) {
	return c.x, c.y
}

type PlayerCoordinate struct {
	Player
	Coordinate
}

type EnemyCoordinate struct {
	CharacterI
	Coordinate
}

type LootCoordinate struct {
	Loot
	Coordinate
}
