package logic

import (
	"fmt"
	"os"

	"github.com/mcustiel/game/input"
	"github.com/mcustiel/game/render"
)

type Game struct {
	level     Level
	character *Character
	renderer  render.Renderer
}

func NewGame(renderer render.Renderer) *Game {
	game := new(Game)
	game.level = initLevel()
	game.character = initCharacter()
	game.renderer = renderer
	return game
}

func (game Game) Play() {
	var running bool = true
	var inputState input.InputStateAccessor

	initRenderer(game.renderer)

	for running {
		inputState = input.KeyboardState()
		if inputState.IsExit() {
			running = false
		}
		game.applyInput(inputState)
		game.applyLogic()
		game.render()
	}
	game.renderer.Terminate()
}

func initRenderer(renderer render.Renderer) {
	err := renderer.Init(LOGICAL_WIDTH * BLOCK_WIDTH_PIXELS)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		panic(err)
	}
}

func (game *Game) applyLogic() {
	applyCharacterLogic(game)
	println("Character speedY", game.character.speedY)
	applyLevelLogic(game)

	//	for i := 0; i < collisionsFound; i++ {
	//		switch (collidingTypes[i]) {
	//		}
	//	}

}

func (game *Game) applyInput(gameInput input.InputStateAccessor) {
	if gameInput.JumpPressed() {
		println("JumpPressed")
	}
	if gameInput.JumpPressed() && !game.character.isJumping() {
		game.character.speedY = -20
		game.character.jumping = true
	}
	if gameInput.LeftPressed() {
		game.character.speedX = -4
	} else {
		game.character.speedX = 0
	}
	if gameInput.RightPressed() {
		game.character.speedX = 4
	}
}

func (game Game) render() {
	screen := game.renderer.Screen(game.character.posX)
	screen.Start()
	for xLevel := 0; xLevel < LOGICAL_WIDTH; xLevel++ {
		for yLevel := 0; yLevel < LOGICAL_HEIGHT; yLevel++ {
			block := game.level[yLevel][xLevel]
			if block != nil {
				screen.Draw(block)
			}
		}
	}
	screen.Draw(game.character)
	screen.End()
}
