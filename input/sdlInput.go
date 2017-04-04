package input

import "github.com/veandco/go-sdl2/sdl"
import "fmt"

func KeyboardState() *Input {
	if event := sdl.PollEvent(); event != nil {
		return createGameInputFromSdlEvent(event)
	}
	return nil
}

func createGameInputFromSdlEvent(event sdl.Event) *Input {
	var code byte

	switch eventType := event.(type) {
	case *sdl.QuitEvent:
		code = CODE_EXIT
		break
	case *sdl.KeyDownEvent:
		println(eventType.Keysym.Sym, ": ",
			sdl.GetKeyName(sdl.Keycode(eventType.Keysym.Sym)))

		fmt.Printf("K_UP %04x ", sdl.K_UP)
		fmt.Printf("EVENT TYPE %04x ", eventType.Type)
		fmt.Printf("Scancode: %02x Sym: %08x Mod: %04x Unicode: %04x\n",
			eventType.Keysym.Scancode, eventType.Keysym.Sym,
			eventType.Keysym.Mod, eventType.Keysym.Unicode)

		switch eventType.Keysym.Sym {
		case sdl.K_SPACE:
			code = CODE_JUMP
			break
		case sdl.K_LEFT:
			code = CODE_LEFT
			break
		case sdl.K_RIGHT:
			code = CODE_RIGHT
			break
		default:
			code = CODE_INVALID
			break
		}
		break
	default:
		return nil
	}

	return NewInput(code)
}
