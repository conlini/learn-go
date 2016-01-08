package main

// functions and retunrs
import (
	"fmt"
)

func swap(val1, val2 string) (string, string) {
	return val2, val1
}

func naked_return_swap(val1, val2 string) (first, second string) {
	first = val2
	second = val1
	return
}

func main() {
	fmt.Println(swap("ball", "base"))
	fmt.Println(naked_return_swap("Potter", "Harry"))
}
