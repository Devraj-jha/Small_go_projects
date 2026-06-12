// a struct is go, is a way to group. related data together.
// isntead of manually, defing things, we can create struct.
// define it and put varaibles in it.

// name:= "devraj "
// age:= 19
package main

import "fmt"
type student struct {
	Name string
	Age int
	Marks int 
}

// creating a function and paassing structs to functions. 

func printStudent(s student){
	println("Name: ", s.Name)
	println("Age", s.Age)
}

// pointers to struct



func main(){
	//1st method to create a struct, pre initialization
	student1 := student{
		Name: "dj",
		Age: 19,
		Marks: 100,
	}
	ptr := &student1
	println(ptr.Name)
	fmt.Println(student1.Age) // TO ACCESS THESE VALUES. 
	//2nd method -> create empty, struct
 
	var student2 student // similar to how variables are declared. 
	student2.Name = "leon"
	student2.Age = 23
	println(student2.Age)
	printStudent(student2)
	


}
