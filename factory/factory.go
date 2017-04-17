package factory

import (
	"github.com/mcustiel/game/logic"
	"github.com/mcustiel/game/render"
	"github.com/mcustiel/game/timing"
)

func CreateGameLogic() *logic.Game {
	return logic.NewGame(CreateRenderer(), CreateFramerateController())
}

func CreateRenderer() render.Renderer {
	return render.NewSdlRenderer()
}

func CreateFramerateController() timing.FrameRateController {
	return timing.NewSdlFrameRateController(timing.FRAME_RATE)
}
