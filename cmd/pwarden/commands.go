package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type command struct {
	reqArgs  []int
	callback func([]string) error
}

var commandMap = map[string]command{
	"start": {
		reqArgs:  []int{0},
		callback: cmdStart,
	},
	"search": {
		reqArgs:  []int{1},
		callback: cmdSearch,
	},
}

func cmdStart(args []string) error {
	stdout := strings.Builder{}
	stderr := strings.Builder{}

	var killList = []string{"Discord"}

	for _, comm := range killList {
		cmd := exec.Command("killall", "-s", "SIGINT", comm)
		cmd.Stdin = os.Stdin
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		err := cmd.Run()

		if err != nil {
			if strings.Contains(stderr.String(), "no process found") {
				continue
			}
			return err
		}

		logMessage(Success, "detected and closed "+comm)
	}

	return nil
}

func cmdSearch(args []string) error {
	stdout := strings.Builder{}

	cmd := exec.Command("ps", "-e", "-o", "pid,comm,cmd")
	cmd.Stdin = os.Stdin
	cmd.Stdout = &stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	for i, pLine := range strings.Split(stdout.String(), "\n") {
		if i == 0 || pLine == "" {
			continue
		}

		pid := strings.Fields(pLine)[0]
		comm := strings.Fields(pLine)[1]
		cmd := strings.Fields(pLine)[2]

		if strings.Contains(strings.ToLower(comm), strings.ToLower(args[2])) {
			fmt.Printf("PID: %v, COMM: %v\n", pid, comm)
			fmt.Printf("CMD: %v\n", cmd)
			fmt.Println()
		}
	}

	return nil
}
