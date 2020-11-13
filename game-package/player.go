package game_package

import (
	"fmt"
	"github.com/eiannone/keyboard"
)

type Player struct {
	Character
	playerName string
	position   *Coordinate
}

func (p *Player) String() string {
	return fmt.Sprintf("%v [HP: %v ATK: %v DEF: %v]", p.playerName,
		p.health, p.attack, p.defense)
}

func (c *Player) move(gmap [][]rune) bool {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	_, key, err := keyboard.GetKey()
	if err != nil {
		panic(err)
	}
	for {
		if key == keyboard.KeyArrowLeft {
			if gmap[c.position.x][c.position.y-1] == '#' {
				break
			}
			c.position.y -= 1
			break
		} else if key == keyboard.KeyArrowRight {
			if gmap[c.position.x][c.position.y+1] == '#' {
				break
			}
			c.position.y += 1
			break
		} else if key == keyboard.KeyArrowUp {
			if gmap[c.position.x-1][c.position.y] == '#' {
				break
			}
			c.position.x -= 1
			break
		} else if key == keyboard.KeyArrowDown {
			if gmap[c.position.x+1][c.position.y] == '#' {
				break
			}
			c.position.x += 1
			break
		} else if key == keyboard.KeyEnter {
			break
		}
		if key == keyboard.KeyEsc {
			return false
		}
	}

	return true
}

func (c *Player) Attack(character CharacterI) {
	fmt.Printf("Player %v attacks! ", c)
	resultingDamage := c.attack - character.GetDefense()
	if resultingDamage < 0 {
		resultingDamage = 0
	}
	fmt.Printf("Dealt %v damage!\n", resultingDamage)
	character.SetHealth(character.GetHealth() - resultingDamage)
}
