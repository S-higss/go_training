package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe	bool		= false
	MaxInt	uint64		= 1<<64 - 1
	z		complex128	= cmplx.Sqrt(-5 + 12i)
)

func Basic_types() {
	fmt.Printf("ToBe(bool:false) -> Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("MaxInt(uint64:1<<64 - 1) -> Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("z(complex128:cmplx.Sqrt(-5 + 12i)) -> Type: %T Value: %v\n", z, z)
}

// The int, uint, and uintptr types are
// 32 bits on 32-bit systems and 64 bits on 64-bit systems. 
// Unless you have a specific reason to use a sized, unsigned integer type, 
// use int when you need an integer variable.