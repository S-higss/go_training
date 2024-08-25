package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return	
	// -> naked return
	// It should only be used in short functions. 
	// Using it in long functions can negatively affect readability.
}

func Named_results() {
	x, y := split(17)
	message := fmt.Sprintf("let 17 split: %d and %d", x, y)
	fmt.Println(message)
}