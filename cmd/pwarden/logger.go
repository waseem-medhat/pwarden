package main

import (
	"fmt"

	codes "github.com/avearmin/stylecodes"
)

type MessageType int

const (
	Success MessageType = iota
	Info
	Error
)

var colorMap = map[MessageType]string{
	Success: codes.ColorGreen,
	Info:    codes.ColorBlue,
	Error:   codes.ColorRed,
}

func logMessage(msgType MessageType, msg string) {
	fmt.Println(colorMap[msgType] + "pwarden: " + msg + codes.ResetColor)
}
