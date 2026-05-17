package main

import (
	"fmt"
	"math/rand/v2"
	
)
func main(){

	var num int
	total_attempt := 0 
	rand_num := rand.IntN(11)

	fmt.Println("--Guess The number( 0 to 10) ---")
	fmt.Println("Enter the number: ")
	fmt.Scan(&num)	
	for {

		if num == rand_num{
		fmt.Println("Correct !!")
		fmt.Println("total_attempt:", total_attempt)
		
		break
			
		}else if num > rand_num{
			fmt.Println("High")
			fmt.Scan(&num)
	
		}else {
			fmt.Println("Low")
			fmt.Scan(&num)
		}
		total_attempt++;
	}
}