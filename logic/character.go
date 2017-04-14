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
	character.solid = true
	return character
}

func characterDiesAfterAction(game *Game) bool {
	game.character.speedX = AccelerateX(game.character)
	game.character.speedY = AccelerateY(game.character)
	var collisionInfo CollisionInformation = GetCollisionInformation(game.character.Entity, game.level)
	if collisionInfo.collidesV {
		game.character.speedY = 0
		game.character.jumping = false
	}
	if collisionInfo.collidesH {
		game.character.speedX = 0
	}
	game.character.posX = collisionInfo.suggestedX
	game.character.posY = collisionInfo.suggestedY

	return collisionInfo.collidesH || collisionInfo.collidesV
}
