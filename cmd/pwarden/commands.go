package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

type command struct {
	callback func()
}

var commandMap = map[string]command{
	"start": {
		callback: cmdStart,
	},
}

func cmdStart() {
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

			log.Fatal(err)
		}

		logMessage(Success, "detected and closed "+comm)
	}
}
