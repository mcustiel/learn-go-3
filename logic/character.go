package logic

type Character struct {
	posX    int
	posY    int
	speedX  int
	speedY  int
	Jumping bool
}

func (character *Character) SetPositionX(posX int) {
	character.posX = posX
}

func (character *Character) SetPositionY(posY int) {
	character.posY = posY
}

func (character *Character) SetSpeedX(speedX int) {
	character.speedX = speedX
}

func (character *Character) SetSpeedY(speedY int) {
	character.speedY = speedY
}

func (character Character) IsMobile() bool {
	return true
}

func (character Character) PositionX() int {
	return character.posX
}

func (character Character) PositionY() int {
	return character.posY
}

func (character Character) SpeedX() int {
	return character.speedX
}

func (character Character) SpeedY() int {
	return character.speedY
}

func (character Character) Width() int {
	return BLOCK_WIDTH_PIXELS
}

func (character Character) Height() int {
	return BLOCK_HEIGHT_PIXELS
}

func initCharacter() *Character {
	character := new(Character)
	character.posX = 0 * BLOCK_WIDTH_PIXELS
	character.posY = 3 * BLOCK_HEIGHT_PIXELS
	character.Jumping = true
	character.speedX = 0
	character.speedY = 0
	return character
}
