package animacy

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"io"

	"github.com/soniakeys/quant/median"
)

// Writer generates animation GIF file from given images.
type Writer struct {
	output *gif.GIF
}

// NewWriter returns new instance of Writer
func NewWriter() *Writer {
	return &Writer{output: &gif.GIF{}}
}

// AppendImage appends image
func (w Writer) AppendImage(img image.Image, duration int) {
	palette := median.Quantizer(256).Quantize(make(color.Palette, 0, 256), img)
	paletted := image.NewPaletted(img.Bounds(), palette)
	draw.FloydSteinberg.Draw(paletted, img.Bounds(), img, image.ZP)

	w.output.Image = append(w.output.Image, paletted)
	w.output.Delay = append(w.output.Delay, duration)
}

// Write writes images to one animation gif file.
func (w Writer) Write(f io.Writer) error {
	return gif.EncodeAll(f, w.output)
}
