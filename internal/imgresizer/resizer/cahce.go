package resizer

import "io"

// CacheResizer is Resizer decorator witch cache resized imgs
type CacheResizer struct {
	wrappedResizer Resizer
}

// NewCacheResizer create new cahced resizer
func NewCacheResizer(r Resizer) *CacheResizer {
	return &CacheResizer{
		wrappedResizer: r,
	}
}

// Resize resize imge
func (c CacheResizer) Resize(reader io.Reader, rc ResizeConfig) ([]byte, error) {
	return c.wrappedResizer.Resize(reader, rc)
}
