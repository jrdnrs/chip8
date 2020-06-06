package main

import (
	"math"

	"github.com/hajimehoshi/oto"
)

// Audio handles audio.
type Audio struct {
	player            *oto.Player
	sampleRate        int
	byteDepth         int
	bufferSizeInBytes int
	toneFrequency     float64
	sequence          float64
	volume            float64
	soundChannel      chan byte
}

// NewAudio returns instance of Audio, with sampling function.
func NewAudio(sampleRate int, byteDepth int, bufferSizeInBytes int, toneFrequency float64, sequence float64, volume float64) (*Audio, error) {
	c, err := oto.NewContext(sampleRate, 1, byteDepth, bufferSizeInBytes)

	if err != nil {
		return nil, err
	}

	return &Audio{
		player:            c.NewPlayer(),
		sampleRate:        sampleRate,
		byteDepth:         byteDepth,
		bufferSizeInBytes: bufferSizeInBytes,
		toneFrequency:     toneFrequency,
		sequence:          sequence,
		volume:            volume,
		soundChannel:      make(chan byte, 60),
	}, nil
}

// InitSound starts a goroutine which waits for sound to be played.
func (a *Audio) InitSound() {
	go func() {
		for b := range a.soundChannel {
			if b > 0x0 {
				if _, err := a.player.Write(a.generateSample()); err != nil {
					panic(err)
				}
			}
		}
	}()
}

func (a *Audio) generateSample() []byte {
	length := a.bufferSizeInBytes / 2
	s := make([]int16, length)
	a.fill(s)
	bytes := make([]byte, a.bufferSizeInBytes)
	for i := range s {
		bytes[2*i] = byte(s[i])
		bytes[2*i+1] = byte(s[i] >> 8)
	}
	return bytes
}

func (a *Audio) fill(sample []int16) {
	vol := float64(a.volume)
	f := float64(a.sampleRate)
	length := int(f / float64(a.toneFrequency))
	var amp int16
	for i := 0; i < len(sample); i++ {
		amp = int16(vol * math.MaxInt16)
		if i%length < int(float64(length)*a.sequence) {
			amp = -amp
		}
		sample[i] = amp
	}
}
