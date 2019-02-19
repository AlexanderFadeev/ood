package shape

import "github.com/AlexanderFadeev/ood/lab7/rect"

type frame struct {
	rect.Rect
}

func NewFrame(rect rect.Rect) *frame {
	return &frame{
		Rect: rect,
	}
}

func (f *frame) GetFrame() *rect.Rect {
	return &f.Rect
}

func (f *frame) SetFrame(frame rect.Rect) {
	if frame.Width() == 0 || frame.Height() == 0 {
		panic("Frame should not have zero width or height")
	}

	f.Rect = frame
}
