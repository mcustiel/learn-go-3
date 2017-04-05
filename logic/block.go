package logic

const BLOCK_HEIGHT_PIXELS int = 50
const BLOCK_WIDTH_PIXELS int = 50

type Block struct {
	mobile bool
	posX   int
	posY   int
	speedX int
	speedY int
}

func (block *Block) SetPositionX(posX int) {
	block.posX = posX
}

func (block *Block) SetPositionY(posY int) {
	block.posY = posY
}

func (block *Block) SetSpeedX(speedX int) {
	block.speedX = speedX
}

func (block *Block) SetSpeedY(speedY int) {
	block.speedY = speedY
}

func (block Block) IsMobile() bool {
	return block.mobile
}

func (block Block) PositionX() int {
	return block.posX
}

func (block Block) PositionY() int {
	return block.posY
}

func (block Block) SpeedX() int {
	return block.speedX
}

func (block Block) SpeedY() int {
	return block.speedY
}

func (block Block) Width() int {
	return BLOCK_WIDTH_PIXELS
}

func (block Block) Height() int {
	return BLOCK_HEIGHT_PIXELS
}
