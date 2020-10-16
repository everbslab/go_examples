package model

import (
	"errors"
	"fmt"
	"os"
)

const (
	exitCmd    string = "exit"
	depositCmd string = "deposit"
	brewCmd    string = "brew"
	statusCmd  string = "status"
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
	c.M.ing.milk -= c.Drink.milk
	if c.M.ing.milk < 0 {
		return errors.New("insufficient milk")
	}

	c.M.ing.sugar -= c.Drink.sugar
	if c.M.ing.sugar < 0 {
		return errors.New("insufficient sugar")
	}

	c.M.ing.bean -= c.Drink.beans
	if c.M.ing.bean < 0 {
		return errors.New("insufficient beans")
	}

	c.M.Deposit += c.Drink.price

	return nil
}

func (c *StatusCommand) Execute() error {
	fmt.Println(c.M.PrintState())
	return nil
}
