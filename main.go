package main

import (
	"fmt"

	"github.com/mcustiel/game/factory"
)

func main() {
	fmt.Println("Starting game")
	game := factory.CreateGameLogic()
	game.Play()
}
