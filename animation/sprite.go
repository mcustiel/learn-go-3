package animation

type Sprite struct {
	framesInSheet []Rectangle
	currentFrame  int
	mustAdvance   bool
}

func (sprite *Sprite) Advance() {
	if sprite.mustAdvance {
		sprite.currentFrame = (sprite.currentFrame + 1) % len(sprite.framesInSheet)
	}
	// To start we are advancing at half the framerate of the game
	sprite.mustAdvance = !sprite.mustAdvance
}

func (sprite Sprite) Current() Rectangle {
	return sprite.framesInSheet[sprite.currentFrame]
}

func NewSprite(frames []Rectangle) *Sprite {
	return &Sprite{frames, 0, true}
}
