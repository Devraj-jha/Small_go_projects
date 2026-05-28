
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
		fmt.Println(i)
	}
}