package logic

import (
	"fmt"
	"os"

	"github.com/mcustiel/game/input"
	"github.com/mcustiel/game/render"
	"github.com/mcustiel/game/timing"

	"github.com/mcustiel/game/animation"
)

type Game struct {
	level     Level
	character *Character
	renderer  render.Renderer
	sprites   Sprites
	timer     timing.FrameRateController
}

func NewGame(renderer render.Renderer, timer timing.FrameRateController) *Game {
	level := NewLevel()
	initLevel(level)

	game := new(Game)
	game.level = level
	game.character = initCharacter()
	game.renderer = renderer
	game.sprites = InitSprites()
	game.timer = timer

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
		if game.applyActionAndGetResult() {
			running = false
		}
		game.render()
		game.timer.Wait()
	}
	game.renderer.Terminate()
}

func initRenderer(renderer render.Renderer) {
	err := renderer.Init(WORLD_WIDTH)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		panic(err)
	}
}

func (game *Game) applyActionAndGetResult() bool {
	if characterDiesAfterAction(game) && false {
		return true
	}
	applyLevelLogic(game)
	return false
}

func (game *Game) applyInput(gameInput input.InputStateAccessor) {
	if gameInput.JumpPressed() && !game.character.isJumping() {
		game.character.speedY = -20
		game.character.jumping = true
	}
	if gameInput.LeftPressed() {
		if game.character.jumping {
			game.character.speedX = -6
		} else {
			game.character.speedX = -4
		}
	} else if gameInput.RightPressed() {
		if game.character.jumping {
			game.character.speedX = 6
		} else {
			game.character.speedX = 4
		}
	} else {
		game.character.speedX = 0
	}
}

func (game Game) render() {
	screen := game.renderer.Screen(game.character.posX)
	screen.Start()
	var sprite *animation.Sprite

	for iterator := game.level.Iterator(); iterator.HasNext(); iterator.Next() {
		block := iterator.Get()
		if block != nil {
			sprite = game.sprites.Get(getBlockSpriteCode(block))
			screen.Draw(block, sprite)
			sprite.Advance()
		}
	}
	sprite = game.sprites.Get(getCharacterSpriteCode(game.character))
	screen.Draw(game.character, sprite)
	sprite.Advance()
	screen.End()
}

func getBlockSpriteCode(block *Block) byte {
	var code byte
	if BlockType(block.Type()) == BLOCK_TYPE_FALLING_ROCK {
		code = FALLING_ROCK
	} else if BlockType(block.Type()) == BLOCK_TYPE_ROCK {
		code = ROCK
	} else if BlockType(block.Type()) == BLOCK_TYPE_FLOOR {
		code = GROUND
	}
	return code
}

func getCharacterSpriteCode(character *Character) byte {
	var code byte = CHARACTER_STANDING
	if character.jumping {
		code = CHARACTER_JUMPING
	} else if character.speedX != 0 {
		code = CHARACTER_WALKING
	}
	return code
}
