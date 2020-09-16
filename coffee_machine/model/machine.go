package model

import "fmt"

// Machine is the struct that represent Coffee machine domain
type Machine struct {
	Deposit int
	Milk    int
}

// Creates Machine instance
func NewMachine(Deposit int, Milk int) *Machine {
	return &Machine{
		Deposit: Deposit,
		Milk:    Milk,
	}
}

// Print state of Machine
func (c *Machine) PrintState() {
	fmt.Printf("Coffee Machine has: \n%d of milk \n$%d of money\n", c.Milk, c.Deposit)
}
