# CHIP-8 Emulator

Implementation of CHIP-8 written in Go.

## Usage

### Flags

      -audiofrequency float
	        Frequency of the audio tone. (default 200)
      -audiosamplerate int
            Sample rate for the audio. (default 44100)
      -audiovolume float
            Multiplier for audio volume, between 0 and 1. (default 0.5)
      -clockspeed int
            The number of cycles executed per second. (default 700)   
      -displayscale float
            Multiplier for screen size. '1' is 64x32. (default 8) 

### Emulation

| Key         	  | Description
|:------------------|:-----------------
| `F1`              | Reset the emulator
| `F2`     	        | Pause
| `F3`              | Continue
| `F4`              | Step (when paused)


### Key Mapping

| Keyboard        	                                                                | CHIP-8 Keypad
|:--------------------------------------------------------------------------------------|:-------------------------------------------------------------------------------------|
| `1` `2` `3` `4`**<br>**`Q` `W` `E` `R`**<br>**`A` `S` `D` `F`**<br>**`Z` `X` `C` `V`  | `1` `2` `3` `C`**<br>**`4` `5` `6` `D`**<br>**`7` `8` `9` `E`**<br>**`A` `0` `B` `F` |

