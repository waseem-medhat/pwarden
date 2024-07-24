package main

import (
	"os"
	"slices"
)

func main() {
	if len(os.Args) < 2 {
		logMessage(Error, "no command provided")
		os.Exit(1)
	}

	command, ok := commandMap[os.Args[1]]
	if !ok {
		logMessage(Error, os.Args[1]+" is not a command")
		os.Exit(1)
	}

	nArgs := len(os.Args) - 2
	if !slices.Contains(command.reqArgs, nArgs) {
		logMessage(Error, "wrong number of arguments for "+os.Args[1])
		os.Exit(1)
	}

	command.callback(os.Args)
}
