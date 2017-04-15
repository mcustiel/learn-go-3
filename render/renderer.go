package render

const WINDOW_TITLE string = "Learning Go/SDL"
const WINDOWS_WIDTH, WINDOWS_HEIGHT int = 800, 600
const FRAME_RATE int = 48

type Renderer interface {
	Init(worldWidth int) error
	Screen(xCenter int) Screen
	Terminate()
}

type Screen interface {
	Start()
	Draw(renderable Renderable)
	End()
}

type Renderable interface {
	PositionX() int
	PositionY() int
	Width() int
	Height() int
	Type() int
	UniqueId() int
}
