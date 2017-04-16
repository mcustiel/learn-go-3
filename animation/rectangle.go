package animation

type Rectangle struct {
	x int32
	y int32
	w int32
	h int32
}

func (rect Rectangle) X() int32 {
	return rect.x
}

func (rect Rectangle) Y() int32 {
	return rect.y
}

func (rect Rectangle) W() int32 {
	return rect.w
}

func (rect Rectangle) H() int32 {
	return rect.h
}
