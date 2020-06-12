package main

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/oto"
)

// Beeper handles audio. Use NewBeeper to initialise.
type Beeper struct {
	memory     *byte
	player     *oto.Player
	sampleRate int
	frequency  float64
	volume     float64
	amplitude  float64
	step       float64
	time       float64

	// whether audio is able to be played
	IsInitialised bool
}

// NewBeeper returns a pointer to Beeper which handles audio.
// The soundTimer pointer represents the register to which the chip8 emulator will write to, indicating when sound is to be played.
// The sampleRate, frequency and volume args affect the audio accordingly.
func NewBeeper(soundTimer *byte, sampleRate int, frequency float64, volume float64) *Beeper {
	b := &Beeper{
		memory:     soundTimer,
		sampleRate: sampleRate,
		frequency:  frequency,
		volume:     volume,
	}

	c, err := oto.NewContext(b.sampleRate, 1, 2, int(b.sampleRate/60*6))
	if err != nil {
		fmt.Println(err)
		b.IsInitialised = false
		return b
	}

	b.player = c.NewPlayer()
	b.amplitude = b.volume * 0x7FFF
	b.step = b.frequency * 2 * math.Pi / float64(b.sampleRate)
	b.IsInitialised = true

	return b
}

// UpdateSound will read the chip8 emulator's soundTimer register. If it is greater than 0, samples will be generated and added to the queue to be played, else nothing will be played.
func (b *Beeper) UpdateSound() {
	if *b.memory > 0 {
		b.player.Write(b.generateSample())
	} else {
		b.player.Write(make([]byte, b.sampleRate/60*2))
	}
}

// generateSample creates enough 16bit single-channel samples for 60th of a second (the rate at which sound is played) and store them 8bit little endian.
func (b *Beeper) generateSample() []byte {
	n := b.sampleRate / 60
	bytes := make([]byte, n*2)

	for i := 0; i < n; i++ {
		s := int16(b.amplitude * curvyTriangle(b.time))

		bytes[2*i] = byte(s)
		bytes[2*i+1] = byte(s >> 8)

		b.time += b.step
	}

	return bytes
}

// wave funcs

func triangle(t float64) float64 {
	return (math.Abs(math.Mod(t, 2)-1) - 0.5) * 2
}

func curvyTriangle(t float64) float64 {
	return math.Pow(math.Abs(math.Mod(t, 2)-1), 3)
}

func square(t float64) float64 {
	if math.Mod(t, 2) < 1 {
		return 1
	}
	return -1
}

func sawtooth(t float64) float64 {
	return math.Mod(t, 2) - 1
}

func sine(t float64) float64 {
	return math.Sin(t)
}
