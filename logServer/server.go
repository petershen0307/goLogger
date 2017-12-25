package main

import (
	"bufio"
	"fmt"
	"net"

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

func pipeReceiver(c chan jobStruct, listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			c <- jobStruct{jobCommand: cmdExit, message: ""}
			return
		}
		bufioReader := bufio.NewReader(conn)
		msg, _ := bufioReader.ReadString('\n')
		job := jobStruct{jobCommand: cmdMessage, message: msg}
		c <- job
		conn.Close()
	}
}

func worker(c chan jobStruct) {
	for {
		job := <-c
		switch job.jobCommand {
		case cmdFlush:
		case cmdMessage:
			fmt.Print(job.message)
		case cmdSplit:
		case cmdExit:
			break
		}
	}
}

func main() {
	c := make(chan jobStruct, 300)
	ln, err := openPipeServer()
	if err != nil {
		fmt.Println(err)
	}
	go pipeReceiver(c, ln)
	worker(c)
}
