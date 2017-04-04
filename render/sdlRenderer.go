package render

import "github.com/veandco/go-sdl2/sdl"
import "github.com/mcustiel/game/logic"

const WINDOW_TITLE string = "Go-SDL2 Events"
const WINDOWS_WIDTH, WINDOWS_HEIGHT int = 800, 600

type SdlDisplay struct {
	window      *sdl.Window
	initialized bool
	renderer    *sdl.Renderer
}

type Renderer interface {
	Init() error
	Render(game *logic.Game)
	Terminate()
}

func Create() *SdlDisplay {
	renderer := new(SdlDisplay)
	renderer.initialized = false
	return renderer
}

func (display *SdlDisplay) Render(game *logic.Game) {
	xStart := game.Character.PositionX() - 6
	if xStart < 0 {
		xStart = 0
	}
	xEnd := game.Character.PositionY() + 6
	if xEnd >= logic.LOGICAL_WIDTH {
		xEnd = logic.LOGICAL_WIDTH - 1
	}
	for x := xStart; x < xEnd; x++ {
		for y := 0; y < 12; y++ {
			block := game.Level[y][x]
			if block != nil {
				rect := sdl.Rect{int32(block.PositionX()), int32(block.PositionY()), int32(block.Width()), int32(block.Height())}
				display.renderer.SetDrawColor(uint8(255-y), uint8(245+y), uint8(235-y), 255)
				display.renderer.DrawRect(&rect)
			}
		}
	}
	display.renderer.Present()
}

func (display *SdlDisplay) Init() error {
	var err error

	sdl.Init(sdl.INIT_EVERYTHING)

	display.window, err = sdl.CreateWindow(WINDOW_TITLE, sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		WINDOWS_WIDTH, WINDOWS_HEIGHT,
		sdl.WINDOW_SHOWN)

	if err != nil {
		return err
	}

	display.renderer, err = sdl.CreateRenderer(display.window, -1, sdl.RENDERER_ACCELERATED)

	return err
}

func (display *SdlDisplay) Terminate() {
	defer display.window.Destroy()
}
