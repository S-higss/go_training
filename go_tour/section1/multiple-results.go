package main

import "fmt"

func swap(a, b string) (string, string) {
	return b, a
}

func Multiple_results() {
	a, b := "hello", "world"
	fmt.Println(a, b)
	fmt.Println("Swap!")
	fmt.Println(swap(a,b))
}