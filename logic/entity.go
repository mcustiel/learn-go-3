package logic

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
	solid         bool
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

func (object Entity) GetNextX() int {
	var newX int = object.posX + object.speedX
	if newX < 0 {
		newX = 0
	} else if newX > WORLD_WIDTH-object.width {
		newX = WORLD_WIDTH - object.width
	}
	return newX
}

func (object Entity) GetNextY() int {
	var newY int = object.posY + object.speedY
	if newY < 0 {
		newY = 0
	} else if newY > WORLD_HEIGHT-object.height {
		newY = WORLD_HEIGHT - object.height
	}
	return newY
}

func (object Entity) GetHorizontalPositionNextTo(otherObject PhysicObject) int {
	var newX int
	if object.SpeedX() > 0 {
		newX = otherObject.PositionX() - object.Width()
	} else {
		newX = otherObject.PositionX() + otherObject.Width()
	}
	return newX
}

func (object Entity) GetVerticalPositionNextTo(otherObject PhysicObject) int {
	var newY int
	if object.SpeedY() > 0 {
		newY = otherObject.PositionY() - object.Height()
	} else {
		newY = otherObject.PositionY() + otherObject.Height()
	}
	return newY
}
