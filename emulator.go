package main

import (
	"fmt"
	"math/rand"
)

var fontset = [80]byte{
	0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
	0x20, 0x60, 0x20, 0x20, 0x70, // 1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
	0x90, 0x90, 0xF0, 0x10, 0x10, // 4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
	0xF0, 0x10, 0x20, 0x40, 0x40, // 7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
	0xF0, 0x90, 0xF0, 0x90, 0x90, // A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
	0xF0, 0x80, 0x80, 0x80, 0xF0, // C
	0xE0, 0x90, 0x90, 0x90, 0xE0, // D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
	0xF0, 0x80, 0xF0, 0x80, 0x80, // F
}

// Emulator opcode interpreter.
type Emulator struct {
	// Current opcode to be executed.
	opcode uint16

	// 4096 8-bit words for main memory.
	// 0x000-0x1FF is used for Chip8 interpreter, however this will contain the font set in emu.
	// 0x050-0x0A0 is used for the built in 4x5 pixel font set.
	// 0x200-0xFFF is used for program rom and rest is work ram.
	memory [4096]byte
	// 15 8-bit general purpose registers names V0, V1...VE. The 16th register (VF) is used for the carry flag.
	register [16]byte

	// 16-bit Index register. Can have value from 0x000-0xFFF.
	i uint16
	// 16-bit Program Counter register. Can have value from 0x000-0xFFF.
	pc uint16

	// Display of 2048 pixels (2:1).
	gfx [64 * 32]byte

	// The timers will count down at 60hz, when set above zero.
	delayTimer byte
	soundTimer byte

	// The stack is used to remember the current location before a jump is performed.
	stack [16]uint16
	sp    uint16

	// HEX based keypad (0x0-0xF).
	key [16]byte

	drawRequired bool
}

// Initialise registers and memory once.
func (emu *Emulator) Initialise() {
	emu.pc = 0x200 // System expects the application to be loaded at memory location 0x200.
	emu.opcode = 0 // Reset current opcode.
	emu.i = 0      // Reset Index register.
	emu.sp = 0     // Reset Stack Pointer.

	// Clear registers V0-VF.
	for i := range emu.register {
		emu.register[i] = 0

	}

	// Load fontset and clear other memory.
	for i := range emu.memory {
		if i < 80 {
			emu.memory[i] = fontset[i]
		} else {
			emu.memory[i] = 0

		}

	}

	// Clear display.
	for i := range emu.gfx {
		emu.gfx[i] = 0

	}

	// clear stack.
	for i := range emu.stack {
		emu.stack[i] = 0

	}

}

// LoadRom accepts slice of bytes and loads into memory, starting at 0x200.
func (emu *Emulator) LoadRom(rom []byte) error {
	for i, byt := range rom {
		emu.memory[0x200+i] = byt
	}

	return nil

}

// EmulateCycle fetches, decodes, executes next opcode and updates timers.
func (emu *Emulator) EmulateCycle() {

	// Opcodes are two bytes long and stored big-endian.
	emu.opcode = uint16(emu.memory[emu.pc])<<8 | uint16(emu.memory[emu.pc+1])

	unknownOpcode := func() {
		fmt.Printf("Unknown Opcode: 0x%X\n", emu.opcode)
	}

	switch emu.opcode & 0xF000 {
	case 0x0000:
		switch emu.opcode & 0x000F {
		case 0x0000:
			emu.x00E0()
		case 0x000E:
			emu.x00EE()
		default:
			unknownOpcode()
		}
	case 0x1000:
		emu.x1NNN()
	case 0x2000:
		emu.x2NNN()
	case 0x3000:
		emu.x3XNN()
	case 0x4000:
		emu.x4XNN()
	case 0x5000:
		emu.x5XY0()
	case 0x6000:
		emu.x6XNN()
	case 0x7000:
		emu.x7XNN()
	case 0x8000:
		switch emu.opcode & 0x000F {
		case 0x0000:
			emu.x8XY0()
		case 0x0001:
			emu.x8XY1()
		case 0x0002:
			emu.x8XY2()
		case 0x0003:
			emu.x8XY3()
		case 0x0004:
			emu.x8XY4()
		case 0x0005:
			emu.x8XY5()
		case 0x0006:
			emu.x8XY6()
		case 0x0007:
			emu.x8XY7()
		case 0x000E:
			emu.x8XYE()
		default:
			unknownOpcode()
		}
	case 0x9000:
		emu.x9XY0()
	case 0xA000:
		emu.xANNN()
	case 0xB000:
		emu.xBNNN()
	case 0xC000:
		emu.xCXNN()
	case 0xD000:
		emu.xDXYN()
	case 0xE000:
		switch emu.opcode & 0x000F {
		case 0x000E:
			emu.xEX9E()
		case 0x0001:
			emu.xEXA1()
		default:
			unknownOpcode()
		}
	case 0xF000:
		switch emu.opcode & 0x00FF {
		case 0x0007:
			emu.xFX07()
		case 0x000A:
			emu.xFX0A()
		case 0x0015:
			emu.xFX15()
		case 0x0018:
			emu.xFX18()
		case 0x001E:
			emu.xFX1E()
		case 0x0029:
			emu.xFX29()
		case 0x0033:
			emu.xFX33()
		case 0x0055:
			emu.xFX55()
		case 0x0065:
			emu.xFX65()
		default:
			unknownOpcode()
		}
	default:
		unknownOpcode()
	}

}

// Increment Program Counter by count * 2, as each instruction takes up two registers in memory.
func (emu *Emulator) incrementPC(count uint16) {
	emu.pc += 2 * count

}

// CHIP-8 has two timers. They both count down at 60 hertz, until they reach 0.
func (emu *Emulator) updateTimers() {
	if emu.delayTimer > 0 {
		emu.delayTimer--

	}

	if emu.soundTimer > 0 {
		emu.soundTimer--

	}

}

/////////////////////////////////
// Opcodes
/////////////////////////////////

// Clears the screen.
func (emu *Emulator) x00E0() {
	for i := range emu.gfx {
		emu.gfx[i] = 0

	}

	emu.incrementPC(1)

	emu.drawRequired = true
}

// Returns from a subroutine.
func (emu *Emulator) x00EE() {
	emu.sp--
	emu.pc = emu.stack[emu.sp]

	emu.incrementPC(1)

}

// Jumps to address NNN.
func (emu *Emulator) x1NNN() {
	nnn := emu.opcode & 0x0FFF
	emu.pc = nnn

}

// Calls subroutine at NNN.
func (emu *Emulator) x2NNN() {
	emu.stack[emu.sp] = emu.pc
	emu.sp++

	nnn := emu.opcode & 0x0FFF
	emu.pc = nnn

}

// Skips the next instruction if VX equals NN.
func (emu *Emulator) x3XNN() {
	x := int((emu.opcode & 0x0F00) >> 8)
	nn := byte(emu.opcode & 0x00FF)

	if emu.register[x] == nn {
		emu.incrementPC(2)

	} else {
		emu.incrementPC(1)

	}

}

// Skips the next instruction if VX doesn't equal NN.
func (emu *Emulator) x4XNN() {
	x := int((emu.opcode & 0x0F00) >> 8)
	nn := byte(emu.opcode & 0x00FF)

	if emu.register[x] != nn {
		emu.incrementPC(2)

	} else {
		emu.incrementPC(1)

	}

}

// Skips the next instruction if VX equals VY.
func (emu *Emulator) x5XY0() {
	x := int((emu.opcode & 0x0F00) >> 8)
	y := int((emu.opcode & 0x00F0) >> 4)

	if emu.register[x] == emu.register[y] {
		emu.incrementPC(2)

	} else {
		emu.incrementPC(1)

	}

}

// Sets VX to NN.
func (emu *Emulator) x6XNN() {
	x := int((emu.opcode & 0x0F00) >> 8)
	nn := byte(emu.opcode & 0x00FF)

	emu.register[x] = nn

	emu.incrementPC(1)

}

// Adds NN to VX. (Carry flag is not changed).
func (emu *Emulator) x7XNN() {
	x := int((emu.opcode & 0x0F00) >> 8)
	nn := byte(emu.opcode & 0x00FF)

	emu.register[x] += nn

	emu.incrementPC(1)

}

// Sets VX to the value of VY.
func (emu *Emulator) x8XY0() {
	x := int((emu.opcode & 0x0F00) >> 8)
	y := int((emu.opcode & 0x00F0) >> 4)

	emu.register[x] = emu.register[y]

	emu.incrementPC(1)

}

// Sets VX to VX or VY. (Bitwise OR operation).
func (emu *Emulator) x8XY1() {
	x := int((emu.opcode & 0x0F00) >> 8)
	y := int((emu.opcode & 0x00F0) >> 4)

	emu.register[x] = emu.register[x] | emu.register[y]

	emu.incrementPC(1)

}

// Sets VX to VX and VY. (Bitwise AND operation).
func (emu *Emulator) x8XY2() {
	x := int((emu.opcode & 0x0F00) >> 8)
	y := int((emu.opcode & 0x00F0) >> 4)

	emu.register[x] = emu.register[x] & emu.register[y]

	emu.incrementPC(1)
}

// Sets VX to VX xor VY.
func (emu *Emulator) x8XY3() {
	x := int((emu.opcode & 0x0F00) >> 8)
	y := int((emu.opcode & 0x00F0) >> 4)

	emu.register[x] = emu.register[x] ^ emu.register[y]

	emu.incrementPC(1)
}

// Adds VY to VX. VF is set to 1 when there's a carry, and to 0 when there isn't.
func (emu *Emulator) x8XY4() {
	x := int((emu.opcode & 0x0F00) >> 8)
	y := int((emu.opcode & 0x00F0) >> 4)

	if int(emu.register[x])+int(emu.register[y]) > 0xFF {
		emu.register[0xF] = 1

	} else {
		emu.register[0xF] = 0

	}

	emu.register[x] += emu.register[y]

	emu.incrementPC(1)

}

// VY is subtracted from VX. VF is set to 0 when there's a borrow, and 1 when there isn't.
func (emu *Emulator) x8XY5() {
	x := int((emu.opcode & 0x0F00) >> 8)
	y := int((emu.opcode & 0x00F0) >> 4)

	if emu.register[y] > emu.register[x] {
		emu.register[0xF] = 0

	} else {
		emu.register[0xF] = 1

	}

	emu.register[x] -= emu.register[y]

	emu.incrementPC(1)

}

// Stores the least significant bit of VX in VF and then shifts VX to the right by 1.
func (emu *Emulator) x8XY6() {
	x := int((emu.opcode & 0x0F00) >> 8)

	emu.register[0xF] = emu.register[x] & 0x01
	emu.register[x] = emu.register[x] >> 1

	emu.incrementPC(1)

}

// Sets VX to VY minus VX. VF is set to 0 when there's a borrow, and 1 when there isn't.
func (emu *Emulator) x8XY7() {
	x := int((emu.opcode & 0x0F00) >> 8)
	y := int((emu.opcode & 0x00F0) >> 4)

	if emu.register[x] > emu.register[y] {
		emu.register[0xF] = 0

	} else {
		emu.register[0xF] = 1

	}

	emu.register[x] = emu.register[y] - emu.register[x]

	emu.incrementPC(1)

}

// Stores the most significant bit of VX in VF and then shifts VX to the left by 1.
func (emu *Emulator) x8XYE() {
	x := int((emu.opcode & 0x0F00) >> 8)

	emu.register[0xF] = (emu.register[x] & 0x80) >> 7
	emu.register[x] = emu.register[x] << 1

	emu.incrementPC(1)

}

// Skips the next instruction if VX doesn't equal VY.
func (emu *Emulator) x9XY0() {
	x := int((emu.opcode & 0x0F00) >> 8)
	y := int((emu.opcode & 0x00F0) >> 4)

	if emu.register[x] != emu.register[y] {
		emu.incrementPC(2)

	} else {
		emu.incrementPC(1)

	}

}

// Sets I to the address NNN.
func (emu *Emulator) xANNN() {
	nnn := emu.opcode & 0x0FFF
	emu.i = nnn

	emu.incrementPC(1)

}

// Jumps to the address NNN plus V0.
func (emu *Emulator) xBNNN() {
	nnn := emu.opcode & 0x0FFF
	emu.pc = nnn + uint16(emu.register[0])

}

// Sets VX to the result of a bitwise and operation on a random number (Typically: 0 to 0xFF) and NN.
func (emu *Emulator) xCXNN() {
	x := int((emu.opcode & 0x0F00) >> 8)
	nn := emu.opcode & 0x00FF

	emu.register[x] = byte(nn) & byte(rand.Intn(0xFF))

	emu.incrementPC(1)

}

// Draws a sprite at coordinate (VX, VY) that has a width of 8 pixels and a height of N pixels.
// Each row of 8 pixels is read as bit-coded starting from memory location I; I value doesn’t change after the execution of this instruction.
// As described above, VF is set to 1 if any screen pixels are flipped from set to unset when the sprite is drawn, and to 0 if that doesn’t happen.
func (emu *Emulator) xDXYN() {
	x := int((emu.opcode & 0x0F00) >> 8)
	y := int((emu.opcode & 0x00F0) >> 4)
	n := int((emu.opcode & 0x000F))

	vx := int(emu.register[x]) // display x coordinate
	vy := int(emu.register[y]) // display y coordinate

	start := 64*vy + vx // number of pixels since origin
	pos := start             // current pixel

	for yLine := 0; yLine < n; yLine++ {
		spr := emu.memory[int(emu.i)+yLine] // bit-coded sprite data

		for xLine := 0; xLine < 8; xLine++ {
			pos = start + (yLine * 64) + xLine

			pix := spr & (0x80 >> xLine) // scan through sprite, one pixel at a time

			if pix >= 1 {
				if emu.gfx[pos] == 1 {
					emu.gfx[pos] = 0
					emu.register[0xF] = 1

				} else {
					emu.gfx[pos] = 1
					emu.register[0xF] = 0

				}
			}
		}
	}
	emu.drawRequired = true
	emu.incrementPC(1)

}

// Skips the next instruction if the key stored in VX is pressed.
func (emu *Emulator) xEX9E() {
	x := int((emu.opcode & 0x0F00) >> 8)

	if emu.key[emu.register[x]] != 0 {
		emu.incrementPC(2)

	} else {
		emu.incrementPC(1)

	}

}

// Skips the next instruction if the key stored in VX isn't pressed.
func (emu *Emulator) xEXA1() {
	x := int((emu.opcode & 0x0F00) >> 8)

	if emu.key[emu.register[x]] == 0 {
		emu.incrementPC(2)

	} else {
		emu.incrementPC(1)

	}

}

// Sets VX to the value of the delay timer.
func (emu *Emulator) xFX07() {
	x := int((emu.opcode & 0x0F00) >> 8)

	emu.register[x] = emu.delayTimer

	emu.incrementPC(1)

}

// A key press is awaited, and then stored in VX. (Blocking Operation. All instruction halted until next key event).
func (emu *Emulator) xFX0A() {
	x := int((emu.opcode & 0x0F00) >> 8)

	pressed := false

	for i := range emu.key {
		if emu.key[i] != 0 {
			emu.register[x] = byte(i)
			pressed = true

		}

	}

	if pressed {
		emu.incrementPC(1)

	}

}

// Sets the delay timer to VX.
func (emu *Emulator) xFX15() {
	x := int((emu.opcode & 0x0F00) >> 8)

	emu.delayTimer = emu.register[x]

	emu.incrementPC(1)

}

// Sets the sound timer to VX.
func (emu *Emulator) xFX18() {
	x := int((emu.opcode & 0x0F00) >> 8)

	emu.soundTimer = emu.register[x]

	emu.incrementPC(1)
}

// Adds VX to I. VF is set to 1 when there is a range overflow (I+VX>0xFFF), and to 0 when there isn't.
func (emu *Emulator) xFX1E() {
	x := int((emu.opcode & 0x0F00) >> 8)

	emu.i += uint16(emu.register[x])

	if emu.i > 0xFFF {
		emu.register[0xF] = 1

	} else {
		emu.register[0xF] = 0

	}

	emu.incrementPC(1)

}

// Sets I to the location of the sprite for the character in VX. Characters 0-F (in hexadecimal) are represented by a 4x5 font.
func (emu *Emulator) xFX29() {
	x := int((emu.opcode & 0x0F00) >> 8)

	emu.i = uint16(emu.register[x]) * 5

	emu.incrementPC(1)

}

// Stores the binary-coded decimal representation of VX, with the most significant of three digits at the address in I, the middle digit at I plus 1, and the least significant digit at I plus 2.
// (In other words, take the decimal representation of VX, place the hundreds digit in memory at location in I, the tens digit at location I+1, and the ones digit at location I+2.)
func (emu *Emulator) xFX33() {
	x := int((emu.opcode & 0x0F00) >> 8)
	vx := emu.register[x]

	emu.memory[emu.i] = vx / 100
	emu.memory[emu.i+1] = (vx / 10) % 10
	emu.memory[emu.i+2] = (vx % 100) % 10

	emu.incrementPC(1)

}

// Stores V0 to VX (including VX) in memory starting at address I. The offset from I is increased by 1 for each value written, but I itself is left unmodified.
func (emu *Emulator) xFX55() {
	x := int((emu.opcode & 0x0F00) >> 8)

	for i := 0; i <= x; i++ {
		emu.memory[int(emu.i)+i] = emu.register[i]

	}

	emu.incrementPC(1)

}

// Fills V0 to VX (including VX) with values from memory starting at address I. The offset from I is increased by 1 for each value written, but I itself is left unmodified.
func (emu *Emulator) xFX65() {
	x := int((emu.opcode & 0x0F00) >> 8)

	for i := 0; i <= x; i++ {
		emu.register[i] = emu.memory[int(emu.i)+i]

	}

	emu.incrementPC(1)

}
