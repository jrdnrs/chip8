package main

import (
	"fmt"
)

type test struct {
	c8 *Chip8
}

func runTests() {
	test := test{}
	test.c8 = &Chip8{}

	test.runAll()

}

func (t *test) runAll() {
	t.testx00E0()
	t.testx00EE()
	t.testx1NNN()
	t.testx2NNN()
	t.testx3XNN()
	t.testx4XNN()
	t.testx5XY0()
	t.testx6XNN()
	t.testx7XNN()
	t.testx8XY0()
	t.testx8XY1()
	t.testx8XY2()
	t.testx8XY3()
	t.testx8XY4()
	t.testx8XY5()
	t.testx8XY6()
	t.testx8XY7()
	t.testx8XYE()
	t.testx9XY0()
	t.testxANNN()
	t.testxBNNN()
	t.testxCXNN()
	t.testxDXYN()
	t.testxEX9E()
	t.testxEXA1()
	t.testxFX07()
	t.testxFX0A()
	t.testxFX15()
	t.testxFX18()
	t.testxFX1E()
	t.testxFX29()
	t.testxFX33()
	t.testxFX55()
	t.testxFX65()

}

func (t *test) testx00E0() {
	t.c8.Initialise()

	// Manually turn pixels on.
	for i := range t.c8.gfx {
		t.c8.gfx[i] = 1

	}

	t.c8.x00E0()

	// Check that all pixels were cleared (reset to 0).
	for i := range t.c8.gfx {
		if t.c8.gfx[i] != 0 {
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

	// Check that stack pointer was correctly decremented.
	if t.c8.sp != 0 {
		fmt.Printf("Opcode 0x00EE halted. Incorrect stack pointer; expected 0, received %d.\n", t.c8.sp)

		return

	}

	// Check that address was correctly popped from stack, loaded into pc, and pc incremented.
	if t.c8.pc != 0x304+0x2 {
		fmt.Printf("Opcode 0x00EE halted. Incorrect program counter; expected 0x306, received 0x%X.\n", t.c8.pc)

		return

	}

	fmt.Println("Opcode 0x00EE test passed.")

}

func (t *test) testx1NNN() {
	t.c8.Initialise()

	t.c8.opcode = 0x1304

	t.c8.x1NNN()

	// Check that pc was set to address specified in opcode.
	if t.c8.pc != 0x304 {
		fmt.Printf("Opcode 0x1NNN halted. Incorrect program counter; expected 0x304, received 0x%X.\n", t.c8.pc)

		return

	}

	fmt.Println("Opcode 0x1NNN test passed.")

}

func (t *test) testx2NNN() {
	t.c8.Initialise()

	t.c8.opcode = 0x2304

	t.c8.x2NNN()

	// Check that the current address was pushed onto stack.
	if t.c8.stack[0] != 0x200 {
		fmt.Printf("Opcode 0x2NNN halted. Incorrect value pushed to stack; expected 0x200, received 0x%X.\n", t.c8.stack[0])

		return

	}

	// Check that stack pointer was incrememented correctly.
	if t.c8.sp != 1 {
		fmt.Printf("Opcode 0x2NNN halted. Incorrect stack pointer; expected 1, received %d.\n", t.c8.sp)

		return

	}

	// Check that pc was set to address specified in opcode.
	if t.c8.pc != 0x304 {
		fmt.Printf("Opcode 0x2NNN halted. Incorrect program counter; expected 0x304, received 0x%X.\n", t.c8.pc)

		return

	}

	fmt.Println("Opcode 0x2NNN test passed.")

}

func (t *test) testx3XNN() {
	t.c8.Initialise()

	t.c8.opcode = 0x3042
	t.c8.register[0] = 0x42

	t.c8.x3XNN()

	// Checks that next address was skipped.
	if t.c8.pc != 0x204 {
		fmt.Printf("Opcode 0x3XNN halted. Incorrect program counter; expected 0x204, received 0x%X.\n", t.c8.pc)

		return

	}

	fmt.Println("Opcode 0x3XNN test passed.")

}

func (t *test) testx4XNN() {
	t.c8.Initialise()

	t.c8.opcode = 0x4042
	t.c8.register[0] = 0x27

	t.c8.x4XNN()

	// Checks that next address was skipped.
	if t.c8.pc != 0x204 {
		fmt.Printf("Opcode 0x4XNN halted. Incorrect program counter; expected 0x204, received 0x%X.\n", t.c8.pc)

		return

	}

	fmt.Println("Opcode 0x4XNN test passed.")

}

func (t *test) testx5XY0() {
	t.c8.Initialise()

	t.c8.opcode = 0x5010
	t.c8.register[0] = 0x42
	t.c8.register[1] = 0x42

	t.c8.x5XY0()

	// Checks that next address was skipped.
	if t.c8.pc != 0x204 {
		fmt.Printf("Opcode 0x5XY0 halted. Incorrect program counter; expected 0x204, received 0x%X.\n", t.c8.pc)

		return

	}

	fmt.Println("Opcode 0x5XY0 test passed.")

}

func (t *test) testx6XNN() {
	t.c8.Initialise()

	t.c8.opcode = 0x6042

	t.c8.x6XNN()

	// Checks that VX was correctly modified.
	if t.c8.register[0] != 0x42 {
		fmt.Printf("Opcode 0x6XNN halted. Incorrect register[0]; expected 0x42, received 0x%X.\n", t.c8.register[0])

		return

	}

	fmt.Println("Opcode 0x6XNN test passed.")

}

func (t *test) testx7XNN() {
	t.c8.Initialise()

	t.c8.opcode = 0x7042

	t.c8.x7XNN()

	// Checks that NN was correctly added to VX.
	if t.c8.register[0] != 0x42 {
		fmt.Printf("Opcode 0x7XNN halted. Incorrect register[0]; expected 0x42, received 0x%X.\n", t.c8.register[0])

		return

	}

	// Checks for change in carry flag (VF).
	if t.c8.register[0xF] != 0 {
		fmt.Printf("Opcode 0x7XNN halted. Incorrect register[0xF]; expected 0, received 0x%X.\n", t.c8.register[0xF])

		return

	}

	fmt.Println("Opcode 0x7XNN test passed.")

}

func (t *test) testx8XY0() {
	t.c8.Initialise()

	t.c8.opcode = 0x8010
	t.c8.register[1] = 0x42

	t.c8.x8XY0()

	// Checks that VX was set to VY.
	if t.c8.register[0] != 0x42 {
		fmt.Printf("Opcode 0x8XY0 halted. Incorrect register[0]; expected 0x42, received 0x%X.\n", t.c8.register[0])

		return

	}

	fmt.Println("Opcode 0x8XY0 test passed.")

}

func (t *test) testx8XY1() {
	t.c8.Initialise()

	t.c8.opcode = 0x8011
	t.c8.register[0] = 0x9
	t.c8.register[1] = 0x2

	t.c8.x8XY1()

	// Checks that OR was performed correctly with VX set to result.
	if t.c8.register[0] != 0xB {
		fmt.Printf("Opcode 0x8XY1 halted. Incorrect register[0]; expected 0xB, received 0x%X.\n", t.c8.register[0])

		return

	}

	fmt.Println("Opcode 0x8XY1 test passed.")

}

func (t *test) testx8XY2() {
	t.c8.Initialise()

	t.c8.opcode = 0x8011
	t.c8.register[0] = 0xB
	t.c8.register[1] = 0xE

	t.c8.x8XY2()

	// Checks that AND was performed correctly with VX set to result.
	if t.c8.register[0] != 0xA {
		fmt.Printf("Opcode 0x8XY2 halted. Incorrect register[0]; expected 0xA, received 0x%X.\n", t.c8.register[0])

		return

	}

	fmt.Println("Opcode 0x8XY2 test passed.")

}

func (t *test) testx8XY3() {
	t.c8.Initialise()

	t.c8.opcode = 0x8011
	t.c8.register[0] = 0xA
	t.c8.register[1] = 0xF

	t.c8.x8XY3()

	// Checks that XOR was performed correctly with VX set to result.
	if t.c8.register[0] != 0x5 {
		fmt.Printf("Opcode 0x8XY3 halted. Incorrect register[0]; expected 0x5, received 0x%X.\n", t.c8.register[0])

		return

	}

	fmt.Println("Opcode 0x8XY3 test passed.")

}

func (t *test) testx8XY4() {
	t.c8.Initialise()

	t.c8.opcode = 0x8014
	t.c8.register[0] = 0xFE
	t.c8.register[1] = 0xF0

	t.c8.x8XY4()

	// Checks that VY was added to VX and overflowed as expected.
	if t.c8.register[0] != 0xEE {
		fmt.Printf("Opcode 0x8XY4 halted. Incorrect register[0]; expected 0xEE, received 0x%X.\n", t.c8.register[0])

		return

	}

	// Checks that carry flag (VF) was set.
	if t.c8.register[0xF] != 1 {
		fmt.Printf("Opcode 0x8XY4 halted. Incorrect register[0xF]; expected 1, received %d.\n", t.c8.register[0xF])

		return

	}

	// Reset to now check with no overflow.
	t.c8.Initialise()

	t.c8.opcode = 0x8014
	t.c8.register[0] = 0x04
	t.c8.register[1] = 0x20

	t.c8.x8XY4()

	// Checks that VY was added to VX and did not overflow.
	if t.c8.register[0] != 0x24 {
		fmt.Printf("Opcode 0x8XY4 halted. Incorrect register[0]; expected 0x24, received 0x%X.\n", t.c8.register[0])

		return

	}

	// Checks that carry flag (VF) was not set.
	if t.c8.register[0xF] != 0 {
		fmt.Printf("Opcode 0x8XY4 halted. Incorrect register[0xF]; expected 0, received %d.\n", t.c8.register[0xF])

		return

	}

	fmt.Println("Opcode 0x8XY4 test passed.")

}

func (t *test) testx8XY5() {
	t.c8.Initialise()

	t.c8.opcode = 0x8015
	t.c8.register[0] = 0xFE
	t.c8.register[1] = 0xF0

	t.c8.x8XY5()

	// Checks that VY was added to VX and doesn't go negative.
	if t.c8.register[0] != 0xE {
		fmt.Printf("Opcode 0x8XY5 halted. Incorrect register[0]; expected 0xE, received 0x%X.\n", t.c8.register[0])

		return

	}

	// Checks that borrow flag (VF) was set to 1 to indicate that there wasn't a borrow.
	if t.c8.register[0xF] != 1 {
		fmt.Printf("Opcode 0x8XY5 halted. Incorrect register[0xF]; expected 1, received %d.\n", t.c8.register[0xF])

		return

	}

	// Reset to now check with negative.
	t.c8.Initialise()

	t.c8.opcode = 0x8015
	t.c8.register[0] = 0x04
	t.c8.register[1] = 0x20

	t.c8.x8XY5()

	// Checks that VY was added to VX and did go negative.
	if t.c8.register[0] != 0xE4 {
		fmt.Printf("Opcode 0x8XY5 halted. Incorrect register[0]; expected 0xE4, received 0x%X.\n", t.c8.register[0])

		return

	}

	// Checks that carry flag (VF) was set to 0 to indicate that there was a borrow.
	if t.c8.register[0xF] != 0 {
		fmt.Printf("Opcode 0x8XY5 halted. Incorrect register[0xF]; expected 0, received %d.\n", t.c8.register[0xF])

		return

	}

	fmt.Println("Opcode 0x8XY5 test passed.")

}

func (t *test) testx8XY6() {
	t.c8.Initialise()

	t.c8.opcode = 0x8016
	t.c8.register[0] = 0xEF

	t.c8.x8XY6()

	// Checks that VX was correctly shifted to the right by one.
	if t.c8.register[0] != 0x77 {
		fmt.Printf("Opcode 0x8XY6 halted. Incorrect register[0]; expected 0x77, received 0x%X.\n", t.c8.register[0])

		return

	}

	// Checks that least significant bit of VX was stored in VF.
	if t.c8.register[0xF] != 1 {
		fmt.Printf("Opcode 0x8XY6 halted. Incorrect register[0xF]; expected 1, received %d.\n", t.c8.register[0xF])

		return

	}

	fmt.Println("Opcode 0x8XY6 test passed.")

}

func (t *test) testx8XY7() {
	t.c8.Initialise()

	t.c8.opcode = 0x8017
	t.c8.register[0] = 0xF0
	t.c8.register[1] = 0xFE

	t.c8.x8XY7()

	// Checks that VY was added to VX and doesn't go negative.
	if t.c8.register[0] != 0xE {
		fmt.Printf("Opcode 0x8XY7 halted. Incorrect register[0]; expected 0xE, received 0x%X.\n", t.c8.register[0])

		return

	}

	// Checks that borrow flag (VF) was set to 1 to indicate that there wasn't a borrow.
	if t.c8.register[0xF] != 1 {
		fmt.Printf("Opcode 0x8XY7 halted. Incorrect register[0xF]; expected 1, received %d.\n", t.c8.register[0xF])

		return

	}

	// Reset to now check with negative.
	t.c8.Initialise()

	t.c8.opcode = 0x8017
	t.c8.register[0] = 0x20
	t.c8.register[1] = 0x04

	t.c8.x8XY7()

	// Checks that VY was added to VX and did go negative.
	if t.c8.register[0] != 0xE4 {
		fmt.Printf("Opcode 0x8XY7 halted. Incorrect register[0]; expected 0xE4, received 0x%X.\n", t.c8.register[0])

		return

	}

	// Checks that carry flag (VF) was set to 0 to indicate that there was a borrow.
	if t.c8.register[0xF] != 0 {
		fmt.Printf("Opcode 0x8XY7 halted. Incorrect register[0xF]; expected 0, received %d.\n", t.c8.register[0xF])

		return

	}

	fmt.Println("Opcode 0x8XY7 test passed.")

}
func (t *test) testx8XYE() {
	t.c8.Initialise()

	t.c8.opcode = 0x801E
	t.c8.register[0] = 0xEF

	t.c8.x8XYE()

	// Checks that VX was correctly shifted to the left by one.
	if t.c8.register[0] != 0xDE {
		fmt.Printf("Opcode 0x8XYE halted. Incorrect register[0]; expected 0xDE, received 0x%X.\n", t.c8.register[0])

		return

	}

	// Checks that most significant bit of VX was stored in VF.
	if t.c8.register[0xF] != 1 {
		fmt.Printf("Opcode 0x8XYE halted. Incorrect register[0xF]; expected 1, received %d.\n", t.c8.register[0xF])

		return

	}

	fmt.Println("Opcode 0x8XYE test passed.")
}

func (t *test) testx9XY0() {
	t.c8.Initialise()

	t.c8.opcode = 0x9010
	t.c8.register[0] = 0x42

	t.c8.x9XY0()

	// Checks that next address was skipped.
	if t.c8.pc != 0x204 {
		fmt.Printf("Opcode 0x9XY0 halted. Incorrect program counter; expected 0x204, received 0x%X.\n", t.c8.pc)

		return

	}

	fmt.Println("Opcode 0x9XY0 test passed.")

}

func (t *test) testxANNN() {
	t.c8.Initialise()

	t.c8.opcode = 0xA304

	t.c8.xANNN()

	// Checks that the index register was set correctly to NNN.
	if t.c8.i != 0x304 {
		fmt.Printf("Opcode 0xANNN halted. Incorrect program counter; expected 0x304, received 0x%X.\n", t.c8.i)

		return

	}

	fmt.Println("Opcode 0xANNN test passed.")

}

func (t *test) testxBNNN() {
	t.c8.Initialise()

	t.c8.opcode = 0xB304
	t.c8.register[0] = 0x5

	t.c8.xBNNN()

	// Checks that the pc was set correctly to NNN + V0.
	if t.c8.pc != 0x309 {
		fmt.Printf("Opcode 0xBNNN halted. Incorrect program counter; expected 0x309, received 0x%X.\n", t.c8.i)

		return

	}

	fmt.Println("Opcode 0xBNNN test passed.")

}

func (t *test) testxCXNN() {
	t.c8.Initialise()

}

func (t *test) testxDXYN() {
	t.c8.Initialise()

	printDisplay := func() {
		fmt.Printf("Printing display...\n\n")

		for i, pixel := range t.c8.gfx {
			if pixel == 0 {
				fmt.Printf("░")
			} else {
				fmt.Printf("▓")
			}

			if (i+1)%64 == 0 {
				fmt.Printf("\n")
			}

		}
	}

	t.c8.opcode = 0xD015
	t.c8.i = 5            // memory address for '0' from fontset
	t.c8.register[0] = 23 // x coordinate
	t.c8.register[1] = 12 // y coordinate

	t.c8.xDXYN()

	targetPixels := []int{793, 856, 857, 921, 985, 1048, 1049, 1050} // pixels that should be set to 1
	var isTargetPixel bool

	// Checks that only target pixels are set to 1.
	for i, pixel := range t.c8.gfx {
		isTargetPixel = intInArray(targetPixels, i)

		if pixel == 1 && !isTargetPixel || pixel == 0 && isTargetPixel {
			fmt.Printf("Opcode 0xDXYN halted. Incorrect pixel %d; received %d.\n", i, pixel)
			printDisplay()

			return
		}
	}

	// Checks that the collision flag is correctly set to 0.
	if t.c8.register[0xF] != 0 {
		fmt.Printf("Opcode 0xDXYN halted. Incorrect register 0xF; expected 0, received 0x%X.\n", t.c8.register[0xF])

		return

	}

	fmt.Println("Opcode 0xDXYN test passed.")

}

func (t *test) testxEX9E() {
	t.c8.Initialise()

	t.c8.opcode = 0xE09E
	t.c8.register[0] = 4
	t.c8.key[4] = 1

	t.c8.xEX9E()

	// Checks that next address was skipped.
	if t.c8.pc != 0x204 {
		fmt.Printf("Opcode 0xEX9E halted. Incorrect program counter; expected 0x204, received 0x%X.\n", t.c8.pc)

		return

	}

	fmt.Println("Opcode 0xEX9E test passed.")

}

func (t *test) testxEXA1() {
	t.c8.Initialise()

	t.c8.opcode = 0xE0A1
	t.c8.register[0] = 4
	t.c8.key[4] = 0

	t.c8.xEXA1()

	// Checks that next address was skipped.
	if t.c8.pc != 0x204 {
		fmt.Printf("Opcode 0xEXA1 halted. Incorrect program counter; expected 0x204, received 0x%X.\n", t.c8.pc)

		return

	}

	fmt.Println("Opcode 0xEXA1 test passed.")

}

func (t *test) testxFX07() {
	t.c8.Initialise()

	t.c8.opcode = 0xF007
	t.c8.delayTimer = 7

	t.c8.xFX07()

	// Checks that VX is equal to delay timer.
	if t.c8.register[0] != 7 {
		fmt.Printf("Opcode 0xFX07 halted. Incorrect register 0; expected 7, received %d.\n", t.c8.register[0])

		return

	}

	fmt.Println("Opcode 0xFX07 test passed.")

}

func (t *test) testxFX0A() {
	t.c8.Initialise()

	t.c8.opcode = 0xF00A

	// Checks that program counter isn't incremented when a key isn't pressed.
	for i := 0; i < 10; i++ {
		t.c8.xFX0A()

		if t.c8.pc != 0x200 {
			fmt.Printf("Opcode 0xFX0A halted. Incorrect program counter; expected 0x200, received 0X%X.\n", t.c8.pc)

			return

		}
	}

	t.c8.key[7] = 1

	t.c8.xFX0A()

	// Checks that program counter increments now that a key has been pressed.
	if t.c8.pc != 0x202 {
		fmt.Printf("Opcode 0xFX0A halted. Incorrect program counter; expected 0x202, received 0X%X.\n", t.c8.pc)

		return

	}

	// Checks that VX is set to the key index.
	if t.c8.register[0] != 7 {
		fmt.Printf("Opcode 0xFX0A halted. Incorrect register 0; expected 7, received %d.\n", t.c8.register[0])

		return

	}

	fmt.Println("Opcode 0xFX0A test passed.")

}

func (t *test) testxFX15() {
	t.c8.Initialise()

	t.c8.opcode = 0xF015
	t.c8.register[0] = 7

	t.c8.xFX15()

	// Checks that delay timer is equal to VX.
	if t.c8.delayTimer != 7 {
		fmt.Printf("Opcode 0xFX15 halted. Incorrect delay timer; expected 7, received %d.\n", t.c8.delayTimer)

		return

	}

	fmt.Println("Opcode 0xFX15 test passed.")

}

func (t *test) testxFX18() {
	t.c8.Initialise()

	t.c8.opcode = 0xF018
	t.c8.register[0] = 7

	t.c8.xFX18()

	// Checks that sound timer is equal to VX.
	if t.c8.soundTimer != 7 {
		fmt.Printf("Opcode 0xFX18 halted. Incorrect sound timer; expected 7, received %d.\n", t.c8.soundTimer)

		return

	}

	fmt.Println("Opcode 0xFX18 test passed.")

}

func (t *test) testxFX1E() {
	t.c8.Initialise()

	t.c8.opcode = 0xF01E
	t.c8.register[0] = 0xFF
	t.c8.i = 0xFF0

	t.c8.xFX1E()

	// Checks that addition of VX to index register is correct.
	if t.c8.i != 0x10EF {
		fmt.Printf("Opcode 0xFX1E halted. Incorrect index register; expected 0x10EF, received 0x%X.\n", t.c8.i)

		return

	}

	// Checks that range overflow flag (VF) is correctly set to 1.
	if t.c8.register[0xF] != 1 {
		fmt.Printf("Opcode 0xFX1E halted. Incorrect register 0xF; expected 1, received %d.\n", t.c8.register[0xF])

		return

	}

	// Reset to check without range overflow.
	t.c8.Initialise()

	t.c8.opcode = 0xF01E
	t.c8.register[0] = 0x5
	t.c8.i = 0xF

	t.c8.xFX1E()

	// Checks that addition of VX to index register is correct.
	if t.c8.i != 0x14 {
		fmt.Printf("Opcode 0xFX1E halted. Incorrect index register; expected 0x14, received 0x%X.\n", t.c8.i)

		return

	}

	// Checks that range overflow flag (VF) is correctly set to 0.
	if t.c8.register[0xF] != 0 {
		fmt.Printf("Opcode 0xFX1E halted. Incorrect register 0xF; expected 0, received %d.\n", t.c8.register[0xF])

		return

	}

	fmt.Println("Opcode 0xFX1E test passed.")

}

func (t *test) testxFX29() {
	t.c8.Initialise()

	t.c8.opcode = 0xF029
	t.c8.register[0] = 4

	t.c8.xFX29()

	if t.c8.i != 20 {
		fmt.Printf("Opcode 0xFX29 halted. Incorrect register 0; expected 20, received %d.\n", t.c8.register[0])

		return

	}

	fmt.Println("Opcode 0xFX29 test passed.")

}

func (t *test) testxFX33() {
	t.c8.Initialise()

	t.c8.opcode = 0xF033
	t.c8.register[0] = 204

	t.c8.xFX33()

	if t.c8.memory[0] != 2 {
		fmt.Printf("Opcode 0xFX33 halted. Incorrect memory 0x0; expected 2, received %d.\n", t.c8.memory[0])

		return
	}

	if t.c8.memory[1] != 0 {
		fmt.Printf("Opcode 0xFX33 halted. Incorrect memory 0x1; expected 0, received %d.\n", t.c8.memory[1])

		return
	}

	if t.c8.memory[2] != 4 {
		fmt.Printf("Opcode 0xFX33 halted. Incorrect memory 0x2; expected 4, received %d.\n", t.c8.memory[2])

		return
	}

	fmt.Println("Opcode 0xFX33 test passed.")

}

func (t *test) testxFX55() {
	t.c8.Initialise()

	t.c8.opcode = 0xFE55

	values := []byte{4, 5, 3, 6, 3, 2, 8, 9, 1, 4, 0, 2, 6, 8, 5}

	for i, v := range values {
		t.c8.register[i] = v
	}

	t.c8.xFX55()

	for i, v := range values {
		if t.c8.memory[i] != v {
			fmt.Printf("Opcode 0xFX55 halted. Incorrect memory 0x%X; expected %d, received %d.\n", i, v, t.c8.memory[i])

			return

		}
	}

	fmt.Println("Opcode 0xFX55 test passed.")

}

func (t *test) testxFX65() {
	t.c8.Initialise()

	t.c8.opcode = 0xFE65

	values := []byte{4, 5, 3, 6, 3, 2, 8, 9, 1, 4, 0, 2, 6, 8, 5}

	for i, v := range values {
		t.c8.memory[i] = v
	}

	t.c8.xFX65()

	for i, v := range values {
		if t.c8.register[i] != v {
			fmt.Printf("Opcode 0xFX65 halted. Incorrect register 0x%X; expected %d, received %d.\n", i, v, t.c8.register[i])

			return

		}
	}

	fmt.Println("Opcode 0xFX65 test passed.")

}
