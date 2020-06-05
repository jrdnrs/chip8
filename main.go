package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
)

const debug bool = false

func main() {

	if debug {
		runTests()
		return
	}

	romPath := os.Args[1]

	game := &Game{}
	game.buffer = make([]byte, 64*32*4)
	game.chip8 = &Chip8{}
	game.chip8.Initialise()

	rom, err := openRom(romPath)

	if err != nil {
		return
	}

	game.chip8.LoadRom(rom)

	// Sepcify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 320)
	ebiten.SetWindowTitle("CHIP-8")
	ebiten.SetMaxTPS(700)

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

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
	chip8  *Chip8
	buffer []byte
}

func (g *Game) updateKeys() {
	keyMap := map[ebiten.Key]int {
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

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {
	// Write your game's logical update.

	g.updateKeys()
	g.chip8.EmulateCycle()

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.

	if g.chip8.drawFlag {

		for i := range g.buffer {
			g.buffer[i] = g.chip8.gfx[int(i/4)] * 255
		}

		screen.ReplacePixels(g.buffer)

	}

}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 64, 32
}
