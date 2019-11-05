package resizer

import (
	"image/jpeg"
	"io"

	"github.com/crazy-genius/resize"
)

// RuntimeResizer provide functionality to resizer image at runtime
type RuntimeResizer struct{}

// NewRuntimeResizer create new instance of Resizer
func NewRuntimeResizer() *RuntimeResizer {
	return &RuntimeResizer{}
}

// Resize resize imge
func (r RuntimeResizer) Resize(reader io.Reader, out io.Writer, rc ResizeConfig) error {

	img, err := jpeg.Decode(reader)
	if err != nil {
		return err
	}

	resized := resize.Resize(rc.Dimenstions.Width, rc.Dimenstions.Height, img, resize.Lanczos3)

	if err := jpeg.Encode(out, resized, nil); err != nil {
		return err
	}

	return nil
}
