package main

import (
	"fmt"
)

type test struct {
	c8 *Chip8
}

func tester() {
	test := test{}
	test.c8 = &Chip8{}

	test.testx00E0()
	test.testx00EE()

}

func (t *test) testx00E0() {
	t.c8.Initialise()

	// Manually turn pixels on.
	for i := range t.c8.gfx {
		t.c8.gfx[i] = 1

	}

	t.c8.x00E0()

	for i := range t.c8.gfx {
		if t.c8.gfx[i] == 1 {
			fmt.Printf("Opcode 0x00E0 halted. Pixel %d was not cleared.\n", i)
			return

		}

	}

	fmt.Println("Opcode 0x00E0 test passed.")

}

func (t *test) testx00EE() {
	t.c8.Initialise()

	// Manually add address to stack.
	t.c8.stack[t.c8.sp] = 0x304
	t.c8.sp++

	t.c8.x00EE()

	if t.c8.sp == 0 {
		if t.c8.pc == 0x304+0x2 {
			fmt.Println("Opcode 0x00EE test passed.")

		} else {
			fmt.Printf("Opcode 0x00EE halted. Incorrect program counter; expected 0x306, received 0x%X.\n", t.c8.pc)

		}

	} else {
		fmt.Printf("Opcode 0x00EE halted. Incorrect stack pointer; expected 0, received %d.\n", t.c8.sp)

	}

}

func (t *test) testx1NNN() {
	t.c8.Initialise()

	t.c8.opcode = 0x1304

	t.c8.x1NNN()

	if t.c8.pc == 0x304 {
		fmt.Println("Opcode 0x1NNN test passed.")

	} else {
		fmt.Printf("Opcode 0x1NNN halted. Incorrect program counter; expected 0x304, received 0x%X.\n", t.c8.pc)

	}

}

func (t *test) testx2NNN() {
	t.c8.Initialise()

}

func (t *test) testx3NNN() {
	t.c8.Initialise()

}

func (t *test) testx4NNN() {
	t.c8.Initialise()

}

func (t *test) testx5XY0() {
	t.c8.Initialise()

}

func (t *test) testx6XNN() {
	t.c8.Initialise()

}

func (t *test) testx7XNN() {
	t.c8.Initialise()

}

func (t *test) testx8XY0() {
	t.c8.Initialise()

}

func (t *test) testx8XY1() {
	t.c8.Initialise()

}

func (t *test) testx8XY2() {
	t.c8.Initialise()

}

func (t *test) testx8XY3() {
	t.c8.Initialise()

}

func (t *test) testx8XY4() {
	t.c8.Initialise()

}

func (t *test) testx8XY5() {
	t.c8.Initialise()

}

func (t *test) testx8XY6() {
	t.c8.Initialise()

}

func (t *test) testx8XY7() {
	t.c8.Initialise()

}
func (t *test) testx8XYE() {
	t.c8.Initialise()

}

func (t *test) testx9XY0() {
	t.c8.Initialise()

}

func (t *test) testxANNN() {
	t.c8.Initialise()

}

func (t *test) testxBNNN() {
	t.c8.Initialise()

}

func (t *test) testxCXNN() {
	t.c8.Initialise()

}

func (t *test) testxDXYN() {
	t.c8.Initialise()

}

func (t *test) testxEX9E() {
	t.c8.Initialise()

}

func (t *test) testxEXA1() {
	t.c8.Initialise()

}

func (t *test) testxFX07() {
	t.c8.Initialise()

}

func (t *test) testxFX0A() {
	t.c8.Initialise()

}

func (t *test) testxFX15() {
	t.c8.Initialise()

}

func (t *test) testxFX18() {
	t.c8.Initialise()

}

func (t *test) testxFX1E() {
	t.c8.Initialise()

}

func (t *test) testxFX29() {
	t.c8.Initialise()

}

func (t *test) testxFX33() {
	t.c8.Initialise()

}

func (t *test) testxFX55() {
	t.c8.Initialise()

}

func (t *test) testxFX65() {
	t.c8.Initialise()

}