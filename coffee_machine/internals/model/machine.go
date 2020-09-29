package model

import (
	"fmt"
	"strings"
)

// Machine is the struct that represents Coffee machine domain
type Machine struct {
	Deposit float64
	Milk    uint
	actions []string
}

// NewMachine creates new instance of Machine
func NewMachine(Deposit float64, Milk uint) *Machine {
	return &Machine{
		Deposit: Deposit,
		Milk:    Milk,
		actions: []string{DepositAction, ExitAction, BrewAction, StatusAction},
	}
}

// PrintState exports formatted string of Machine state
func (c *Machine) PrintState() string {
	return fmt.Sprintf(`Coffee Machine has:
%d ml of milk
$%f of money`, c.Milk, c.Deposit)
}

// AvailableActions outputs available actions for certain Machine
func (c *Machine) AvailableActions(sep string) string {
	return strings.Join(c.actions, sep)
}

// MakeExit closes the Machine
func (c *Machine) MakeExit() Command {
	return &ExitCommand{
		M: c,
	}
}

// MakeDeposit updates the state with deposit amount
func (c *Machine) MakeDeposit(amount float64) Command {
	return &DepositCommand{
		Amount: amount,
		M:      c,
	}
}

func (c *Machine) MakeBrew(coffee *Coffee) Command {
	return &BrewCommand{
		Drink: coffee,
		M:     c,
	}
}

func (c *Machine) MakeStatus() Command {
	return &StatusCommand{
		M: c,
	}
}
