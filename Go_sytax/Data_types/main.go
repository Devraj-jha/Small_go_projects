package main

import "fmt"

func main() {

	//! Signed Integers (whole numbers with +/-)
	var a int = 10               //? Platform dependent (32 or 64-bit)
	var b int8 = 127             //? Range: -128 to 127
	var c int16 = 32000          //? Range: -32768 to 32767
	var d int32 = 200000         //? Range: -2.1B to 2.1B
	var e int64 = 9000000000     //? Range: -9.2Q to 9.2Q

	//! Unsigned Integers (only positive)
	var u uint = 100             //? Platform dependent
	var u8 uint8 = 255           //? Range: 0 to 255

	//! Floating Point (decimals)
	var f1 float32 = 3.14        //? ~6-7 digits precision
	var f2 float64 = 99.999999   //? ~15-16 digits precision

	//! Complex Numbers (real + imaginary)
	var comp complex64 = 2 + 3i  //? float32 real + imag
	var comp2 complex128 = 5 + 7i //? float64 real + imag

	//! Boolean (true/false)
	var isGoFun bool = true      //? Only true or false

	//! String (text)
	var language string = "Go Programming"  //? UTF-8 encoded text

	//! Array (fixed size)
	var numbers [3]int = [3]int{1, 2, 3}  //? Length is part of type

	//! Slice (dynamic array)
	var fruits []string = []string{"Apple", "Banana", "Mango"}  //? Can grow/shrink

	//! Map (key-value pairs)
	var student map[string]int = map[string]int{  //? Like dictionary/hash
		"Love": 90,
		"Raj":  85,
	}

	//! Struct (custom type grouping)
	type Person struct {
		name string  //? Field 1
		age  int     //? Field 2
	}
	p1 := Person{name: "Love", age: 21}

	//! Pointer (memory address)
	var x int = 50
	var ptr *int = &x   //? & = address, * = value at address

	//// Output Section
	fmt.Println("int:", a)
	fmt.Println("int8:", b)
	fmt.Println("int16:", c)
	fmt.Println("int32:", d)
	fmt.Println("int64:", e)

	fmt.Println("uint:", u)
	fmt.Println("uint8:", u8)

	fmt.Println("float32:", f1)
	fmt.Println("float64:", f2)

	fmt.Println("complex64:", comp)
	fmt.Println("complex128:", comp2)

	fmt.Println("bool:", isGoFun)
	fmt.Println("string:", language)
	fmt.Println("array:", numbers)
	fmt.Println("slice:", fruits)
	fmt.Println("map:", student)
	fmt.Println("struct:", p1)
	fmt.Println("pointer value:", *ptr)
	fmt.Println("pointer address:", ptr)
}