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

// Chip8 opcode interpreter.
type Chip8 struct {
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

	drawFlag bool
}

// Initialise registers and memory once.
func (c8 *Chip8) Initialise() {
	c8.pc = 0x200 // System expects the application to be loaded at memory location 0x200.
	c8.opcode = 0 // Reset current opcode.
	c8.i = 0      // Reset Index register.
	c8.sp = 0     // Reset Stack Pointer.

	// Clear registers V0-VF.
	for i := range c8.register {
		c8.register[i] = 0

	}

	// Load fontset and clear other memory.
	for i := range c8.memory {
		if i < 80 {
			c8.memory[i] = fontset[i]
		} else {
			c8.memory[i] = 0

		}

	}

	// Clear display.
	for i := range c8.gfx {
		c8.gfx[i] = 0

	}

	// clear stack.
	for i := range c8.stack {
		c8.stack[i] = 0

	}

}

// LoadRom accepts slice of bytes and loads into memory, starting at 0x200.
func (c8 *Chip8) LoadRom(rom []byte) error {
	for i, byt := range rom {
		c8.memory[0x200 + i] = byt
	}

	return nil
	
}

// EmulateCycle fetches, decodes, executes next opcode and updates timers.
func (c8 *Chip8) EmulateCycle() {

	// Opcodes are two bytes long and stored big-endian.
	c8.opcode = uint16(c8.memory[c8.pc])<<8 | uint16(c8.memory[c8.pc+1])

	unknownOpcode := func() {
		fmt.Printf("Unknown Opcode: 0x%X\n", c8.opcode)
	}

	switch c8.opcode & 0xF000 {
	case 0x0000:
		switch c8.opcode & 0x000F {
		case 0x0000:
			c8.x00E0()
		case 0x000E:
			c8.x00EE()
		default:
			unknownOpcode()
		}
	case 0x1000:
		c8.x1NNN()
	case 0x2000:
		c8.x2NNN()
	case 0x3000:
		c8.x3XNN()
	case 0x4000:
		c8.x4XNN()
	case 0x5000:
		c8.x5XY0()
	case 0x6000:
		c8.x6XNN()
	case 0x7000:
		c8.x7XNN()
	case 0x8000:
		switch c8.opcode & 0x000F {
		case 0x0000:
			c8.x8XY0()
		case 0x0001:
			c8.x8XY1()
		case 0x0002:
			c8.x8XY2()
		case 0x0003:
			c8.x8XY3()
		case 0x0004:
			c8.x8XY4()
		case 0x0005:
			c8.x8XY5()
		case 0x0006:
			c8.x8XY6()
		case 0x0007:
			c8.x8XY7()
		case 0x000E:
			c8.x8XYE()
		default:
			unknownOpcode()
		}
	case 0x9000:
		c8.x9XY0()
	case 0xA000:
		c8.xANNN()
	case 0xB000:
		c8.xBNNN()
	case 0xC000:
		c8.xCXNN()
	case 0xD000:
		c8.xDXYN()
	case 0xE000:
		switch c8.opcode & 0x000F {
		case 0x000E:
			c8.xEX9E()
		case 0x0001:
			c8.xEXA1()
		default:
			unknownOpcode()
		}
	case 0xF000:
		switch c8.opcode & 0x00FF {
		case 0x0007:
			c8.xFX07()
		case 0x000A:
			c8.xFX0A()
		case 0x0015:
			c8.xFX15()
		case 0x0018:
			c8.xFX18()
		case 0x001E:
			c8.xFX1E()
		case 0x0029:
			c8.xFX29()
		case 0x0033:
			c8.xFX33()
		case 0x0055:
			c8.xFX55()
		case 0x0065:
			c8.xFX65()
		default:
			unknownOpcode()
		}
	default:
		unknownOpcode()
	}

	c8.updateTimers()

}

// Increment Program Counter by count * 2, as each instruction takes up two registers in memory.
func (c8 *Chip8) incrementPC(count uint16) {
	c8.pc += 2 * count

}

// CHIP-8 has two timers. They both count down at 60 hertz, until they reach 0.
func (c8 *Chip8) updateTimers() {
	if c8.delayTimer > 0 {
		c8.delayTimer--

	}

	if c8.soundTimer > 0 {
		fmt.Println("BEEP")
		c8.soundTimer--

	}

}

/////////////////////////////////
// Opcodes
/////////////////////////////////

// Clears the screen.
func (c8 *Chip8) x00E0() {
	for i := range c8.gfx {
		c8.gfx[i] = 0

	}

	c8.incrementPC(1)

}

// Returns from a subroutine.
func (c8 *Chip8) x00EE() {
	c8.sp--
	c8.pc = c8.stack[c8.sp]

	c8.incrementPC(1)

}

// Jumps to address NNN.
func (c8 *Chip8) x1NNN() {
	nnn := c8.opcode & 0x0FFF
	c8.pc = nnn

}

// Calls subroutine at NNN.
func (c8 *Chip8) x2NNN() {
	c8.stack[c8.sp] = c8.pc
	c8.sp++

	nnn := c8.opcode & 0x0FFF
	c8.pc = nnn

}

// Skips the next instruction if VX equals NN.
func (c8 *Chip8) x3XNN() {
	x := int((c8.opcode & 0x0F00) >> 8)
	nn := byte(c8.opcode & 0x00FF)

	if c8.register[x] == nn {
		c8.incrementPC(2)

	} else {
		c8.incrementPC(1)

	}

}

// Skips the next instruction if VX doesn't equal NN.
func (c8 *Chip8) x4XNN() {
	x := int((c8.opcode & 0x0F00) >> 8)
	nn := byte(c8.opcode & 0x00FF)

	if c8.register[x] != nn {
		c8.incrementPC(2)

	} else {
		c8.incrementPC(1)

	}

}

// Skips the next instruction if VX equals VY.
func (c8 *Chip8) x5XY0() {
	x := int((c8.opcode & 0x0F00) >> 8)
	y := int((c8.opcode & 0x00F0) >> 4)

	if c8.register[x] == c8.register[y] {
		c8.incrementPC(2)

	} else {
		c8.incrementPC(1)

	}

}

// Sets VX to NN.
func (c8 *Chip8) x6XNN() {
	x := int((c8.opcode & 0x0F00) >> 8)
	nn := byte(c8.opcode & 0x00FF)

	c8.register[x] = nn

	c8.incrementPC(1)

}

// Adds NN to VX. (Carry flag is not changed).
func (c8 *Chip8) x7XNN() {
	x := int((c8.opcode & 0x0F00) >> 8)
	nn := byte(c8.opcode & 0x00FF)

	c8.register[x] += nn

	c8.incrementPC(1)

}

// Sets VX to the value of VY.
func (c8 *Chip8) x8XY0() {
	x := int((c8.opcode & 0x0F00) >> 8)
	y := int((c8.opcode & 0x00F0) >> 4)

	c8.register[x] = c8.register[y]

	c8.incrementPC(1)

}

// Sets VX to VX or VY. (Bitwise OR operation).
func (c8 *Chip8) x8XY1() {
	x := int((c8.opcode & 0x0F00) >> 8)
	y := int((c8.opcode & 0x00F0) >> 4)

	c8.register[x] = c8.register[x] | c8.register[y]

	c8.incrementPC(1)

}

// Sets VX to VX and VY. (Bitwise AND operation).
func (c8 *Chip8) x8XY2() {
	x := int((c8.opcode & 0x0F00) >> 8)
	y := int((c8.opcode & 0x00F0) >> 4)

	c8.register[x] = c8.register[x] & c8.register[y]

	c8.incrementPC(1)
}

// Sets VX to VX xor VY.
func (c8 *Chip8) x8XY3() {
	x := int((c8.opcode & 0x0F00) >> 8)
	y := int((c8.opcode & 0x00F0) >> 4)

	c8.register[x] = c8.register[x] ^ c8.register[y]

	c8.incrementPC(1)
}

// Adds VY to VX. VF is set to 1 when there's a carry, and to 0 when there isn't.
func (c8 *Chip8) x8XY4() {
	x := int((c8.opcode & 0x0F00) >> 8)
	y := int((c8.opcode & 0x00F0) >> 4)

	c8.register[x] += c8.register[y]

	if int(x)+int(y) > 0xFF {
		c8.register[0xF] = 1

	} else {
		c8.register[0xF] = 0

	}

	c8.incrementPC(1)

}

// VY is subtracted from VX. VF is set to 0 when there's a borrow, and 1 when there isn't.
func (c8 *Chip8) x8XY5() {
	x := int((c8.opcode & 0x0F00) >> 8)
	y := int((c8.opcode & 0x00F0) >> 4)

	c8.register[x] -= c8.register[y]

	if y > x {
		c8.register[0xF] = 0

	} else {
		c8.register[0xF] = 1

	}

	c8.incrementPC(1)

}

// Stores the least significant bit of VX in VF and then shifts VX to the right by 1.
func (c8 *Chip8) x8XY6() {
	x := int((c8.opcode & 0x0F00) >> 8)

	c8.register[0xF] = c8.register[x] & 0x000F
	c8.register[x] = c8.register[x] >> 1

	c8.incrementPC(1)

}

// Sets VX to VY minus VX. VF is set to 0 when there's a borrow, and 1 when there isn't.
func (c8 *Chip8) x8XY7() {
	x := int((c8.opcode & 0x0F00) >> 8)
	y := int((c8.opcode & 0x00F0) >> 4)

	c8.register[x] = c8.register[y] - c8.register[x]

	if x > y {
		c8.register[0xF] = 0

	} else {
		c8.register[0xF] = 1

	}

	c8.incrementPC(1)

}

// Stores the most significant bit of VX in VF and then shifts VX to the left by 1.
func (c8 *Chip8) x8XYE() {
	x := int((c8.opcode & 0x0F00) >> 8)

	c8.register[0xF] = c8.register[x] & 0x000F
	c8.register[x] = c8.register[x] << 1

	c8.incrementPC(1)

}

// Skips the next instruction if VX doesn't equal VY.
func (c8 *Chip8) x9XY0() {
	x := int((c8.opcode & 0x0F00) >> 8)
	y := int((c8.opcode & 0x00F0) >> 4)

	if c8.register[x] != c8.register[y] {
		c8.incrementPC(2)

	} else {
		c8.incrementPC(1)

	}

}

// Sets I to the address NNN.
func (c8 *Chip8) xANNN() {
	nnn := c8.opcode & 0x0FFF
	c8.i = nnn

	c8.incrementPC(1)

}

// Jumps to the address NNN plus V0.
func (c8 *Chip8) xBNNN() {
	nnn := c8.opcode & 0x0FFF
	c8.pc = nnn + uint16(c8.register[0])

}

// Sets VX to the result of a bitwise and operation on a random number (Typically: 0 to 0xFF) and NN.
func (c8 *Chip8) xCXNN() {
	x := int((c8.opcode & 0x0F00) >> 8)
	nn := c8.opcode & 0x00FF

	c8.register[x] = byte(nn) & byte(rand.Intn(0xFF))

	c8.incrementPC(1)

}

// Draws a sprite at coordinate (VX, VY) that has a width of 8 pixels and a height of N pixels.
// Each row of 8 pixels is read as bit-coded starting from memory location I; I value doesn’t change after the execution of this instruction.
// As described above, VF is set to 1 if any screen pixels are flipped from set to unset when the sprite is drawn, and to 0 if that doesn’t happen.
func (c8 *Chip8) xDXYN() {
	x := int((c8.opcode & 0x0F00) >> 8)
	y := int((c8.opcode & 0x00F0) >> 4)
	n := int((c8.opcode & 0x000F))

	vx := c8.register[x] // display x coordinate
	vy := c8.register[y] // display y coordinate

	start := int(64*vy + vx) // number of pixels since origin
	pos := start             // current pixel

	for yLine := 0; yLine < n; yLine++ {
		spr := c8.memory[int(c8.i)+yLine] // bit-coded sprite data

		for xLine := 0; xLine < 8; xLine++ {
			pos = start + (yLine * 64) + xLine

			pix := spr & (0x80 >> xLine) // scan through sprite, one pixel at a time

			if pix == 1 {
				if c8.gfx[pos] == 1 {
					c8.gfx[pos] = 0
					c8.register[0xF] = 1

				} else {
					c8.gfx[pos] = 1
					c8.register[0xF] = 0

				}
			}
		}
	}
	c8.drawFlag = true
	c8.incrementPC(1)

}

// Skips the next instruction if the key stored in VX is pressed.
func (c8 *Chip8) xEX9E() {
	x := int((c8.opcode & 0x0F00) >> 8)

	if c8.key[c8.register[x]] != 0 {
		c8.incrementPC(2)

	} else {
		c8.incrementPC(1)

	}

}

// Skips the next instruction if the key stored in VX isn't pressed.
func (c8 *Chip8) xEXA1() {
	x := int((c8.opcode & 0x0F00) >> 8)

	if c8.key[c8.register[x]] == 0 {
		c8.incrementPC(2)

	} else {
		c8.incrementPC(1)

	}

}

// Sets VX to the value of the delay timer.
func (c8 *Chip8) xFX07() {
	x := int((c8.opcode & 0x0F00) >> 8)

	c8.register[x] = c8.delayTimer

	c8.incrementPC(1)

}

// A key press is awaited, and then stored in VX. (Blocking Operation. All instruction halted until next key event).
func (c8 *Chip8) xFX0A() {
	x := int((c8.opcode & 0x0F00) >> 8)

	keyState := c8.key
	pressed := false

	for i := range keyState {
		if keyState[i] == 0 && c8.key[i] != 0 {
			c8.register[x] = byte(i)
			pressed = true

		}

	}

	if pressed {
		c8.incrementPC(1)

	}

}

// Sets the delay timer to VX.
func (c8 *Chip8) xFX15() {
	x := int((c8.opcode & 0x0F00) >> 8)

	c8.delayTimer = c8.register[x]

	c8.incrementPC(1)

}

// Sets the sound timer to VX.
func (c8 *Chip8) xFX18() {
	x := int((c8.opcode & 0x0F00) >> 8)

	c8.soundTimer = c8.register[x]

	c8.incrementPC(1)
}

// Adds VX to I. VF is set to 1 when there is a range overflow (I+VX>0xFFF), and to 0 when there isn't.
func (c8 *Chip8) xFX1E() {
	x := int((c8.opcode & 0x0F00) >> 8)

	c8.i += uint16(c8.register[x])

	if c8.i > 0xFFF {
		c8.register[0xF] = 1

	} else {
		c8.register[0xF] = 0

	}

	c8.incrementPC(1)

}

// Sets I to the location of the sprite for the character in VX. Characters 0-F (in hexadecimal) are represented by a 4x5 font.
func (c8 *Chip8) xFX29() {
	x := int((c8.opcode & 0x0F00) >> 8)

	c8.i = uint16(c8.register[x]) * 5

	c8.incrementPC(1)

}

// Stores the binary-coded decimal representation of VX, with the most significant of three digits at the address in I, the middle digit at I plus 1, and the least significant digit at I plus 2.
// (In other words, take the decimal representation of VX, place the hundreds digit in memory at location in I, the tens digit at location I+1, and the ones digit at location I+2.)
func (c8 *Chip8) xFX33() {
	x := int((c8.opcode & 0x0F00) >> 8)
	vx := c8.register[x]

	c8.memory[c8.i] = vx / 100
	c8.memory[c8.i+1] = (vx / 10) % 10
	c8.memory[c8.i+2] = (vx % 100) % 10

	c8.incrementPC(1)

}

// Stores V0 to VX (including VX) in memory starting at address I. The offset from I is increased by 1 for each value written, but I itself is left unmodified.
func (c8 *Chip8) xFX55() {
	x := int((c8.opcode & 0x0F00) >> 8)

	for i := 0; i <= x; i++ {
		c8.memory[int(c8.i)+i] = c8.register[i]

	}

	c8.incrementPC(1)

}

// Fills V0 to VX (including VX) with values from memory starting at address I. The offset from I is increased by 1 for each value written, but I itself is left unmodified.
func (c8 *Chip8) xFX65() {
	x := int((c8.opcode & 0x0F00) >> 8)

	for i := 0; i <= x; i++ {
		c8.register[i] = c8.memory[int(c8.i)+i]

	}

	c8.incrementPC(1)

}
