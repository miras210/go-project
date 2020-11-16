package game_package

import (
	"fmt"
	"github.com/eiannone/keyboard"
)

type Player struct {
	Character
	playerName      string
	position        *Coordinate
	defenseArtifact Loot
	attackArtifact  Loot
}

func (p *Player) addAttackArtifact(a Loot) {
	if p.attackArtifact == nil {
		p.attackArtifact = a
	} else {
		p.attackArtifact.setNext(a)
	}
}
func (p *Player) addDefenseArtifact(d Loot) {
	if p.defenseArtifact == nil {
		p.defenseArtifact = d
	} else {
		p.defenseArtifact.setNext(d)
	}
}

func (p *Player) GetAttack() int {
	if p.attackArtifact != nil {
		return p.attack + p.attackArtifact.getValue()
	} else {
		return p.attack
	}
}
func (p *Player) GetDefense() int {
	if p.defenseArtifact != nil {
		return p.defense + p.defenseArtifact.getValue()
	} else {
		return p.defense
	}
}

func (p *Player) String() string {
	return fmt.Sprintf("%v [HP: %v ATK: %v DEF: %v]", p.playerName,
		p.health, p.GetAttack(), p.GetDefense())
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
	fmt.Printf("Player %v attacks %v! ", c, character)
	coeff := float64(c.GetAttack()) / float64(character.GetDefense())
	if coeff > 1 {
		coeff = 1
	}
	resultingDamage := int(float64(c.GetAttack()) * coeff)
	fmt.Printf("Dealt %v damage!\n", resultingDamage)
	character.SetHealth(character.GetHealth() - resultingDamage)
}
