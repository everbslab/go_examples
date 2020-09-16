package model

import "fmt"

type Machine struct {
	Deposit int
	Milk    int
}

func NewMachine(Deposit int, Milk int) *Machine {
	return &Machine{
		Deposit: Deposit,
		Milk:    Milk,
	}
}

func (c *Machine) PrintState() {
	fmt.Printf("Coffee Machine has: \n%d of milk \n$%d of money\n", c.Milk, c.Deposit)
}
