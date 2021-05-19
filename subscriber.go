package logger

import (
	"fmt"
	"time"
)

type Subscriber interface {
	notify(msg *Message)
}

type Message struct {
	timestamp int
	msg string
	level Level
}


func NewMessage(timestamp int, msg string, level Level) Message {
	return Message{
		timestamp,
		msg,
		level,
	}
}


func NewStdoutSubscriber() Subscriber {
	return &stdoutSubscriber{}
}

// subscribers that parses and print the logs to stdout.
type stdoutSubscriber struct {}

func (s *stdoutSubscriber) notify(msg *Message){
	fmt.Print(messageToString(msg))
}

func messageToString(msg *Message) string {
	var timestamp string
	var levelStr string

	if msg.level == DEBUG {
		levelStr = "DEBUG"
	} else if msg.level == INFO {
		levelStr = "INFO"
	} else if msg.level == WARNING {
		levelStr = "WARNING"
	} else if msg.level == ERROR {
		levelStr = "ERROR"
	} else {
		levelStr = "CRITICAL"
	}

	timestamp = time.Unix(int64(msg.timestamp), 0).UTC().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("[%s]:[%s]:%s", levelStr, timestamp, msg.msg)
}
