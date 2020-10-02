package model

import (
	"fmt"
	"strings"
)

type Ingredient int

type Ingredients struct {
	milk  Ingredient
	bean  Ingredient
	sugar Ingredient
	cup   Ingredient
}

// Machine is the struct that represents Coffee machine domain
type Machine struct {
	Deposit float64
	ing     *Ingredients
	actions []Action
}

// New creates new instance of Machine
func New(deposit float64, is *Ingredients) *Machine {
	if is == nil {
		is = &Ingredients{}
	}

	return &Machine{
		Deposit: deposit,
		ing:     is,
		actions: []Action{DepositAction, ExitAction, BrewAction, StatusAction},
	}
}

// PrintState exports formatted string of Machine state
func (c *Machine) PrintState() string {
	return fmt.Sprintf(`Coffee Machine has:
%d ml of milk
$%.2f of money`, c.ing.milk, c.Deposit)
}

// AvailableActions outputs available actions for certain Machine
func (c *Machine) AvailableActions(sep string) string {
	return strings.Join(func(as []Action) []string {
		var ss []string
		for _, v := range as {
			ss = append(ss, string(v))
		}
		return ss
	}(c.actions), sep)
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
