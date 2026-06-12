package main

//if struct is data, 
// method is a behavior.
// struct shows info.
// cars also do things. start stop. accelerate.

import "fmt"
type Car struct{
	Brand string
	speed int
}
func (c Car)printCar(){ // give every car a print car method. 

	// c.brand menad this current car brand thing.
	fmt.Println("brand: ", c.Brand)
	fmt.Println("speed : " , c.speed)

}

type BankAccount struct {
	owner string 
	balance int 
}
func (b *BankAccount) Deposit(amount int){
	b.balance += amount

}
func (b BankAccount) ShowBalance() {
    fmt.Println("Balance:", b.balance)
}
func main(){

	car := Car {
		Brand: "Bmw",
		speed: 150,
	}
	account := BankAccount{
    owner: "Devraj",
    balance: 1000,
}

account.Deposit(500)
account.ShowBalance()

	car.printCar()
}
//