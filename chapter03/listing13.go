package main

import "fmt"

func main() {
	const x string = "Hello, World"
	x = "something else"
	fmt.Println(x)
}
