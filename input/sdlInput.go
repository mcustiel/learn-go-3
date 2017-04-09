package input

import "github.com/veandco/go-sdl2/sdl"

func KeyboardState() *Input {
	input := NewInput()
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		setInputDataFromSdlEvent(event, input)
	}
	state := sdl.GetKeyboardState()
	if state[sdl.SCANCODE_SPACE] == 1 {
		println(state[sdl.SCANCODE_SPACE])
		input.jump = true
	}
	if state[sdl.SCANCODE_RIGHT] == 1 {
		println(state[sdl.SCANCODE_RIGHT])
		input.right = true
	}
	if state[sdl.SCANCODE_LEFT] == 1 {
		println(state[sdl.SCANCODE_LEFT])
		input.left = true
	}

	return input
}

func setInputDataFromSdlEvent(event sdl.Event, input *Input) {
	switch event.(type) {
	case *sdl.QuitEvent:
		input.exit = true
	}
}
