package logic

const BLOCK_HEIGHT_PIXELS int = 50
const BLOCK_WIDTH_PIXELS int = 50

const LOGICAL_WIDTH = 100
const LOGICAL_HEIGHT = 12

type Level [][]*Block

type Block struct {
	mobile bool
	posX   int
	posY   int
	speedX int
	speedY int
}

type Renderable interface {
	PositionX() int
	PositionY() int
	SpeedX() int
	SpeedY() int
	Width() int
	Heigth() int
}

type Character struct {
	posX   int
	posY   int
	speedX int
	speedY int
}

type Game struct {
	Level      Level
	Character  *Character
	GameWidth  int
	GameHeight int
}

func CreateGameState() *Game {
	game := new(Game)
	game.Level = initLevel()
	game.Character = initCharacter()
	return game
}

func (character *Character) PositionX() int {
	return character.posX
}

func (character *Character) PositionY() int {
	return character.posY
}

func (character *Character) SpeedX() int {
	return character.speedX
}

func (character *Character) SpeedY() int {
	return character.speedY
}

func (character *Character) Width() int {
	return 40
}

func (character *Character) Height() int {
	return 40
}

func (block *Block) IsMobile() bool {
	return block.mobile
}

func (block *Block) PositionX() int {
	return block.posX
}

func (block *Block) PositionY() int {
	return block.posY
}

func (block *Block) SpeedX() int {
	return block.speedX
}

func (block *Block) SpeedY() int {
	return block.speedY
}

func (block *Block) Width() int {
	return BLOCK_WIDTH_PIXELS
}

func (block *Block) Height() int {
	return BLOCK_HEIGHT_PIXELS
}

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

func initCharacter() *Character {
	character := new(Character)
	character.posX = 0 * BLOCK_WIDTH_PIXELS
	character.posY = 5 * BLOCK_HEIGHT_PIXELS
	return character
}
