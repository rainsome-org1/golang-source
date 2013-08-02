package main

import (
	"fmt"
)

func main() {

	f := func(x, y int) int {
		return x - y
	}

	// it made func equals f, then you can use f to call the function
	// so you can pass th eparamer to control the function
	fmt.Println(f(3, 4))

	var j int = 5

	a := func() (func()) {
		var i int = 10
		return func() {
			i, j = j, i
			fmt.Printf("%d : %d\n", i, j)
		}
	}()

	a()
}
