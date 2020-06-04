package main

func intInArray(array []int, value int) bool {
	for _, v := range array {
		if v == value {
			return true

		}

	}

	return false

}
