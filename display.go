package main

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/hajimehoshi/ebiten"
)

var (
	off = color.RGBA{0xC5, 0xCA, 0xA4, 0xFF}
	on  = color.RGBA{0x48, 0x52, 0x39, 0xFF}
)

// Display handles rendering to screen. Use NewDisplay to initialise.
type Display struct {
	memory *[256]byte
	buffer *image.RGBA

	// multiplier for display resolution
	DisplayScale float64
}

// NewDisplay returns a pointer to Display which handles rendering to screen.
// The displayMemory pointer represents the 64x32 resolution bit-coded array of pixels, which is updated by the chip8 emulator.
// The displayScale arg is used as a multiplier for the resolution.
func NewDisplay(displayMemory *[256]byte, displayScale float64) *Display {
	d := &Display{
		memory:       displayMemory,
		DisplayScale: displayScale,
	}

	i := image.NewRGBA(image.Rect(0, 0, 64, 32))
	d.buffer = i
	d.fillBuffer(off)

	return d
}

// Render reads from the chip8 emulator's display memory and draws the final image to screen.
func (d *Display) Render(screen *ebiten.Image) {
	if ebiten.IsDrawingSkipped() {
		return
	}
	d.updateBuffer()
	screen.ReplacePixels(d.buffer.Pix)
}

// updateBuffer updates the offscreen image as a buffer before rendering to screen.
func (d *Display) updateBuffer() {
	d.fillBuffer(off)
	for pos, b := range d.memory {
		if b != 0 {
			for i := 0; i < 8; i++ {
				if b&(0x80>>i) != 0 {
					d.buffer.SetRGBA(((pos%8)*8)+i, pos/8, on)
				}
			}
		}
	}
}

// fillBuffer fills the image buffer with the given colour.
func (d *Display) fillBuffer(c color.RGBA) {
	draw.Draw(d.buffer, d.buffer.Bounds(), &image.Uniform{c}, image.Point{}, draw.Src)
}
