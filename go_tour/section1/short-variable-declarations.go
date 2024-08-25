package main

import "fmt"

func Short_variable_declatations() {
	// ":=" can be used only in functions
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}