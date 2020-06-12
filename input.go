package main

import (
	"github.com/hajimehoshi/ebiten"
)

// Input handles key presses. Use NewInput to initialise.
type Input struct {
	memory       *[16]byte
	gameKeys     map[ebiten.Key]byte
	functionKeys map[ebiten.Key]func()

	reset func()
	pause func()
	cont  func()
	step  func()
}

// NewInput returns a pointer to Input which handles key presses. This includes game keys and function keys.
// The keyMemory pointer represents the byte array from which the chip8 emulator will read from to determine which keys are currently pressed.
// The reset, pause, continue and step func args will be called with F1, F2, F3 and F4 respectively.
func NewInput(keyMemory *[16]byte, reset func(), pause func(), cont func(), step func()) *Input {
	i := &Input{
		memory: keyMemory,
		reset:  reset,
		pause:  pause,
		cont:   cont,
		step:   step,
	}

	i.gameKeys = map[ebiten.Key]byte{
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
	i.functionKeys = map[ebiten.Key]func(){
		ebiten.KeyF1: i.reset,
		ebiten.KeyF2: i.pause,
		ebiten.KeyF3: i.cont,
		ebiten.KeyF4: i.step,
	}

	return i
}

// UpdateInput checks which keys are currently pressed and either updates the chip8 emulator's key memory or calls the relevant func.
func (i *Input) UpdateInput() {
	for k, v := range i.gameKeys {
		if ebiten.IsKeyPressed(k) {
			i.memory[v] = 1
		} else {
			i.memory[v] = 0
		}
	}

	for k, v := range i.functionKeys {
		if ebiten.IsKeyPressed(k) {
			v()
			break
		}
	}
}
