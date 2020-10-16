package cmd

import (
	"bufio"
	"fmt"
	machine "github.com/everbslab/go_examples/coffee_machine/internals/model"
	"go.uber.org/zap"
	"os"
)

func BootMachine() {
	fmt.Println("Coffee Machine example.")

	logger, _ := zap.NewProduction()
	defer func() {
		_ = logger.Sync()
	}() // flushes buffer, if any

	cm := machine.New(500, nil)
	fmt.Println(cm.PrintState())

	fmt.Println("---")
	fmt.Printf("Available actions: %s.\n", cm.AvailableActions(", "))
	fmt.Println("---")

	fmt.Println("Enter action:")
	fmt.Print("> ")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()

		if err := cm.ExecCommand(text); err != nil {
			logger.Fatal(err.Error())
		}

		//switch text {
		//case machine.exitCmd:
		//	err := cm.MakeExit().Execute()
		//	failOnError(err)
		//case machine.depositCmd:
		//	cmd := cm.MakeDeposit()
		//	fmt.Print(":: Enter deposit value > ")
		//	var am float64
		//	if _, err := fmt.Scanln(&am); err != nil {
		//		fmt.Println("ERROR: Wrong deposit amount.")
		//	} else {
		//		err := cmd.Execute()
		//		failOnError(err)
		//	}
		//case machine.brewCmd:
		//	c := machine.NewCoffee(100, 50, 2, 50, 2.50)
		//	cmd := cm.MakeBrew(c)
		//	if err := cmd.Execute(); err == nil {
		//		logger.Info("Brewing coffee succeed.")
		//	} else {
		//		logger.Fatal("ERROR: " + err.Error())
		//	}
		//case machine.statusCmd:
		//	err := cm.MakeStatus().Execute()
		//	failOnError(err)
		//default:
		//	fmt.Printf("Unknown action. Available actions: %v.\n", cm.AvailableActions(", "))
		//}

		fmt.Print("> ")
	}
}
