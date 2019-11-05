package resizer

import (
	"io"
)

// Resizer provide functionality to resizer image
type Resizer interface {
	Resize(reader io.Reader, out io.Writer, rc ResizeConfig) error
}

// ResizeConfig represents a desired img resize configuration
type ResizeConfig struct {
	Dimenstions Dimenstions
	Quality     uint
}

// NewResizer create new instance of Resizer
func NewResizer() Resizer {
	return NewRuntimeResizer()
}
