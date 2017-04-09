package logic

import "math"

type Entity struct {
	posX          int
	posY          int
	speedX        int
	speedY        int
	accelerationX int
	accelerationY int
	width         int
	height        int
	friction      int
}

func (entity Entity) Friction() int {
	return entity.friction
}

func (entity Entity) PositionX() int {
	return entity.posX
}

func (entity Entity) PositionY() int {
	return entity.posY
}

func (entity Entity) SpeedX() int {
	return entity.speedX
}

func (entity Entity) SpeedY() int {
	return entity.speedY
}

func (entity Entity) AccelerationX() int {
	return entity.accelerationX
}

func (entity Entity) AccelerationY() int {
	return entity.accelerationY
}

func (entity Entity) Width() int {
	return entity.width
}

func (entity Entity) Height() int {
	return entity.height
}

func (entity Entity) Collides(X int, Y int, W int, H int) bool {
	var collidesX bool = entity.posX+entity.width > X && entity.posX < X+W

	var collidesY bool = entity.posY+entity.height > Y && entity.posY < Y+H

	return collidesX && collidesY
}

func (object Entity) collitionDirection(X int, Y int) CollitionDirection {
	distX := math.Abs(float64(object.posX - X))
	distY := math.Abs(float64(object.posY - Y))

	var direction CollitionDirection
	if distX > distY {
		direction = DIRECTION_HORIZONTAL
	} else if distY > distX {
		direction = DIRECTION_VERTICAL
	} else {
		direction = DIRECTION_BOTH
	}
	return direction
}
