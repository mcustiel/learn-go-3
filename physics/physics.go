package physics

import "github.com/mcustiel/game/logic"
import "math"

type PhysicObject interface {
	IsMobile() bool
	PositionX() int
	PositionY() int
	SpeedX() int
	SpeedY() int
	Width() int
	Height() int
	SetPositionX(posX int)
	SetPositionY(posY int)
	SetSpeedX(speedX int)
	SetSpeedY(speedY int)
}

const SPEED_Y int = 1

func Apply(game *logic.Game) {
	for x := 0; x < logic.LOGICAL_WIDTH; x++ {
		for y := 0; y < logic.LOGICAL_HEIGHT; y++ {
			block := game.Level[y][x]
			if block != nil && block.IsMobile() {
				newX := block.PositionX() + block.SpeedX()
				newY := block.PositionY() + block.SpeedY()

				colliding := getColliding(newX, newY, block.Width(), block.Height(), game, x, y)
				collidingX := getColliding(newX, block.PositionY(), block.Width(), block.Height(), game, x, y)
				collidingY := getColliding(block.PositionX(), newY, block.Width(), block.Height(), game, x, y)

				if colliding != nil {
					if collidingX != nil {
						block.SetPositionY(newY)
						block.SetPositionX(collidingX.PositionX() - block.Width())
						block.SetSpeedX(0)
					} else {
						block.SetPositionX(newX)
						block.SetPositionY(collidingY.PositionY() - block.Height())
						block.SetSpeedY(0)
					}
				} else {
					block.SetPositionX(newX)
					block.SetPositionY(newY)
				}

				block.SetSpeedY(block.SpeedY() + SPEED_Y)
			}
		}
	}

	newX := game.Character.PositionX() + game.Character.SpeedX()
	newY := game.Character.PositionY() + game.Character.SpeedY()
	colliding := getColliding(newX, newY, game.Character.Width(), game.Character.Height(), game, -1, -1)
	collidingX := getColliding(newX, game.Character.PositionY(), game.Character.Width(), game.Character.Height(), game, -1, -1)
	collidingY := getColliding(game.Character.PositionX(), newY, game.Character.Width(), game.Character.Height(), game, -1, -1)

	if colliding != nil {
		if collidingX != nil {
			println("character is colliding in x")
			if collidingY == nil {
				game.Character.SetPositionY(newY)
			}
			if game.Character.SpeedX() > 0 {
				game.Character.SetPositionX(collidingX.PositionX() - game.Character.Width())
			} else if game.Character.SpeedX() < 0 {
				game.Character.SetPositionX(collidingX.PositionX() + collidingX.Width())
			}
			game.Character.SetSpeedX(0)
		}
		if collidingY != nil {
			println("character is colliding in y")
			if collidingX == nil {
				game.Character.SetPositionX(newX)
			}
			if game.Character.SpeedY() > 0 {
				game.Character.SetPositionY(collidingY.PositionY() - game.Character.Height())
				if game.Character.Jumping {
					game.Character.Jumping = false
				}
			} else if game.Character.SpeedY() < 0 {
				game.Character.SetPositionY(collidingY.PositionY() + collidingY.Height())
			}
			game.Character.SetSpeedY(0)
		}
	} else {
		game.Character.SetPositionX(newX)
		game.Character.SetPositionY(newY)
	}
	game.Character.SetSpeedY(game.Character.SpeedY() + SPEED_Y)
}

func getColliding(xPos int, yPos int, width int, height int, game *logic.Game, objectX int, objectY int) *logic.Block {
	for x := 0; x < logic.LOGICAL_WIDTH; x++ {
		for y := 0; y < logic.LOGICAL_HEIGHT; y++ {
			block := game.Level[y][x]
			if block != nil && (objectX != x || objectY != y) {

				diffX := xPos - block.PositionX()
				diffY := yPos - block.PositionY()

				if int(math.Abs(float64(diffX))) < width && int(math.Abs(float64(diffY))) < height {
					return block
				}
			}
		}
	}
	return nil
}
