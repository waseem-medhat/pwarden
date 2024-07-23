package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	codes "github.com/avearmin/stylecodes"
)

type MessageType int

const (
	Success MessageType = iota
	Info
	Error
)

var colorMap = map[MessageType]string{
	Success: codes.Color.Green,
	Info:    codes.Color.Blue,
}

func logMessage(msgType MessageType, msg string) {
	fmt.Println(colorMap[msgType] + "pwarden: " + msg + codes.Color.Reset)
}

func main() {
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
