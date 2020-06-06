package main

import (
	"github.com/hajimehoshi/ebiten"
)

type Input struct {
	memory *[16]byte
	keyMap map[ebiten.Key]byte
}

func NewInput(keyMemory *[16]byte) *Input {
	return &Input{
		memory: keyMemory,
	}
}

func (i *Input) InitInput() {
	i.keyMap = map[ebiten.Key]byte{
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
}

func (i *Input) UpdateInput() {
	for k, v := range i.keyMap {
		if ebiten.IsKeyPressed(k) {
			i.memory[v] = 1
		} else {
			i.memory[v] = 0
		}
	}
}
