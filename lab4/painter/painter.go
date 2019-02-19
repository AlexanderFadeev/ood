package painter

import (
	"github.com/AlexanderFadeev/ood/lab4/canvas"
	"github.com/AlexanderFadeev/ood/lab4/picture_draft"
)

type Painter interface {
	Paint(picture_draft.PictureDraft, canvas.Canvas)
}

type painter struct{}

func New() Painter {
	return new(painter)
}

func (painter) Paint(draft picture_draft.PictureDraft, canvas canvas.Canvas) {
	for _, shape := range draft {
		shape.Draw(canvas)
	}
}
