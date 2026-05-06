package main

import "fmt"

func main() {
	var todos []string
	var choice int

	for {
		fmt.Println("\n1. Add  2. List  3. Quit")
		fmt.Print("> ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var task string
			fmt.Print("Task: ")
			fmt.Scan(&task)
			todos = append(todos, task)

		case 2:
			for i, t := range todos {
				fmt.Printf("%d: %s\n", i+1, t)
			}

		case 3:
			return
		}
	}
}