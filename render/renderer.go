package render

const WINDOW_TITLE string = "Learning Go/SDL"
const WINDOWS_WIDTH, WINDOWS_HEIGHT int = 800, 600

type Renderer interface {
	Init(worldWidth int) error
	Screen(xCenter int) Screen
	Terminate()
}

type Renderable interface {
	PositionX() int
	PositionY() int
	Width() int
	Height() int
}
