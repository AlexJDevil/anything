package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello everyone, I'm Alex and I'm in mars...")

	fooGo()

	fmt.Println("something more please!!")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}

	n, _ := fmt.Println("Hello, Playground", 42, true)
	fmt.Println(n)
}

func fooGo() {
	fmt.Println("I'm in the loo doing some foo go baby!!")
}
