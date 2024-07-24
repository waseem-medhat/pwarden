package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 {
		logMessage(Error, "not enough arguments")
		os.Exit(1)
	}

	if command, ok := commandMap[os.Args[1]]; ok {
		command.callback()
	} else {
		logMessage(Error, os.Args[1]+" is not a command")
		os.Exit(1)
	}
	// SEARCH HELPER (abstracts `ps` command)
	//
	// stdout := strings.Builder{}
	//
	// cmd := exec.Command("ps", "-e", "-o", "pid,comm,cmd")
	// cmd.Stdin = os.Stdin
	// cmd.Stdout = &stdout
	// cmd.Stderr = os.Stderr
	//
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// for i, pLine := range strings.Split(stdout.String(), "\n") {
	// 	if i == 0 || pLine == "" {
	// 		continue
	// 	}
	//
	// 	pid := strings.Fields(pLine)[0]
	// 	comm := strings.Fields(pLine)[1]
	// 	cmd := strings.Fields(pLine)[2]
	//
	// 	if comm == "Discord" {
	// 		fmt.Printf("PID: %v, COMM: %v\n", pid, comm)
	// 		fmt.Printf("CMD: %v\n", cmd)
	// 		fmt.Println()
	// 	}
	// }
}
