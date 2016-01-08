package main

import (
	"fmt"
)

var pack_level bool
var first, second bool = true, false 
func main() {
	var methodLevel bool
	methodLevel = true
	implicit := false
	fmt.Println(pack_level, methodLevel, implicit)
	fmt.Println(first, second)
}
