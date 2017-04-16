package timing

import "github.com/veandco/go-sdl2/sdl"

type SdlFramerateController int

func NewSdlFrameRateController(framesPerSecond int) SdlFramerateController {
	var controller SdlFramerateController = SdlFramerateController(framesPerSecond)
	return controller
}

func (frameRate SdlFramerateController) Wait() {
	sdl.Delay(uint32(1000 / frameRate))
}
