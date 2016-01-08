package main

// for statments 
import (
	"fmt"
)

func main() {
	const f = "%v"
	for i := 10; i > 0; i-- {
		fmt.Printf(f, i)
	}

	sum := 1
	for sum < 5 {
		sum += sum
	}
	fmt.Printf(f, sum)
	if sum > 10 {
		fmt.Println("greater")
	} else {
		fmt.Println("lesser")
	}

	if sum = sum+10; sum > 10 {
		fmt.Println("greatness achieved")
	}
}
