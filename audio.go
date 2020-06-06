package main

import (
	"math"

	"github.com/hajimehoshi/oto"
)

// Beeper handles audio, utilising the soundTimer register.
type Beeper struct {
	player            *oto.Player
	sampleRate        int
	byteDepth         int
	bufferSizeInBytes int
	toneFrequency     float64
	volume            float64
	amplitude         int16
	step              float64
	counter           float64
	soundChannel      chan byte
}

// NewBeeper returns instance of Beeper, with sampling function. SoundTimer register values should be passed through the soundChannel at 60hz.
func NewBeeper(sampleRate int, toneFrequency float64, volume float64) (*Beeper, error) {
	c, err := oto.NewContext(sampleRate, 1, 2, int(sampleRate/15))

	if err != nil {
		return nil, err
	}

	return &Beeper{
		player:            c.NewPlayer(),
		sampleRate:        sampleRate,
		bufferSizeInBytes: int(sampleRate / 15),
		toneFrequency:     toneFrequency,
		volume:            volume,
		amplitude:         int16(volume * 0x7FFF),
		step:              toneFrequency * 2 * math.Pi / float64(sampleRate),
		soundChannel:      make(chan byte, 60),
	}, nil
}

// InitBeeper starts a goroutine which listens to the soundChannel and generates square wave samples when passed value is positive.
func (b *Beeper) InitBeeper() {
	go func() {
		for i := range b.soundChannel {
			if i > 0 {
				if _, err := b.player.Write(b.generateSample()); err != nil {
					panic(err)
				}
			}
		}
	}()
}

// Generate number of 16bit samples equal to half size of sound buffer and store them in little endian. 
func (b *Beeper) generateSample() []byte {
	numSamples := b.bufferSizeInBytes / 2	
	samples := make([]int16, numSamples)
	b.generateSquareWave(samples)
	bytes := make([]byte, b.bufferSizeInBytes)
	for i, s := range samples {
		bytes[2*i] = byte(s)
		bytes[2*i+1] = byte(s >> 8)
	}
	return bytes
}

// Writes positive or negative amplitude according to sin of counter. Counter increments using specified toneFrequency and sampleRate.
func (b *Beeper) generateSquareWave(sample []int16) {
	for i := range sample {
		if math.Sin(b.counter) < 0 {
			sample[i] = -b.amplitude
		} else {
			sample[i] = b.amplitude
		}
		b.counter += b.step
	}
}
