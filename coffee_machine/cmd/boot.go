package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/everbslab/go_examples/coffee_machine/internals/action"
	"github.com/everbslab/go_examples/coffee_machine/internals/model"
)

func BootMachine() {
	fmt.Println("Coffee Machine example.")

	machine := model.NewMachine(500, 100)
	fmt.Println(machine.PrintState())

	fmt.Println("---")
	fmt.Printf("Available actions: %s.\n", machine.AvailableActions(", "))
	fmt.Println("---")

	fmt.Println("Write down required action:")
	fmt.Print("> ")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()

		switch text {
		case action.ExitAction:
			action.Exit(0)
		case action.DepositAction:
			// @todo: write payload
		default:
			fmt.Printf("Unknown action. Available actions: %v.\n", machine.AvailableActions(", "))
			break
		}
		fmt.Print("> ")
	}
}
