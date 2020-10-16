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
	actions map[string]Command
}

// New creates new instance of Machine
func New(deposit float64, is *Ingredients) *Machine {
	if is == nil {
		is = &Ingredients{}
	}

	return &Machine{
		Deposit: deposit,
		ing:     is,
		actions: map[string]Command{
			depositCmd: &DepositCommand{},
			exitCmd:    &ExitCommand{},
			brewCmd:    &BrewCommand{},
			statusCmd:  &StatusCommand{},
		},
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
	return strings.Join(func(as map[string]Command) []string {
		s := make([]string, 0, len(as))
		for k, _ := range as {
			s = append(s, k)
		}
		return s
	}(c.actions), sep)
}

// MakeExit closes the Machine
func (c *Machine) MakeExit() Command {
	return &ExitCommand{
		M: c,
	}
}

// MakeDeposit updates the state with deposit amount
func (c *Machine) MakeDeposit(am float64) Command {
	return &DepositCommand{
		Amount: am,
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

func (c *Machine) ExecCommand(cmd string) error {
	if command := c.actions[cmd]; command == nil {
		fmt.Println(c.AvailableActions(", "))
	} else {
		err := command.Execute()
		return err
	}

	return nil
}
