package main

import "fmt"

type Rect struct {
	x, y float64
	width, height float64
}

func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}

func (rect *Rect)Area() float64 {
	return rect.x * rect.y
}

type Integer int

func (a Integer) Less(b Integer) Integer {
	return a - b
}

func (a *Integer) Add(b Integer) {
	*a+=b
}

type LessAdd interface {
	Less(a Integer) Integer
	Add(b Integer)
}

func main() {
	rect := new(Rect)
	rect.x = 434.45

	fmt.Println("Using new(Rect): ", rect.x)

	rect1 := &Rect{x: 23, width: 34}
	fmt.Println("Using &Rect: ", rect1.x, rect1. width)

	rect2 := NewRect(3,4, 3, 4)
	fmt.Println("Using &Rect: ", rect2.x, rect2. width)

	var a Integer = 2
	fmt.Println(a.Less(34))
	a.Add(3)
	fmt.Println(a)

	fmt.Println("The usage of interface: ")
	var b LessAdd = &a
	fmt.Println(b)
	fmt.Println(b.Less(1))
	b.Add(10)
	fmt.Println(b)


}
