package render

import (
	"github.com/mcustiel/game/animation"
	"github.com/veandco/go-sdl2/sdl"
)

type SdlDisplay struct {
	window      *sdl.Window
	renderer    *sdl.Renderer
	spritesheet *sdl.Texture
	worldWidth  int
}

type SdlScreen struct {
	xPos    int
	yPos    int
	display *SdlDisplay
}

func NewSdlRenderer() *SdlDisplay {
	display := new(SdlDisplay)
	return display
}

func (display *SdlDisplay) Screen(xCenter int) Screen {
	screen := new(SdlScreen)
	screen.display = display
	if xCenter < WINDOWS_WIDTH/2 {
		screen.xPos = 0
	} else if xCenter > display.worldWidth-WINDOWS_WIDTH/2 {
		screen.xPos = display.worldWidth - WINDOWS_WIDTH
	} else {
		screen.xPos = xCenter - WINDOWS_WIDTH/2
	}
	return screen
}

func (screen SdlScreen) Start() {
	screen.display.renderer.SetDrawColor(0x33, 0x33, 0x33, 0xFF)
	screen.display.renderer.Clear()
}

func (screen SdlScreen) End() {
	screen.display.renderer.Present()
}

func (screen SdlScreen) Draw(renderable Renderable, sprite *animation.Sprite) {
	if screen.isOutOfTheScreen(renderable) {
		return
	}

	screenPos := sdl.Rect{
		int32(renderable.PositionX() - screen.xPos),
		int32(renderable.PositionY()),
		int32(renderable.Width()),
		int32(renderable.Height()),
	}

	if sprite == nil {
		screen.display.renderer.SetDrawColor(255, 245, 235, 255)
		screen.display.renderer.DrawRect(&screenPos)
	} else {
		framePos := sdl.Rect{
			sprite.Current().X,
			sprite.Current().Y,
			sprite.Current().W,
			sprite.Current().H,
		}
		screen.display.renderer.Copy(screen.display.spritesheet, &framePos, &screenPos)
	}
}

func (display *SdlDisplay) Init(worldWidth int) error {
	var err error

	sdl.Init(sdl.INIT_EVERYTHING)

	display.worldWidth = worldWidth
	display.window, err = sdl.CreateWindow(WINDOW_TITLE, sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		WINDOWS_WIDTH, WINDOWS_HEIGHT,
		sdl.WINDOW_SHOWN)

	if err != nil {
		return err
	}

	display.renderer, err = sdl.CreateRenderer(display.window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return err
	}
	display.spritesheet, err = loadAssets(display.renderer)
	if err != nil {
		return err
	}

	return err
}

func (display *SdlDisplay) Terminate() {
	display.spritesheet.Destroy()
	display.renderer.Destroy()
	display.window.Destroy()
}

func (screen SdlScreen) isOutOfTheScreen(renderable Renderable) bool {
	return renderable.PositionX() > screen.xPos+WINDOWS_WIDTH ||
		renderable.PositionX()+renderable.Width() < screen.xPos
}
