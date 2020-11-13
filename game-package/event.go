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
	index := randomGen(0, 1)
	index = 1
	switch index {
	case 0:

		player.SetHealth(player.GetHealth() + 20)
		fmt.Println("You found some medical pills! +20 HP")
	case 1:
		player.addAttackArtifact(&FlameRing{})
		fmt.Println("You found some Flame Rings!")
	}
	fmt.Println("Press Enter.")
	fmt.Scanln()
}
