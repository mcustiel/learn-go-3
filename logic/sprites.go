package logic

import (
	"github.com/mcustiel/game/animation"
)

const (
	CHARACTER_WALKING  byte = 1
	CHARACTER_STANDING byte = 2
	CHARACTER_JUMPING  byte = 3
	GROUND             byte = 4
	FALLING_ROCK       byte = 5
	ROCK               byte = 6
)

type Sprites map[byte]*animation.Sprite

func (sprites Sprites) Get(code byte) *animation.Sprite {
	return sprites[code]
}

func InitSprites() Sprites {
	var sprites Sprites = make(Sprites)

	sprites[CHARACTER_WALKING] = getCharacterWalkingSprite()
	sprites[CHARACTER_STANDING] = getCharacterStandingSprite()
	sprites[CHARACTER_JUMPING] = getCharacterJumpingSprite()
	sprites[GROUND] = getGroundSprite()
	sprites[FALLING_ROCK] = getFallingBlockSprite()
	sprites[ROCK] = getRockSprite()

	return sprites
}

func getCharacterJumpingSprite() *animation.Sprite {
	frame := make([]animation.Rectangle, 2)
	frame[0] = animation.Rectangle{805, 545, 44, 53}
	frame[1] = animation.Rectangle{764, 0, 44, 54}

	return animation.NewSprite(frame)
}

func getCharacterWalkingSprite() *animation.Sprite {
	frame := make([]animation.Rectangle, 5)
	frame[0] = animation.Rectangle{759, 812, 45, 54}
	frame[1] = animation.Rectangle{760, 380, 45, 54}
	frame[2] = animation.Rectangle{759, 503, 45, 52}
	frame[3] = animation.Rectangle{713, 157, 49, 45}
	frame[4] = animation.Rectangle{585, 779, 64, 40}

	return animation.NewSprite(frame)
}

func getGroundSprite() *animation.Sprite {
	frame := make([]animation.Rectangle, 1)
	frame[0] = animation.Rectangle{520, 0, 64, 64}
	return animation.NewSprite(frame)
}

func getFallingBlockSprite() *animation.Sprite {
	frame := make([]animation.Rectangle, 1)
	frame[0] = animation.Rectangle{455, 914, 64, 64}
	return animation.NewSprite(frame)
}

func getRockSprite() *animation.Sprite {
	frame := make([]animation.Rectangle, 1)
	frame[0] = animation.Rectangle{455, 849, 64, 64}
	return animation.NewSprite(frame)
}

func getCharacterStandingSprite() *animation.Sprite {
	frame := make([]animation.Rectangle, 1)
	frame[0] = animation.Rectangle{762, 203, 45, 54}
	return animation.NewSprite(frame)
}
