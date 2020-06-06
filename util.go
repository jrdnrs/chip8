package main

import (
	"io/ioutil"
)

func intInArray(array []int, value int) bool {
	for _, v := range array {
		if v == value {
			return true

		}

	}

	return false

}

func openRom(path string) ([]byte, error) {
	rom, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return rom, nil

}