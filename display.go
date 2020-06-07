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
	memory        *[32][8]byte
	previousFrame [32][8]byte
	display       *ebiten.Image
	drawRequired  *bool
}

// NewDisplay exposes Render function which uses gfxMemory to draw image to screen.
func NewDisplay(gfxMemory *[32][8]byte, drawRequired *bool) *Display {
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
	for y, xBytes := range d.memory { // loop through array of rows

		for x := range xBytes { // loop through bytes in row

			if d.memory[y][x] != d.previousFrame[y][x] { // don't check further if entire byte is same as previous frame

				var mask byte

				for b := 0; b < 8; b++ { // loop through bits with mask to update those that have changed

					mask = 0x80 >> b

					if d.memory[y][x]&mask != d.previousFrame[y][x]&mask {
						if d.memory[y][x]&mask != 0 {
							d.display.Set(x*8+b, y, on)
						} else {
							d.display.Set(x*8+b, y, off)
						}
					}
				}
			}
		}
	}
}
