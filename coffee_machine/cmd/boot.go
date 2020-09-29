package cmd

import (
	"bufio"
	"fmt"
	"github.com/everbslab/go_examples/coffee_machine/internals/model"
	"go.uber.org/zap"
	"os"
)

func BootMachine() {
	fmt.Println("Coffee Machine example.")

	logger, _ := zap.NewProduction()
	defer func() {
		_ = logger.Sync()
	}() // flushes buffer, if any

	machine := model.NewMachine(500, 100)
	fmt.Println(machine.PrintState())

	fmt.Println("---")
	fmt.Printf("Available actions: %s.\n", machine.AvailableActions(", "))
	fmt.Println("---")

	fmt.Println("Enter action:")
	fmt.Print("> ")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()

		switch text {
		case model.ExitAction:
			cmd := machine.MakeExit()
			if err := cmd.Execute(); err != nil {
				logger.Fatal(err.Error())
			}
		case model.DepositAction:
			fmt.Print(":: Enter deposit value > ")
			var am float64
			if _, err := fmt.Scanln(&am); err != nil {
				fmt.Println("ERROR: Wrong deposit amount.")
			} else {
				cmd := machine.MakeDeposit(am)
				if err := cmd.Execute(); err != nil {
					logger.Fatal(err.Error())
				}
			}
		case model.BrewAction:
			c := model.NewCoffee(100, 50, 2, 50, 2.50)
			cmd := machine.MakeBrew(c)
			if err := cmd.Execute(); err == nil {
				logger.Info("Brewing coffee succeed.")
			} else {
				logger.Fatal("ERROR: " + err.Error())
			}
		case model.StatusAction:
			cmd := machine.MakeStatus()
			_ = cmd.Execute()
		default:
			fmt.Printf("Unknown action. Available actions: %v.\n", machine.AvailableActions(", "))
		}

		fmt.Print("> ")
	}
}
