package main

import (
	"fmt"
)
func greet(){
	fmt.Println("hello")
}
func add( a int, b int ) int {
	return a + b
}
func sub( a,b int) int {
	return a - b
}
func divide(a, b int) (int, int) {
    return a / b, a % b
}
func main(){
	greet()
	fmt.Println(add(3,4))
	q, r := divide(10, 3)
	fmt.Println(q,r)
}

//func name(parameters) return type {
// }