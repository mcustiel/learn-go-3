package logic

const LOGICAL_WIDTH = 100
const LOGICAL_HEIGHT = 12

type Level [][]*Block

func initLevel() Level {
	level := make([][]*Block, 12, 12)
	for i := 0; i < 12; i++ {
		level[i] = make([]*Block, 100, 100)
	}

	for i := 0; i < 100; i++ {
		object := new(Block)
		object.mobile = false
		object.posX = i * BLOCK_WIDTH_PIXELS
		object.posY = 11 * BLOCK_HEIGHT_PIXELS
		level[11][i] = object
	}

	for i := 9; i < 100; i += 10 {
		object := new(Block)
		object.mobile = true
		object.posX = i * BLOCK_WIDTH_PIXELS
		object.posY = 0 * BLOCK_HEIGHT_PIXELS
		level[0][i] = object
	}

	return level
}
