package logic

type Character struct {
	Entity
	jumping bool
}

func (character Character) UniqueId() int {
	return -1
}

func (character Character) Type() int {
	return -1
}

func (character Character) isJumping() bool {
	return character.jumping
}

func initCharacter() *Character {
	character := new(Character)
	character.posX = 0 * BLOCK_WIDTH_PIXELS
	character.posY = 3 * BLOCK_HEIGHT_PIXELS
	character.jumping = true
	character.speedX = 0
	character.speedY = 0
	character.accelerationX = 0
	character.accelerationY = 1
	character.width = BLOCK_WIDTH_PIXELS
	character.height = BLOCK_HEIGHT_PIXELS
	character.friction = 0
	return character
}

func applyCharacterLogic(game *Game) {
	game.character.speedX = AccelerateX(game.character)
	game.character.speedY = AccelerateY(game.character)

	var newX, newY int = GetNewPosition(game.character)
	if newX < 0 {
		newX = 0
	} else if newX > LOGICAL_WIDTH*BLOCK_WIDTH_PIXELS {
		newX = (LOGICAL_WIDTH * BLOCK_WIDTH_PIXELS) - game.character.width
	}
	if newY < 0 {
		newY = 0
	} else if newY > LOGICAL_HEIGHT*BLOCK_HEIGHT_PIXELS {
		newX = (LOGICAL_HEIGHT * BLOCK_HEIGHT_PIXELS) - game.character.height
	}

	var collidingBlocks map[int]*Block = make(map[int]*Block)
	var collitionsFound int = 0
	var collidesH, collidesV bool = false, false

	for xLevel := 0; xLevel < LOGICAL_WIDTH; xLevel++ {
		for yLevel := 0; yLevel < LOGICAL_HEIGHT; yLevel++ {
			block := game.level[yLevel][xLevel]
			if block != nil && block.Collides(newX, newY, game.character.width, game.character.height) {
				println("Character collides with block ", block.id, block.posY)
				collidingBlocks[collitionsFound] = block

				if block.Collides(newX, game.character.posY, game.character.width, game.character.height) {
					if game.character.speedX > 0 {
						newX = block.posX - game.character.width
					} else {
						newX = block.posX + block.width
					}
					collidesH = true
				}
				if block.Collides(game.character.posX, newY, game.character.width, game.character.height) {
					if game.character.speedY > 0 {
						newY = block.posY - game.character.height
					} else {
						newY = block.posY + block.height
					}
					collidesV = true
				}

				collitionsFound++
			}
		}
	}

	if collidesV {
		println("Collides vertically")
		game.character.speedY = 0
		game.character.jumping = false
	}
	if collidesH {
		println("Collides horizontally")
		game.character.speedX = 0
	}

	println("Character new position ", newX, newY)

	game.character.posX = newX
	game.character.posY = newY

}
