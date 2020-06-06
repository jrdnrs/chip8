package main

import (
	"math"
)

// Generates tone for given duration is ms.
func genTone(duration int, sampleRate int, frequency int, channels int) []byte {
	numSamples := duration * sampleRate / 1000
	samples := make([]float64, numSamples)
	finalSound := make([]byte, 2*numSamples*channels)

	// Fill array with samples https://www.desmos.com/calculator/6qpozvgl9m.
	for i := range samples {
		samples[i] = math.Sin(float64(2*i) * math.Pi / float64(sampleRate/frequency))
	}


	var val int16
	for i, dval := range samples {
		// Scale to maximum amplitude.
		val = int16(dval * 0x7FFF)

		// Per channel, 2 bytes in little endian
		for j := 0; j < channels; j++ {
			finalSound[i*2*channels+(j*2)] = byte(val & 0x00FF)
			finalSound[i*2*channels+(j*2)+1] = byte((uint16(val) & 0xFF00) >> 8)
		}
	}

	return finalSound

}
