package main

import (
	"fmt"
)

func main() {
	var op string 
	var a, b float64

	fmt.Print("Enter operation (+, -, *, /): ")
	fmt.Scan(&op)

	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)

	switch op {
	case "+":
		fmt.Println("Result:", a+b)
	case "-":
		fmt.Println("Result:", a-b)
	case "*":
		fmt.Println("Result:", a*b)
	case "/":
		if b == 0 {
			fmt.Println("Cannot divide by zero")
			return
		}
		fmt.Println("Result:", a/b)
	default:
		fmt.Println("Invalid operation")
	}
}
//A package => folder related go code

//