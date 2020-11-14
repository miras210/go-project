package game_package

import "fmt"

func eventBattle(player *Player, difficulty Difficulty) bool {
	enemy := difficulty.getEnemy()
	for player.isAlive() && enemy.isAlive() {
		player.Attack(enemy)
		fmt.Scanln()
		if enemy.GetHealth() > 0 {
			enemy.Attack(player)
		}
		fmt.Scanln()
	}
	var res bool
	if player.isAlive() {
		fmt.Println("You survived!")
		res = true
	} else {
		fmt.Println("You died.")
		res = false
	}
	fmt.Println("Press Enter.")
	fmt.Scanln()
	return res
}

func eventLoot(player *Player, difficulty Difficulty) {
	index := randomGen(1, 100)
	if index >= 1 && index <= 60 {
		loot := NewLoot()
		if loot.getType() == "attack" {
			fmt.Printf("You found %v! +%v ATK\n", loot.String(), loot.getValue())
			player.addAttackArtifact(loot)
		} else {
			fmt.Printf("You found %v! +%v DEF\n", loot.String(), loot.getValue())
			player.addDefenseArtifact(loot)
		}
	} else {
		sav := randomGen(1, 100)
		if sav > 0 && sav <= 40 {
			player.SetHealth(player.GetHealth() + 20)
			fmt.Println("You found some medical pills! +20 HP")
		} else {
			player.SetHealth(player.GetHealth() + 10)
			fmt.Println("You found some bandage! +10 HP")
		}

	}
	/*index := randomGen(0, 1)
	switch index {
	case 0:
		player.SetHealth(player.GetHealth() + 20)
		fmt.Println("You found some medical pills! +20 HP")
	case 1:
		player.addAttackArtifact(&FlameRing{})
		fmt.Println("You found some Flame Ring! +10 +ATK")
	case 2:
		player.addDefenseArtifact(&Armor{})
		fmt.Println("You found some chain Armor! +10 DEF")
	}*/
	fmt.Println("Press Enter.")
	fmt.Scanln()
}
