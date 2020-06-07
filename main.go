package main

import (
	"os"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	debug = false

	clockSpeed = 500 // hz

	beeperToneFrequency = 555   // hz
	beeperSampleRate    = 44100 // hz
	beeperVolume        = 0.05  // 0-1

	displayScale = 10
)

func main() {

	if debug {
		runTests()
		return
	}

	// Get rom path from arg and open file
	romPath := os.Args[1]
	rom, err := openRom(romPath)
	if err != nil {
		panic(err)
	}

	// Prepare emu
	chip8 := &Chip8{}
	chip8.emu = &Emulator{}
	chip8.emu.Initialise()

	// Audio
	b, err := NewBeeper(beeperSampleRate, beeperToneFrequency, beeperVolume)
	if err != nil {
		panic(err)
	}
	chip8.audio = b
	chip8.audio.InitBeeper()

	// Display
	chip8.display = NewDisplay(&chip8.emu.gfx, &chip8.emu.drawRequired)
	chip8.display.InitDisplay()

	//Input
	chip8.input = NewInput(&chip8.emu.key)
	chip8.input.InitInput()

	// Load rom
	chip8.emu.LoadRom(rom)

	soundTicker := time.NewTicker(time.Second / 60)
	cpuTicker := time.NewTicker(time.Second / clockSpeed)

	// Run timers
	go func() {
		for range cpuTicker.C {
			chip8.input.UpdateInput()
			chip8.emu.EmulateCycle()
		}
	}()

	go func() {
		for range soundTicker.C {
			chip8.emu.updateTimers()
			chip8.audio.soundChannel <- chip8.emu.soundTimer
		}
	}()

	// Run
	ebiten.Run(chip8.display.Render, 64, 32, displayScale, "CHIP-8")

}

// Chip8 implements ebiten.Chip8 interface.
type Chip8 struct {
	emu     *Emulator
	audio   *Beeper
	display *Display
	input   *Input
}
