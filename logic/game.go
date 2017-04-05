package logic

import "github.com/mcustiel/game/input"

type Game struct {
	Level      Level
	Character  *Character
	GameWidth  int
	GameHeight int
}

func CreateGameState() *Game {
	game := new(Game)
	game.Level = initLevel()
	game.Character = initCharacter()
	return game
}

func (game *Game) ApplyInput(gameInput *input.Input) {
	if gameInput.JumpPressed() && !game.Character.Jumping {
		game.Character.SetSpeedY(-16)
		game.Character.Jumping = true
	}
	if gameInput.LeftPressed() {
		game.Character.SetSpeedX(-4)
	} else {
		game.Character.SetSpeedX(0)
	}
	if gameInput.RightPressed() {
		game.Character.SetSpeedX(4)
	}
}
