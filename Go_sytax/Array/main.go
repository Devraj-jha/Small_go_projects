package main

import (
	"fmt"
)


func main(){
 var numbers [5]int 
 //var name [size] data type
 numbers[0] = 10
numbers[1] = 20
numbers[2] = 30
numbers[3] = 40
numbers[4] = 50
fmt.Print(numbers[3])
num := [5]int{5,10,15,20,25};
var sum int = 0
for _,value := range num {
	sum = sum + value
}
	fmt.Println(sum)

}