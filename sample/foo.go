package main

import (
	"fmt"
	"sample/sub"
)

func main(){
	fmt.Println("Hello world!")
	fmt.Println(sub.Add(3,5))
	sub.Hello()
}