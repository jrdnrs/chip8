package main

import (
	// "log"
	// "io/ioutil"
	// "os"

	// "github.com/hajimehoshi/ebiten"

)

func main() {

	tester()


	// romPath := os.Args[1]

	// game := &Game{}
	// game.buffer = make([]byte, 64*32*4)
	// game.chip8 = &emulator.Chip8{}
	// game.chip8.Initialise()

	// rom, err := openRom(romPath)

	// if err != nil {
	// 	return
	// }

	// game.chip8.LoadRom(rom)

	// // Sepcify the window size as you like. Here, a doubled size is specified.
	// ebiten.SetWindowSize(640, 320)
	// ebiten.SetWindowTitle("CHIP-8")

	// // Call ebiten.RunGame to start your game loop.
	// if err := ebiten.RunGame(game); err != nil {
	// 	log.Fatal(err)
	// }

}

// func openRom(path string) ([]byte, error) {
// 	rom, err := ioutil.ReadFile(path)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return rom, nil

// }

// // Game implements ebiten.Game interface.
// type Game struct {
// 	chip8  *emulator.Chip8
// 	buffer []byte
// }

// // Update proceeds the game state.
// // Update is called every tick (1/60 [s] by default).
// func (g *Game) Update(screen *ebiten.Image) error {
// 	// Write your game's logical update.

// 	g.chip8.EmulateCycle()

// 	return nil
// }

// // Draw draws the game screen.
// // Draw is called every frame (typically 1/60[s] for 60Hz display).
// func (g *Game) Draw(screen *ebiten.Image) {
// 	// Write your game's rendering.

// 	if g.chip8.drawFlag {

// 		for i := 0; i < 64*32; i++ {
// 			g.buffer[i] = g.chip8.gfx[i] * 255
// 			g.buffer[i+1] = g.chip8.gfx[i] * 255
// 			g.buffer[i+2] = g.chip8.gfx[i] * 255
// 			g.buffer[i+3] = byte(255)
// 		}

// 		screen.ReplacePixels(g.buffer)

// 	}

// }

// // Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// // If you don't have to adjust the screen size with the outside size, just return a fixed size.
// func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
// 	return 64, 32
// }
