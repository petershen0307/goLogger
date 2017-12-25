package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"

	winio "github.com/Microsoft/go-winio"
)

type command int

const (
	cmdFlush command = iota
	cmdMessage
	cmdSplit
	cmdExit
)

type jobStruct struct {
	jobCommand command
	message    string
}

func openPipeServer() (listener net.Listener, err error) {
	listener, err = winio.ListenPipe(`\\.\pipe\mypipename`, nil)
	return
}

func pipeReceiver(jobQueue chan jobStruct, listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			jobQueue <- jobStruct{jobCommand: cmdExit, message: ""}
			return
		}
		bufioReader := bufio.NewReader(conn)
		msg, _ := bufioReader.ReadString('\n')
		job := jobStruct{jobCommand: cmdMessage, message: msg}
		jobQueue <- job
		conn.Close()
	}
}

func worker(jobQueue chan jobStruct) {
	for {
		job := <-jobQueue
		switch job.jobCommand {
		case cmdFlush:
		case cmdMessage:
			fmt.Print(job.message)
		case cmdSplit:
		case cmdExit:
			if 0 != len(jobQueue) {
				continue
			}
			break
		}
	}
}

func main() {
	jobQueue := make(chan jobStruct, 300)
	listener, err := openPipeServer()
	if err != nil {
		fmt.Println(err)
	}
	go pipeReceiver(jobQueue, listener)
	go worker(jobQueue)
	osEvent := make(chan os.Signal, 1)
	signal.Notify(osEvent, os.Interrupt)
	for sig := range osEvent {
		// sig is a ^C, handle it
		fmt.Println(sig)
		break
	}
}
