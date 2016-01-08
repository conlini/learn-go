package main
// Case statement

import (
	"fmt"
	"runtime"
)

var (
	os string = runtime.GOOS
)

const ()

func myfuntion(x int, y string) (string, int, string) {
	return y, x, y
}

func main() {
	switch os {
	case "darwin":
		fmt.Println("os x")
	default:
		fmt.Println(os)
	}

	defer fmt.Println("Hello you fool")
	fmt.Println(myfuntion(1, "a"))
}
