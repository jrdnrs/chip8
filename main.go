package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	clockSpeed := flag.Int64("clockspeed", 700, "The number of cycles executed per second.")
	displayScale := flag.Float64("displayscale", 8, "Multiplier for screen size. '1' is 64x32.")
	audioSampleRate := flag.Int("audiosamplerate", 44100, "Sample rate for the audio.")
	audioFrequency := flag.Float64("audiofrequency", 200, "Frequency of the audio tone.")
	audioVolume := flag.Float64("audiovolume", 0.5, "Multiplier for audio volume, between 0 and 1.")
	flag.Parse()
	romPath := flag.Arg(0)

	if *clockSpeed < 0 {
		fmt.Println("Clock speed of 0 or greater is required.")
		os.Exit(1)
	}
	if *displayScale < 1 {
		fmt.Println("Display scale of 1 or greater is required.")
		os.Exit(1)
	}
	if *audioSampleRate < 0 {
		fmt.Println("Audio sample rate of 0 or greater is required.")
		os.Exit(1)
	}
	if *audioVolume < 0 || *audioVolume > 1 {
		fmt.Println("Audio volume between 0 and 1 is required.")
		os.Exit(1)
	}

	chip8 := NewChip8(*clockSpeed, *displayScale, *audioSampleRate, *audioFrequency, *audioVolume, romPath)
	chip8.Run()
}

// Chip8 contains implementation of chip8 emulator as well as facilities to play sound, render to screen and read input.
type Chip8 struct {
	emu     *Emulator
	audio   *Beeper
	display *Display
	input   *Input
}

// Run starts the emulation.
func (c8 *Chip8) Run() {
	ebiten.SetRunnableInBackground(true)
	ebiten.SetMaxTPS(60)
	ebiten.Run(c8.loop, 64, 32, c8.display.DisplayScale, "CHIP-8")
}

// NewChip8 provides a pointer to an initialised Chip8, using provided args.
func NewChip8(clockSpeed int64, displayScale float64, audioSampleRate int, audioFrequency float64, audioVolume float64, romPath string) *Chip8 {
	rom, err := ioutil.ReadFile(romPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c8 := &Chip8{}
	c8.emu = NewEmulator(clockSpeed, rom)
	c8.audio = NewBeeper(&c8.emu.SoundTimer, audioSampleRate, audioFrequency, audioVolume)
	c8.display = NewDisplay(&c8.emu.Display, displayScale)
	c8.input = NewInput(&c8.emu.Key, c8.emu.Reset, c8.emu.Pause, c8.emu.Continue, c8.emu.EmulateCycle)

	return c8
}

// Main loop for ebiten to run every tick.
func (c8 *Chip8) loop(screen *ebiten.Image) error {
	c8.input.UpdateInput()
	c8.emu.Process()
	if c8.audio.IsInitialised {
		c8.audio.UpdateSound()
	}
	c8.emu.UpdateTimers()
	c8.display.Render(screen)

	return nil
}
