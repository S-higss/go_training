package main

import "fmt"

func Type_inference() {
	v := 42				//int
	f := 3.142			// float64
	g := 0.867 + 0.5i	// complex128
	fmt.Printf("v(=%d) is of type %T\n", v, v)
	fmt.Printf("f(=%f) is of type %T\n", f, f)
	fmt.Printf("g(=%f) is of type %T\n", g, g)
}