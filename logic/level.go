package logic

import "sync"

const LOGICAL_WIDTH = 100
const LOGICAL_HEIGHT = 12

type Level [][]*Block

func initLevel() Level {
	level := make([][]*Block, LOGICAL_HEIGHT, LOGICAL_HEIGHT)
	for i := 0; i < LOGICAL_HEIGHT; i++ {
		level[i] = make([]*Block, LOGICAL_WIDTH, LOGICAL_WIDTH)
	}

	for i := 0; i < LOGICAL_WIDTH; i++ {
		object := new(Block)
		object.accelerationX = 0
		object.accelerationY = 0
		object.blockType = BLOCK_TYPE_FLOOR
		object.friction = -1
		object.id = (LOGICAL_HEIGHT-1)*LOGICAL_WIDTH + i
		object.posX = i * BLOCK_WIDTH_PIXELS
		object.posY = (LOGICAL_HEIGHT - 1) * BLOCK_HEIGHT_PIXELS
		object.speedX = 0
		object.speedY = 0
		object.width = BLOCK_WIDTH_PIXELS
		object.height = BLOCK_HEIGHT_PIXELS

		level[LOGICAL_HEIGHT-1][i] = object
	}

	for i := 5; i < LOGICAL_WIDTH; i += 10 {
		object := new(Block)
		object.accelerationX = 0
		object.accelerationY = 0
		object.blockType = BLOCK_TYPE_FALLING_ROCK
		object.friction = 0
		object.id = i
		object.posX = i * BLOCK_WIDTH_PIXELS
		object.posY = 0
		object.speedX = 0
		object.speedY = 0
		object.width = BLOCK_WIDTH_PIXELS
		object.height = BLOCK_HEIGHT_PIXELS

		level[0][i] = object
	}

	return level
}

func applyLevelLogic(game *Game) {
	waitGroup := new(sync.WaitGroup)
	for xLevel := 0; xLevel < LOGICAL_WIDTH; xLevel++ {
		for yLevel := 0; yLevel < LOGICAL_HEIGHT; yLevel++ {
			block := game.level[yLevel][xLevel]
			if block != nil {
				waitGroup.Add(1)
				go func() {
					applyBlockLogic(block, game)
					waitGroup.Done()
				}()
			}
		}
	}
	waitGroup.Wait()
}
