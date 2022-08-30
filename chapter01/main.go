package main

/*
creates an executable program that references the fmt library and contains
one function. that function takes no arguments and doesn't return anything.
It accesses the println function contained inside the fmt package and
invokes it using one argument - the string Hello, World
*/

// import fmt package
import "fmt"

func main() {
	fmt.Println("Hello, World")
}
