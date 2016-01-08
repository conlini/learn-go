package main

// structs pointers
import ("fmt")

type Person struct {
	name string
	age int
}

func (p *Person) me() string {
	return p.name + " is my name"
}

func main() {
	i:= 42
	pi:= &i
	fmt.Println(pi)
	*pi = 111
	fmt.Println(i)
	dummy := Person{"dummy", 10}
	anotherDummy := Person{"second", 1}
	fmt.Println(dummy, anotherDummy)
	fmt.Println(dummy.name)

	d := &dummy
	*d = Person{"yogi bear", 25}
	fmt.Println(d, *d)
	fmt.Println(dummy.me())
}
