package main

// vars const prints

import (
	"fmt"
)

var (
	ToBe   bool   = false
	MaxInt uint64 = 1<<64 - 1
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
}
