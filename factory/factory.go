package factory

import (
	"github.com/mcustiel/game/logic"
	"github.com/mcustiel/game/render"
)

func CreateGameLogic() *logic.Game {
	return logic.NewGame(CreateRenderer())
}

func CreateRenderer() render.Renderer {
	return render.NewSdlRenderer()
}
