package render

import "github.com/veandco/go-sdl2/sdl"
import "github.com/mcustiel/game/logic"

type SdlObjects [][]*sdl.Rect

type SdlDisplay struct {
	window          *sdl.Window
	initialized     bool
	renderer        *sdl.Renderer
	characterObject *sdl.Rect
	sdlObjects      SdlObjects
}

func Create() *SdlDisplay {
	renderer := new(SdlDisplay)
	renderer.initialized = false
	return renderer
}

func (display *SdlDisplay) Render(game *logic.Game) {
	display.renderer.SetDrawColor(0x33, 0x33, 0x33, 0xFF)
	display.renderer.Clear()
	renderLevelSection(display, game.Level, getRenderStartX(game.Character), getRenderEndX(game.Character))
	renderCharacter(display, game.Character)
	display.renderer.Present()
	sdl.Delay(uint32(1000 / 48))
}

func renderLevelSection(display *SdlDisplay, level logic.Level, xStart int, xEnd int) {
	for x := xStart; x < xEnd; x++ {
		for y := 0; y < 12; y++ {
			block := level[y][x]
			if block != nil {
				renderBlock(display, block, x, y)
			}
		}
	}
}

func renderBlock(display *SdlDisplay, block *logic.Block, x int, y int) {
	if display.sdlObjects[y][x] == nil {
		rect := sdl.Rect{int32(block.PositionX()),
			int32(block.PositionY()),
			int32(block.Width()),
			int32(block.Height())}
		display.sdlObjects[y][x] = &rect
	} else {
		display.sdlObjects[y][x].X = int32(block.PositionX())
		display.sdlObjects[y][x].Y = int32(block.PositionY())
	}
	display.renderer.SetDrawColor(255, 245, 235, 255)
	display.renderer.DrawRect(display.sdlObjects[y][x])
}

func getRenderEndX(character *logic.Character) int {
	xEnd := character.PositionX()/logic.BLOCK_WIDTH_PIXELS + 8
	if xEnd >= logic.LOGICAL_WIDTH {
		xEnd = logic.LOGICAL_WIDTH - 1
	}
	if xEnd < 16 {
		xEnd = 16
	}
	return xEnd
}

func getRenderStartX(character *logic.Character) int {
	xStart := character.PositionX()/logic.BLOCK_WIDTH_PIXELS - 8
	if xStart < 0 {
		xStart = 0
	}
	if xStart >= logic.LOGICAL_WIDTH-16 {
		xStart = logic.LOGICAL_WIDTH - 16
	}
	return xStart
}

func renderCharacter(display *SdlDisplay, character *logic.Character) {
	if display.characterObject == nil {
		rect := sdl.Rect{int32(character.PositionX()),
			int32(character.PositionY()),
			int32(character.Width()),
			int32(character.Height())}
		display.characterObject = &rect
	} else {
		display.characterObject.X = int32(character.PositionX())
		display.characterObject.Y = int32(character.PositionY())
	}
	display.renderer.SetDrawColor(uint8(255), uint8(245), uint8(235), 255)
	display.renderer.DrawRect(display.characterObject)
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
	display.initialized = true
	display.sdlObjects = make([][]*sdl.Rect, 12, 12)
	for i := 0; i < 12; i++ {
		display.sdlObjects[i] = make([]*sdl.Rect, 100, 100)
	}

	return err
}

func (display *SdlDisplay) Terminate() {
	defer display.renderer.Destroy()
	defer display.window.Destroy()
}
