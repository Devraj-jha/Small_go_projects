package main // Every executable Go program starts with package main

import (
	"fmt"     // Printing
	"strings" // String utility functions
)

// Struct Type
// A struct groups related data together
type User struct {
	Name string // field of type string
	Age  int    // field of type int
}
type Person struct {
	Name string 
	Age int
}
// Method
// A method is a function attached to a type
func (u User) Greet() {
	fmt.Printf("Hello, I'm %s and I'm %d years old\n", u.Name, u.Age)
}

// Function
// Takes an int and returns a bool
func isAdult(age int) bool {
	return age >= 18
}

func main() {

	// Short variable declaration
	// Go infers the type automatically
	name := "Devraj" // string
	age := 20        // int

	// Explicit variable declaration
	var city string = "Delhi"

	// Constant
	// Value cannot change
	const country = "India"

	fmt.Println(country)

	// Slice
	// Dynamic array that can grow/shrink
	skills := []string{
		"Go",
		"JavaScript",
		"Git",
	}

	// Map
	// Key-value data structure
	scores := map[string]int{
		"math":    95,
		"science": 88,
	}

	// Struct Instance
	// Creating a User value
	user := User{
		Name: name,
		Age:  age,
	}

	// Calling a method
	user.Greet()

	// If / Else
	if isAdult(user.Age) {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}

	// Traditional for loop
	for i := 0; i < len(skills); i++ {
		fmt.Println(skills[i])
	}

	// Range loop over slice
	// index = position
	// skill = value
	for index, skill := range skills {
		fmt.Println(index, skill)
	}

	// Range loop over map
	for subject, score := range scores {
		fmt.Println(subject, score)
	}

	// Switch statement
	switch city {
	case "Delhi":
		fmt.Println("Capital city")
	case "Mumbai":
		fmt.Println("Financial city")
	default:
		fmt.Println("Unknown city")
	}

	// Calling package function
	fmt.Println(strings.ToUpper(name))

	// Append adds element to slice
	skills = append(skills, "Docker")

	fmt.Println(skills)

	// Anonymous Function
	// Function stored inside a variable
	printMessage := func(msg string) {
		fmt.Println(msg)
	}

	printMessage("Learning Go")

	// Pointer Example
	x := 10

	// & gets memory address
	ptr := &x

	// * gets value stored at address
	fmt.Println(*ptr)

	// Function returning value + error
	value, err := divide(10, 2)

	// Standard Go error handling
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Result:", value)
}

// Function returning two values
// (result, error)
func divide(a, b int) (int, error) {

	// Guard clause
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	return a / b, nil
}