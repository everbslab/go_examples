package model

import (
	"fmt"
	"os"
)

const (
	ExitAction    = "exit"
	DepositAction = "deposit"
	BrewAction    = "brew"
	StatusAction  = "status"
)

type Command interface {
	Execute() error
}

type StatusCommand struct {
	M *Machine
}

type BrewCommand struct {
	Drink *Coffee
	M     *Machine
}

type DepositCommand struct {
	Amount float64
	M      *Machine
}

type ExitCommand struct {
	M *Machine
}

func (c *ExitCommand) Execute() error {
	os.Exit(0)
	return nil
}

func (c *DepositCommand) Execute() error {
	c.M.Deposit += c.Amount
	fmt.Printf("New deposit amount: %v\n", c.Amount)
	fmt.Println(c.M.PrintState())
	return nil
}

func (c *BrewCommand) Execute() error {
	c.M.Milk -= c.Drink.Milk
	c.M.Deposit += c.Drink.Price

	return nil
}

func (c *StatusCommand) Execute() error {
	fmt.Println(c.M.PrintState())
	return nil
}
