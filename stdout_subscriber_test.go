package logger

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"os"
	"sync"
	"testing"
)

var FAKE_TIMESTAMP = 1621359006

var debugMsg = NewMessage(FAKE_TIMESTAMP, "bonjour je suis un msg de log", DEBUG)
var infoMsg = NewMessage(FAKE_TIMESTAMP, "bonjour je suis un msg de log", INFO)
var warningMsg = NewMessage(FAKE_TIMESTAMP, "bonjour je suis un msg de log", WARNING)
var errorMsg = NewMessage(FAKE_TIMESTAMP, "bonjour je suis un msg de log", ERROR)
var criticalMsg = NewMessage(FAKE_TIMESTAMP, "bonjour je suis un msg de log", CRITICAL)

func TestMsgAreWellPrintedOnStdout(t *testing.T) {
	sub := NewStdoutSubscriber()
	assert := assert.New(t)

	rs := captureOutput(sub.notify, &debugMsg)
	assert.Equal("[DEBUG]:[2021-05-18 17:30:06]:bonjour je suis un msg de log", rs)

	rs = captureOutput(sub.notify, &infoMsg)
	assert.Equal("[INFO]:[2021-05-18 17:30:06]:bonjour je suis un msg de log", rs)

	rs = captureOutput(sub.notify, &warningMsg)
	assert.Equal("[WARNING]:[2021-05-18 17:30:06]:bonjour je suis un msg de log", rs)

	rs = captureOutput(sub.notify, &errorMsg)
	assert.Equal("[ERROR]:[2021-05-18 17:30:06]:bonjour je suis un msg de log", rs)

	rs = captureOutput(sub.notify, &criticalMsg)
	assert.Equal("[CRITICAL]:[2021-05-18 17:30:06]:bonjour je suis un msg de log", rs)
}

func captureOutput(f func(*Message), m *Message) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f(m)
	writer.Close()
	return <-out
}
