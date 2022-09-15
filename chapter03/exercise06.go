package main

import "fmt"

func main() {
	fmt.Print("Enter length in feet:")
	var input float64
	fmt.Scanf("%f", &input)
	output := (input * 0.3048)
	fmt.Printf("%v feet is %2.2f meters\n",input,output)
}
