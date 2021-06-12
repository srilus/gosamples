package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Println(add(2, 3))
}

//first commit
func add(x, y int) (z int) {
	z = x + y
	return
}
