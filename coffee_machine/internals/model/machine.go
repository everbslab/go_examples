package model

import "fmt"

// Machine is the struct that represent Coffee machine domain
type Machine struct {
	Deposit int
	Milk    int
}

// NewMachine creates new instance of Machine
func NewMachine(Deposit int, Milk int) *Machine {
	return &Machine{
		Deposit: Deposit,
		Milk:    Milk,
	}
}

// PrintState prints to stdout the Machine state
func (c *Machine) PrintState() string {
	return fmt.Sprintf("Coffee Machine has:\n%d ml of milk\n$%d of money\n", c.Milk, c.Deposit)
}
