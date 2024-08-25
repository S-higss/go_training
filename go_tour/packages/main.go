package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	Packages()
	Imports()
	Exported_names()

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(100)
	rand.Seed(time.Now().UnixNano())
	y := rand.Intn(100)
	fmt.Printf("%d plus %d is %d\n", x, y, Add(x,y))
}