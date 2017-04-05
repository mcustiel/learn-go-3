package main

import (
	"fmt"

	"os"

	"github.com/mcustiel/game/factory"
	"github.com/mcustiel/game/input"
	"github.com/mcustiel/game/physics"
)

func main() {
	fmt.Println("Hello world")

	renderer := factory.CreateRenderer()
	game := factory.CreateGameLogic()
	err := renderer.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		panic(err)
	}

	var play = true
	for renderer.Render(game); play; {
		keyboardState := input.KeyboardState()

		game.ApplyInput(keyboardState)
		if keyboardState != nil {
			if keyboardState.IsExit() {
				play = false
			}
		}

		physics.Apply(game)
		renderer.Render(game)
	}
	renderer.Terminate()

	//	fmt.Println("Hello world")
	//	os.Exit(0)
	//	var gameLoader gameLoader.GameLoader = factory.GetGameLoader()

	//	var game game.GameLogic = factory.GetGameLogic()
	//	var physics game.Physics = factory.GetPhysics()
	//	var collisionHandler collision.CollisionHandler = factory.GetCollisionHandler()
	//	var renderer game.Renderer = factory.GetRenderer()
	//	// var character game.Character                    = factory.GetCharacter()
	//	// var currentLevel int8 = 0
	//	// var levelMap []int8 = gameLoader.LoadLevel(currentLevel)
	//	var gameState game.GameState = game.init()
	//	var playing bool = true
	//	var input input.Input

	//	for playing {
	//		input = inputManager.GetCurrentInput()
	//		playing = game.parseInput(input, gameState)
	//		physics.applyPhysics(gameState)
	//		collisionHandler.detectCollision(gameState)
	//		game.applyGameLogic(gameState)
	//		renderer.render(gameState)
	//	}

}
