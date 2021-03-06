package main

import (
	"fmt"
	g "golang-project/game-package"
)

func main() {
	fmt.Println("Welcome to the club buddie :)")
	fmt.Println("My name is Van Darkholme. What is your name my sweetie ?)")
	var playerName string
	_, _ = fmt.Scanln(&playerName)
	fmt.Printf("Wow, that's beautiful name, my %v\n", playerName)
	var difficulty string
	for {
		fmt.Println("If you come here, do you wanna play hard with me ? Choose [easy, medium, hard]")
		_, _ = fmt.Scanln(&difficulty)
		myGame := g.NewGame(playerName, difficulty)
		res := myGame.StartGame()
		fmt.Println(res)
		var ans string
		fmt.Println("Wanna play again? [y/n]")
		for {
			_, _ = fmt.Scanln(&ans)
			if ans == "y" || ans == "n" {
				break
			}
		}
		if ans == "n" {
			fmt.Println("Goodbye!")
			break
		}
	}
}
