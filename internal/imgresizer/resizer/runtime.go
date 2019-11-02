package resizer

import (
	"io"
	"io/ioutil"

	"gopkg.in/gographics/imagick.v3/imagick"
)

// RuntimeResizer provide functionality to resizer image at runtime
type RuntimeResizer struct{}

// NewRuntimeResizer create new instance of Resizer
func NewRuntimeResizer() *RuntimeResizer {
	return &RuntimeResizer{}
}

// Resize resize imge
func (r RuntimeResizer) Resize(reader io.Reader, rc ResizeConfig) ([]byte, error) {
	imagick.Initialize()

	defer imagick.Terminate()
	var err error

	mw := imagick.NewMagickWand()

	err = readImageFromReader(reader, mw)
	if err != nil {
		return nil, err
	}

	// Get original size
	currentDimenstions := Dimenstions{
		Height: mw.GetImageHeight(),
		Width:  mw.GetImageWidth(),
	}

	// Calculate half the sized
	dimenstions := preserveAspectRatoDimensions(currentDimenstions, rc.Dimenstions, PreserveWidth)

	// Resize the image using the Lanczos filter
	err = mw.ResizeImage(dimenstions.Width, dimenstions.Height, imagick.FILTER_LANCZOS)
	if err != nil {
		return nil, err
	}

	// Set the compression quality to 95 (high quality = low compression)
	err = mw.SetImageCompressionQuality(rc.Quality)
	if err != nil {
		return nil, err
	}

	return mw.GetImageBlob(), nil
}

func readImageFromReader(reader io.Reader, mw *imagick.MagickWand) error {
	fileBlob, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	return mw.ReadImageBlob(fileBlob)
}
