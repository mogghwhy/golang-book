package main

import "fmt"

func main() {
	fmt.Print("Enter degrees in fahrenhei:")
	var input float64
	fmt.Scanf("%f", &input)
	output := (input - 32.0) * (5.0 / 9.0)
	fmt.Printf("%v degrees fahrenhei is %2.2f degrees celsius\n",input,output)
}
