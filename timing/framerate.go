package timing

const FRAME_RATE int = 48

type FrameRateController interface {
	Wait()
}
