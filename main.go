package main

import (
	"fmt"

	"os"

	"github.com/mcustiel/game/input"
	"github.com/veandco/go-sdl2/sdl"
)

var winTitle string = "Go-SDL2 Events"
var winWidth, winHeight int = 800, 600

func main() {
	fmt.Println("Hello world")

	var window *sdl.Window

	var err error

	sdl.Init(sdl.INIT_EVERYTHING)

	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight,
		sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		os.Exit(1)
	}
	var play = true
	for play {
		keyboardState := input.KeyboardState()

		if keyboardState != nil {
			if keyboardState.IsExit() {
				play = false
			}
			fmt.Println(keyboardState.Code())
		}
		//fmt.Println(keyboardState.key)
	}
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
	defer window.Destroy()
}
