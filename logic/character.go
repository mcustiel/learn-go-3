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
	var collidingTypes []BlockType = make([]BlockType, 8, 8)
	var collitionsFound int = 0
	var collidesH, collidesV bool = false, false

	for xLevel := 0; xLevel < LOGICAL_WIDTH; xLevel++ {
		for yLevel := 0; yLevel < LOGICAL_HEIGHT; yLevel++ {
			block := game.level[yLevel][xLevel]
			if block != nil && block.Collides(newX, newY, game.character.width, game.character.height) {
				println("Character collides with block ", block.id, block.posY)
				collidingTypes[collitionsFound] = block.blockType

				if block.Collides(newX, game.character.posY, game.character.width, game.character.height) {
					if game.character.speedX > 0 {
						newX = block.posX - game.character.width
					} else {
						newX = block.posX + block.width
					}
					collidesH = true
				} else {
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
