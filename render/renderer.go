package render

import "github.com/mcustiel/game/logic"

const WINDOW_TITLE string = "Go-SDL2 Events"
const WINDOWS_WIDTH, WINDOWS_HEIGHT int = 800, 600

type Renderer interface {
	Init() error
	Render(game *logic.Game)
	Terminate()
}

type Renderable interface {
	PositionX() int
	PositionY() int
	SpeedX() int
	SpeedY() int
	Width() int
	Heigth() int
}
