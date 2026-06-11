package main

//store data -> make decision -> repeat actions


import "fmt"

func main(){
	//? -> 1st example
	age := 20

	if age >= 18{
		fmt.Println("Adult")

	}else {
		fmt.Println("kid")
	}
	println(age == 20)

	//? -> 2nd example 

	marks := 75

	if marks >= 90 {
		fmt.Println("topper")
	}else if marks >= 70 {
		fmt.Println("average")
	}else {
		fmt.Println("billioanire")

	}
	class := "a"

	if  class == "a" {
		fmt.Println("a");
	}

}