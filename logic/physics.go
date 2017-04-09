package logic

import "math"

type CollitionDirection int8

const (
	DIRECTION_BOTH       CollitionDirection = 0
	DIRECTION_VERTICAL   CollitionDirection = -1
	DIRECTION_HORIZONTAL CollitionDirection = 1
)

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
	return int(math.Abs(float64(object1.PositionX() - object2.PositionX()))), int(math.Abs(float64(object1.PositionY() - object2.PositionY())))
}
