package render

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

type SdlObject struct {
	sdlRect *sdl.Rect
	sprite  *Sprite
}

type SdlObjects map[int]*SdlObject

type SdlDisplay struct {
	window      *sdl.Window
	renderer    *sdl.Renderer
	spritesheet *sdl.Texture
	sdlObjects  SdlObjects
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
	sdl.Delay(uint32(1000 / FRAME_RATE))
}

func (screen SdlScreen) Draw(renderable Renderable) {
	if screen.isOutOfTheScreen(renderable) {
		return
	}

	sdlObject := screen.getSdlObjectToDraw(renderable)
	sdlObject.sdlRect.X = int32(renderable.PositionX() - screen.xPos)
	sdlObject.sdlRect.Y = int32(renderable.PositionY())

	if sdlObject.sprite == nil {
		screen.display.renderer.SetDrawColor(255, 245, 235, 255)
		screen.display.renderer.DrawRect(sdlObject.sdlRect)
	} else {
		dest := sdl.Rect{
			sdlObject.sprite.frames[sdlObject.sprite.current].x,
			sdlObject.sprite.frames[sdlObject.sprite.current].y,
			sdlObject.sprite.frames[sdlObject.sprite.current].w,
			sdlObject.sprite.frames[sdlObject.sprite.current].h,
		}
		sdlObject.sprite.advance()
		screen.display.renderer.Copy(screen.display.spritesheet, &dest, sdlObject.sdlRect)
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
	image, err := img.Load("assets/spritesheet_complete.png")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", err)
		return err
	}
	defer image.Free()

	display.spritesheet, err = display.renderer.CreateTextureFromSurface(image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		return err
	}

	display.sdlObjects = make(SdlObjects)

	frame := make([]Frame, 5)
	frame[0] = Frame{759, 812, 45, 54}
	frame[1] = Frame{760, 380, 45, 54}
	frame[2] = Frame{759, 503, 45, 52}
	frame[3] = Frame{713, 157, 49, 45}
	frame[4] = Frame{585, 779, 64, 40}
	display.sdlObjects[-1] = new(SdlObject)
	display.sdlObjects[-1].sprite = &Sprite{frame, 0, true}

	return err
}

func (display *SdlDisplay) Terminate() {
	display.spritesheet.Destroy()
	display.renderer.Destroy()
	display.window.Destroy()
}

func (screen SdlScreen) getSdlObjectToDraw(renderable Renderable) *SdlObject {
	if screen.display.sdlObjects[renderable.UniqueId()] == nil {
		screen.display.sdlObjects[renderable.UniqueId()] = new(SdlObject)
	}
	sdlObject := screen.display.sdlObjects[renderable.UniqueId()]
	if sdlObject.sdlRect == nil {
		rect := sdl.Rect{int32(renderable.PositionX() - screen.xPos),
			int32(renderable.PositionY()),
			int32(renderable.Width()),
			int32(renderable.Height())}
		sdlObject.sdlRect = &rect
	}
	return sdlObject
}

func (screen SdlScreen) isOutOfTheScreen(renderable Renderable) bool {
	return renderable.PositionX() > screen.xPos+WINDOWS_WIDTH ||
		renderable.PositionX()+renderable.Width() < screen.xPos
}
