package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

var (
	off = color.RGBA{0x90, 0x90, 0x90, 0xFF}
	on  = color.RGBA{0x1E, 0x1E, 0x1E, 0xFF}
)

// Display handles rendering functions.
type Display struct {
	memory        *[2048]byte
	previousFrame [2048]byte
	display       *ebiten.Image
	drawRequired  *bool
}

// NewDisplay exposes Render function which uses gfxMemory to draw image to screen.
func NewDisplay(gfxMemory *[2048]byte, drawRequired *bool) *Display {

	return &Display{
		memory:        gfxMemory,
		previousFrame: *gfxMemory,
		drawRequired:  drawRequired,
	}
}

// InitDisplay creates blank image for drawing to.
func (d *Display) InitDisplay() {
	image, _ := ebiten.NewImage(64, 32, ebiten.FilterDefault)
	image.Fill(off)
	d.display = image
}

// Render draws the pixels
func (d *Display) Render(screen *ebiten.Image) error {
	if *(d.drawRequired) {
		d.drawPixels()
		d.previousFrame = *d.memory
		*(d.drawRequired) = false
	}

	if err := screen.DrawImage(d.display, &ebiten.DrawImageOptions{}); err != nil {
		return err
	}

	return nil
}

func (d *Display) drawPixels() {
	for i, b := range d.memory {
		if b != d.previousFrame[i] {
			if b != 0x0 {
				d.display.Set(i%64, int(i/64), on)
			} else {
				d.display.Set(i%64, int(i/64), off)
			}
		}
	}
}
