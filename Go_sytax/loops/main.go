
package main

import "fmt"

func main() {
	num := 0

	for {
		fmt.Println(num)
		num++
		if num == 4 {
			break
		}
	}
	for i:= 0; i < 5; i++ {
		for j:= 0; j < 5; j++{
		fmt.Print(j);
	}
	fmt.Println(" ")
	}
	
}