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
	type process struct {
		pid  string
		comm string
		cmd  string
	}

	cmd := exec.Command("ps", "-e", "-o", "pid,comm,cmd")
	stdout := strings.Builder{}
	cmd.Stdin = os.Stdin
	cmd.Stdout = &stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	processMatches := []process{}
	for i, pLine := range strings.Split(stdout.String(), "\n") {
		if i == 0 || pLine == "" {
			continue
		}

		pid := strings.Fields(pLine)[0]
		comm := strings.Fields(pLine)[1]
		cmd := strings.Join(strings.Fields(pLine)[2:], " ")

		commLower := strings.ToLower(comm)
		cmdLower := strings.ToLower(comm)
		query := strings.ToLower(args[2])

		if strings.Contains(commLower, query) || strings.Contains(cmdLower, query) {
			processMatches = append(processMatches, process{
				pid:  pid,
				comm: comm,
				cmd:  cmd,
			})
		}
	}

	for i, p := range processMatches {
		fmt.Printf("%2v: PID: %v, COMM: %v\n", i, p.pid, p.comm)
	}

	return nil
}
