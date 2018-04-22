// package for converting image
package con

import (
	"image"
	"io"
)

type DecodeEncoder interface {
	Decoder
	Encoder
}

type Decoder interface {
	Decode(io.Reader) (image.Image, error)
}

type Encoder interface {
	Encode(io.Writer, image.Image) error
}
