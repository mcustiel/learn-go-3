package render

import (
	"github.com/mcustiel/game/animation"
)

type Screen interface {
	Start()
	Draw(renderable Renderable, sprite *animation.Sprite)
	End()
}
