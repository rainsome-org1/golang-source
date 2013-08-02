package main

import "fmt"

func testint(args ...int) {
	for _, num := range args {
		fmt.Println("testint: ", num)
	}
}

func testvar(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is a int type")
		case float64:
			fmt.Println(arg, "is a float64 type")
		case string:
			fmt.Println(arg, "is a string type")
		default :
			fmt.Println(arg, "unknow type")
		}
	}
}

func main() {
	var array [4]int = [4]int{3, 45, 56, 1}

	fmt.Println("test separet array element")
	testint(array[0], array[1])

	fmt.Println("test useing a array")

	fmt.Println("test varible arguments")
	testvar(2, "youngcy", 23.32)
}
