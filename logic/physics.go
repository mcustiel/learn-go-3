package logic

import "math"

type PhysicObject interface {
	PositionX() int
	PositionY() int
	SpeedX() int
	SpeedY() int
	AccelerationX() int
	AccelerationY() int
	Friction() int
	Width() int
	Height() int
	Collides(X int, Y int, W int, H int) bool
	GetNextX() int
	GetNextY() int
	GetHorizontalPositionNextTo(otherObject PhysicObject) int
	GetVerticalPositionNextTo(otherObject PhysicObject) int
}

type CollisionInformation struct {
	collidingsBlock map[int]*Block
	collidesH       bool
	collidesV       bool
	suggestedX      int
	suggestedY      int
}

func GetNewPosition(object PhysicObject) (int, int) {
	return GetNewX(object), GetNewY(object)
}

func GetNewX(object PhysicObject) int {
	return object.PositionX() + object.SpeedX()
}

func GetNewY(object PhysicObject) int {
	return object.PositionY() + object.SpeedY()
}

func AccelerateX(object PhysicObject) int {
	return object.SpeedX() + object.AccelerationX()
}

func AccelerateY(object PhysicObject) int {
	return object.SpeedY() + object.AccelerationY()
}

func Distance(object1 PhysicObject, object2 PhysicObject) (int, int) {
	return int(math.Abs(float64(object1.PositionX() - object2.PositionX()))),
		int(math.Abs(float64(object1.PositionY() - object2.PositionY())))
}

func GetCollisionInformation(object Entity, level Level) CollisionInformation {
	var newX, newY int = object.GetNextX(), object.GetNextY()
	var collidesH, collidesV bool = false, false
	var block *Block
	var collidingBlocks map[int]*Block = make(map[int]*Block)
	var collisionsFound int = 0
	for iterator := level.Iterator(); iterator.HasNext(); iterator.Next() {
		block = iterator.Get()
		if block != nil && block.Entity != object && block.solid && block.Collides(newX, newY, object.width, object.height) {
			collidingBlocks[collisionsFound] = block
			collisionsFound++
			if block.Collides(newX, object.posY, object.width, object.height) {
				newX = object.GetHorizontalPositionNextTo(block.Entity)
				collidesH = true
			}
			if block.Collides(object.posX, newY, object.width, object.height) {
				newY = object.GetVerticalPositionNextTo(block.Entity)
				collidesV = true
			}
		}
	}

	return CollisionInformation{collidingBlocks, collidesH, collidesV, newX, newY}
}
