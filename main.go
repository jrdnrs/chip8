package main

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const debug bool = false

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
	game := &Game{}
	game.chip8 = &Chip8{}
	game.chip8.Initialise()
	cpuTicker := time.NewTicker(time.Second / 500)

	// Audio
	a, err := NewBeeper(44100, 600, 0.005)
	if err != nil {
		panic(err)
	}
	game.audio = a
	game.audio.InitBeeper()
	soundTicker := time.NewTicker(time.Second / 60)

	

	// Display
	game.display = NewDisplay(&game.chip8.gfx, &game.chip8.drawRequired)
	game.display.InitDisplay()

	// Load rom
	game.chip8.LoadRom(rom)

	// Run timers
	go func() {
		for range cpuTicker.C {
			game.updateKeys()
			game.chip8.EmulateCycle()
		}
	}()

	go func() {
		for range soundTicker.C {
			game.chip8.updateTimers()
			game.audio.soundChannel <- game.chip8.soundTimer
		}
	}()

	// Run
	ebiten.Run(game.display.Render, 64, 32, 10, "CHIP-8")

}

func openRom(path string) ([]byte, error) {
	rom, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return rom, nil

}

// Game implements ebiten.Game interface.
type Game struct {
	chip8   *Chip8
	audio   *Beeper
	display *Display
}

func (g *Game) updateKeys() {
	keyMap := map[ebiten.Key]int{
		ebiten.KeyX: 0x0,
		ebiten.Key1: 0x1,
		ebiten.Key2: 0x2,
		ebiten.Key3: 0x3,
		ebiten.KeyQ: 0x4,
		ebiten.KeyW: 0x5,
		ebiten.KeyE: 0x6,
		ebiten.KeyA: 0x7,
		ebiten.KeyS: 0x8,
		ebiten.KeyD: 0x9,
		ebiten.KeyZ: 0xA,
		ebiten.KeyC: 0xB,
		ebiten.Key4: 0xC,
		ebiten.KeyR: 0xD,
		ebiten.KeyF: 0xE,
		ebiten.KeyV: 0xF,
	}

	for k, v := range keyMap {
		if ebiten.IsKeyPressed(k) {
			g.chip8.key[v] = 1
		} else {
			g.chip8.key[v] = 0
		}
	}

}