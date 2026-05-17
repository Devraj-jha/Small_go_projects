package main

import (
	"fmt"
	"math/rand/v2"
	
)
func main(){

	var num int
	rand_num := rand.IntN(101)

	fmt.Println("--Guess The number( 0 to 100) ---")
	fmt.Println("Enter the number: ")
	fmt.Scan(&num)	
	for {

		if num == rand_num{
		fmt.Println("Correct !!")
		break
			
		}else if num > rand_num{
			fmt.Println("High")
			fmt.Scan(&num)
	
		}else {
			fmt.Println("Low")
			fmt.Scan(&num)
		}
	}
}