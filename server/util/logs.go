package util

import (
	"github.com/gookit/color"
)

const (
	LogInfo = iota
	LogWarning
	LogError
)

type LogMessage struct {
	LogType  int
	Message  string
	Args     []interface{}
}

var LogChannel chan LogMessage

func init() {
	LogChannel = make(chan LogMessage)
}

func LogHandler() {
	for msg := range LogChannel {
		switch msg.LogType {
			case LogInfo:
				color.LightBlue.Print("\r[INFO]: ")
			case LogWarning:
				color.Yellow.Print("\r[WARNING]: ")
			case LogError:
				color.Red.Print("\r[ERROR]: ")
		}

		color.White.Printf(msg.Message, msg.Args...)
		color.White.Print(" (press enter to continue)")
	}
}

func Info(message string, args ...interface{}) {
	LogChannel <- LogMessage{
		LogType: LogInfo,
		Message: message,
		Args:    args,
	}
}

func Warning(message string, args ...interface{}) {
	LogChannel <- LogMessage{
		LogType: LogWarning,
		Message: message,
		Args:    args,
	}
}

func Error(message string, args ...interface{}) {
	LogChannel <- LogMessage{
		LogType: LogError,
		Message: message,
		Args:    args,
	}
}