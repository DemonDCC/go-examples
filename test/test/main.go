package main

import "fmt"

func main() {
	testMapLength()
}

func testChannelClose() {
	ch := make(chan int, 1)

	ch <- 1
	close(ch)

	fmt.Println(<-ch)
}

func testMapLength() {
	m := make(map[int]string)

	m[1] = "demon"
	m[2] = "abser"
	m[3] = "caicai"

	fmt.Println(len(m))
}
