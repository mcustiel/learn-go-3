package render

type Frame struct {
	x int32
	y int32
	w int32
	h int32
}

type Sprite struct {
	frames      []Frame
	current     int
	mustAdvance bool
}

func (sprite *Sprite) advance() {
	if sprite.mustAdvance {
		sprite.current = (sprite.current + 1) % len(sprite.frames)
	}
	sprite.mustAdvance = !sprite.mustAdvance
}
