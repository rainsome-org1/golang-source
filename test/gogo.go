package main

import "fmt"
import "math"

func  slice_and_slice() {
	row := []int{1, 2, 3, 4}
	col := row[2:]

	fmt.Println("col[4]", col[1])
}

func IsEqual (f1, f2, p float64) bool {
	return math.Abs(f1 - f2) < p
}

func char_check(str string) bool {
	n := len(str)

	for i := 0; i< n; i++ {
		ch := str[i]
		fmt.Println(i, ch)
	}

	return true
}

func array_oper(str string) bool {
	for i:= 0; i < len(str); i++ {
		fmt.Printf("%d : %c", str[i], str[i])
	}

	return true
}

func use_range(str string) bool {
	for i, ch := range str {
		fmt.Println(i, ch)
	}

	return true
}

func value_test(array [3]int) bool {
	array[2] = 32;
	fmt.Println("Value_test---The array[2] is: ", array[2])

	return true
}

func my_slice(i int) bool {
	var array [4] int = [4]int{34, 43, 45, 2}
	var mySlice []int = array[:2]

	fmt.Printf("Element of array: \n")
	for _, num := range array {
		fmt.Print(num, "")
	}

	fmt.Printf("Element of mySlice: \n")
	for _, num := range mySlice {
		fmt.Print(num, " ")
	}

	slice := make([]int, 5)
	slice[3] = 1
	slice[4] = 32

	for i, num := range slice {
		fmt.Printf("The slice[%d] is : %d\n", i, num)
	}

	return true
}

func new_use_of_slice() {

	fmt.Println("The useage of slice append")
	slice := make([]int, 5, 10)
	other := []int{3, 45, 56}

	fmt.Println("len(array): ", len(slice))
	fmt.Println("cap(array): ", cap(slice))
	for i, num := range slice {
		fmt.Println("Old slice: ", i, num)
	}
	slice = append(slice, other...)
	for i, num := range slice {
		fmt.Println("New slice: ", i, num)
	}
}

func use_of_copy() {

	arr1 := []int{23, 34, 4, 34}
	arr2 := []int{34,67}

	copy(arr1, arr2)
	fmt.Println("copy(arr2 -> arr1)")
	for i, num := range arr1 {
		fmt.Println(i, num)
	}

	copy(arr2, arr1)
	fmt.Println("copy(arr1 -> arr2)")
	for i, num := range arr2 {
		fmt.Println(i, num)
	}
}

func main() {

	var str string;

	str = "hello"
	var array [3]int = [3]int{1, 2, 3}

	fmt.Printf("The origin string is \"%s\", its length is %d\n", str, len(str))
	fmt.Println("f1 is equal f2 ? : ", IsEqual(1.0, 1.0, 0.000001))

	fmt.Println(str + str)
	fmt.Printf("%c\n", "hello"[1])

	fmt.Printf("The char_check function")
	char_check("I LoVe YoU")

	fmt.Println("Array operation: ")
	array_oper("I love you")

	fmt.Println("Use of range")
	use_range("I LOVE YOU")

	fmt.Println("The origin array[2] is: ", array[2])
	value_test(array)
	fmt.Println("The changed array[2] is: ", array[2])

	my_slice(1)
	new_use_of_slice()

	slice_and_slice()

	use_of_copy()
}
