package main

import "fmt"
import "time"

func Count(ch chan int) {
	fmt.Println("counting...")
	ch <- 1
}

func test_channel() {
	ch := make(chan int, 1)
	for {
		select {
		case ch <- 0:
		case  ch <- 1:
		}
		i := <-ch
		fmt.Println(i)
		time.Sleep(1e9)
	}
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}

	var sum int
	for _, ch := range(chs) {
		sum += <-ch
	}

	fmt.Println("The sum is: ", sum)

	var ch chan int
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(2e9)
		timeout <- true
	} ()

	select {
	case <-ch:
	case <-timeout:
	}
//	test_channel()
}
