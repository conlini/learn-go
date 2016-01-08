package main

import (
	"fmt"
	"math/rand"
)

func compute(x, y int) int {
	rand.Seed(100)
	return ((x + y) * rand.Intn(100))
}

func main() {
	fmt.Println(compute(10, 12))
}
