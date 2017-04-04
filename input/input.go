package input

const CODE_EXIT byte = 0
const CODE_JUMP byte = 1
const CODE_RIGHT byte = 2
const CODE_LEFT byte = 3
const CODE_INVALID byte = 4

type Input struct {
	key byte
}

type GameInput interface {
	IsExit() bool
	IsValid() bool
	Code() Input
}

func NewInput(c byte) *Input {
	input := new(Input)
	input.key = c
	return input
}

func (inputData Input) IsExit() bool {
	return inputData.key == CODE_EXIT
}

func (inputData Input) IsValid() bool {
	return inputData.key != CODE_INVALID
}

func (inputData Input) Code() byte {
	return inputData.key
}
