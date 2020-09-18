package model

import (
	"fmt"
	"github.com/everbslab/go_examples/coffee_machine/internals/action"
	"strings"
)

// Machine is the struct that represents Coffee machine domain
type Machine struct {
	Deposit int
	Milk    int
	actions []string
}

// NewMachine creates new instance of Machine
func NewMachine(Deposit int, Milk int) *Machine {
	return &Machine{
		Deposit: Deposit,
		Milk:    Milk,
		actions: []string{action.DepositAction, action.ExitAction},
	}
}

// PrintState exports formatted string of Machine state
func (c *Machine) PrintState() string {
	return fmt.Sprintf(`Coffee Machine has:
%d ml of milk
$%d of money`, c.Milk, c.Deposit)
}

// AvailableActions outputs available actions for certain Machine
func (c *Machine) AvailableActions(sep string) string {
	return strings.Join(c.actions, sep)
}
