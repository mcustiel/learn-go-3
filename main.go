package main

import (
	"fmt"

	"github.com/mcustiel/game/factory"
)

func main() {
	fmt.Println("Hello world")
	game := factory.CreateGameLogic()
	game.Play()
}
