package logic

const (
	BLOCK_HEIGHT_PIXELS int = 50
	BLOCK_WIDTH_PIXELS  int = 50
)

type BlockType int

const (
	BLOCK_TYPE_NONE          BlockType = 0
	BLOCK_TYPE_FLOOR         BlockType = 1
	BLOCK_TYPE_FALLING_ROCK  BlockType = 2
	BLOCK_TYPE_UNDERGROUND   BlockType = 3
	BLOCK_TYPE_UNDERWATER    BlockType = 4
	BLOCK_TYPE_WATER_SURFACE BlockType = 5
)

type Block struct {
	Entity
	id        int
	blockType BlockType
	friction  int
}

func (block Block) UniqueId() int {
	return block.id
}

func (block Block) Type() int {
	return int(block.blockType)
}

func (block Block) Friction() int {
	return block.friction
}

func (block Block) Width() int {
	return BLOCK_WIDTH_PIXELS
}

func (block Block) Height() int {
	return BLOCK_HEIGHT_PIXELS
}

func applyBlockLogic(block *Block, game *Game) {
	if distance, _ := Distance(block, game.character); BlockType(block.Type()) == BLOCK_TYPE_FALLING_ROCK && distance < 75 {
		block.accelerationY = 1
	}

	block.speedX = AccelerateX(block)
	block.speedY = AccelerateY(block)

	if block.speedX != 0 || block.speedY != 0 {
		var collisionInfo CollisionInformation = GetCollisionInformation(block.Entity, game.level)

		if collisionInfo.collidesV {
			block.speedY = 0
		}
		if collisionInfo.collidesH {
			block.speedX = 0
		}
		block.posX = collisionInfo.suggestedX
		block.posY = collisionInfo.suggestedY
	}
}
