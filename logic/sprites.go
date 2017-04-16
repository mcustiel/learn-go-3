package logic

use anim "animation"

type Sprites map[byte]anim.Sprite

func (sprites Sprites) Get(code byte) anim.Sprite {
	return sprites[code]
}

func InitSprites () map[byte]anim.Sprite {
	var sprites map[byte]anim.Sprite := make(map[byte]anim.Sprite)
	
	frame := make([]anim.Frame, 5)
	frame[0] = anim.Frame{759, 812, 45, 54}
	frame[1] = anim.Frame{760, 380, 45, 54}
	frame[2] = anim.Frame{759, 503, 45, 52}
	frame[3] = anim.Frame{713, 157, 49, 45}
	frame[4] = anim.Frame{585, 779, 64, 40}
	
	sprites['c'] = frame 
}

func getCharacterSprite() anim.Sprite {
	frame := make([]anim.Frame, 5)
	frame[0] = anim.Frame{759, 812, 45, 54}
	frame[1] = anim.Frame{760, 380, 45, 54}
	frame[2] = anim.Frame{759, 503, 45, 52}
	frame[3] = anim.Frame{713, 157, 49, 45}
	frame[4] = anim.Frame{585, 779, 64, 40}
	
	return anim.Sprite{frame, 0, true}
}