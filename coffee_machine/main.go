package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/everbslab/go_examples/coffee_machine/internals/model"
)

func main() {
	fmt.Println("Hello!")

	machine := model.NewMachine(500, 100)
	machine.PrintState()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Write down required action:")
	fmt.Print("> ")

	for scanner.Scan() {
		text := scanner.Text()

		switch text {
		case "exit":
			os.Exit(0)
		default:
			// @todo: write payload
			break
		}
		fmt.Print("> ")
	}
}
