package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Add(x, y int) int {
	return x + y
}
// func Add(x int, y int) int {
// 	return x + y
// }

func Functions() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(100)
	rand.Seed(time.Now().UnixNano())
	y := rand.Intn(100)
	fmt.Printf("%d plus %d is %d\n", x, y, Add(x,y))
}