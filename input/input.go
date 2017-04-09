package input

const CODE_EXIT byte = 0
const CODE_JUMP byte = 1
const CODE_RIGHT byte = 2
const CODE_LEFT byte = 3
const CODE_INVALID byte = 4

type Input struct {
	jump    bool
	right   bool
	left    bool
	invalid bool
	exit    bool
}

type InputStateAccessor interface {
	IsExit() bool
	IsValid() bool
	JumpPressed() bool
	LeftPressed() bool
	RightPressed() bool
}

func (inputData Input) IsExit() bool {
	return inputData.exit
}

func (inputData Input) IsValid() bool {
	return !inputData.invalid
}

func (inputData Input) JumpPressed() bool {
	return inputData.jump
}

func (inputData Input) LeftPressed() bool {
	return inputData.left
}

func (inputData Input) RightPressed() bool {
	return inputData.right
}

func NewInput() *Input {
	input := new(Input)
	input.invalid = false
	input.exit = false
	input.jump = false
	input.left = false
	input.right = false
	return input
}
