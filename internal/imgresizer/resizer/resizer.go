package resizer

import (
	"io"
	"io/ioutil"

	"gopkg.in/gographics/imagick.v3/imagick"
)

// Resizer provide functionality to resizer image
type Resizer struct {
}

// ResizeConfig represents a desired img resize configuration
type ResizeConfig struct {
	Height uint
	Width  uint
}

// NewResizer create new instance of Resizer
func NewResizer() *Resizer {

	return &Resizer{}
}

// Resize resize imge
func (r *Resizer) Resize(reader io.Reader, configuration ResizeConfig) []byte {
	imagick.Initialize()

	defer imagick.Terminate()
	var err error

	mw := imagick.NewMagickWand()

	err = readImageFromReader(reader, mw)
	if err != nil {
		panic(err)
	}

	// Get original size
	width := mw.GetImageWidth()
	height := mw.GetImageHeight()

	// Calculate half the size
	hWidth := uint(width / 2)
	hHeight := uint(height / 2)

	// Resize the image using the Lanczos filter
	err = mw.ResizeImage(hWidth, hHeight, imagick.FILTER_LANCZOS)
	if err != nil {
		panic(err)
	}

	// Set the compression quality to 95 (high quality = low compression)
	err = mw.SetImageCompressionQuality(95)
	if err != nil {
		panic(err)
	}

	return mw.GetImageBlob()
}

func readImageFromReader(reader io.Reader, mw *imagick.MagickWand) error {
	fileBlob, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	return mw.ReadImageBlob(fileBlob)
}
