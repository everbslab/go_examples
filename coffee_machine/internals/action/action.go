package action

import (
	"fmt"
	"os"
	"strings"
)

const (
	ExitAction    = "exit"
	DepositAction = "deposit"
)

// Exit used to stop the program
func Exit(ec int) {
	fmt.Printf("[Action %s] chosen.\n", strings.ToUpper(ExitAction))
	os.Exit(ec)
}
