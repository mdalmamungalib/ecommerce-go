package main

import (
	// "ecommerce/cmde"
	"fmt"
	"os"
)

type People interface {
	PrintDetails()
	ReceiveMoney(amount float64) float64
	WithdrawMoney(amount float64) float64
}

type BankUser interface {
	WithdrawMoney(amount float64) float64
}

type user struct {
	Name  string
	Age   int
	Money float64
}

func (obj user) PrintDetails() {
	fmt.Println("Name:", obj.Name)
	fmt.Println("Age:", obj.Age)
	fmt.Println("Money:", obj.Money)
}

func (obj user) WithdrawMoney(amount float64) float64 {
	obj.Money = obj.Money - amount
	return obj.Money
}

func (obj user) ReceiveMoney(amount float64) float64 {
	obj.Money = obj.Money - amount
	return obj.Money
}

func main() {

	var usr1 People
	usr1 = user{
		Name:  "John Doe",
		Age:   30,
		Money: 1000.0,
	}

	var usr2 People
	usr2 = user{
		Name:  "Jane Smith",
		Age:   25,
		Money: 1500.0,
	}

	var usr3 BankUser
	usr3 = user{
		Name: "Alice Johnson",
		Age:  28,
		Money: 2000.0,
	}

	usr3.WithdrawMoney(10)
	obj, ok := usr3.(user)
	if !ok {
		fmt.Println("Sorry usr3 is not type of user struct")
		os.Exit(1)
	}
	obj.PrintDetails()
	obj.ReceiveMoney(10)

	var usr4 People
	usr4 = user{
		Name:  "Bob Brown",
		Age:   35,
		Money: 500.0,
	}

	usr1.PrintDetails()
	usr2.PrintDetails()
	usr1.ReceiveMoney(100)
	usr4.PrintDetails()

	// cmde.Serve()
}
