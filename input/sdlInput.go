package input

import "github.com/veandco/go-sdl2/sdl"

//import "fmt"

func KeyboardState() *Input {
	input := NewInput()
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		setInputDataFromSdlEvent(event, input)
	}
	state := sdl.GetKeyboardState()
	if state[sdl.SCANCODE_UP] == 1 {
		println(state[sdl.SCANCODE_UP])
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
		//	case *sdl.KeyDownEvent:
		//		println(eventType.Keysym.Sym, ": ",
		//			sdl.GetKeyName(sdl.Keycode(eventType.Keysym.Sym)))

		//		fmt.Printf("K_UP %04x ", sdl.K_UP)
		//		fmt.Printf("EVENT TYPE %04x ", eventType.Type)
		//		fmt.Printf("Scancode: %02x Sym: %08x Mod: %04x Unicode: %04x\n",
		//			eventType.Keysym.Scancode, eventType.Keysym.Sym,
		//			eventType.Keysym.Mod, eventType.Keysym.Unicode)

		//		switch eventType.Keysym.Sym {
		//		case sdl.K_SPACE:
		//			input.jump = true
		//			break
		//		case sdl.K_LEFT:
		//			input.left = true
		//			break
		//		case sdl.K_RIGHT:
		//			input.right = true
		//			break
		//		default:
		//			input.invalid = true
		//			break
		//		}
		//		break
	}
}
