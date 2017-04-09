package render

import "github.com/veandco/go-sdl2/sdl"

type SdlObjects map[int]*sdl.Rect

type SdlDisplay struct {
	window          *sdl.Window
	renderer        *sdl.Renderer
	characterObject *sdl.Rect
	sdlObjects      SdlObjects
	worldWidth      int
}

type SdlScreen struct {
	xPos       int
	sdlObjects SdlObjects
	renderer   *sdl.Renderer
}

func NewSdlRenderer() *SdlDisplay {
	renderer := new(SdlDisplay)

	return renderer
}

func (display *SdlDisplay) Screen(xCenter int) Screen {
	screen := new(SdlScreen)
	screen.renderer = display.renderer
	screen.sdlObjects = display.sdlObjects
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
	screen.renderer.SetDrawColor(0x33, 0x33, 0x33, 0xFF)
	screen.renderer.Clear()
}

func (screen SdlScreen) End() {
	screen.renderer.Present()
	sdl.Delay(uint32(1000 / FRAME_RATE))
}

func (screen SdlScreen) Draw(renderable Renderable) {
	if renderable.PositionX() > screen.xPos+WINDOWS_WIDTH {

		return
	}
	if renderable.PositionX()+renderable.Width() < screen.xPos {
		return
	}
	if screen.sdlObjects[renderable.UniqueId()] == nil {
		rect := sdl.Rect{int32(renderable.PositionX() - screen.xPos),
			int32(renderable.PositionY()),
			int32(renderable.Width()),
			int32(renderable.Height())}
		screen.sdlObjects[renderable.UniqueId()] = &rect
	} else {
		screen.sdlObjects[renderable.UniqueId()].X = int32(renderable.PositionX() - screen.xPos)
		screen.sdlObjects[renderable.UniqueId()].Y = int32(renderable.PositionY())
	}

	screen.renderer.SetDrawColor(255, 245, 235, 255)
	screen.renderer.DrawRect(screen.sdlObjects[renderable.UniqueId()])
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
	display.sdlObjects = make(SdlObjects)

	return err
}

func (display *SdlDisplay) Terminate() {
	defer display.renderer.Destroy()
	defer display.window.Destroy()
}
