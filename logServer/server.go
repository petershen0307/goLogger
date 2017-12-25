package main

import (
	"bufio"
	"fmt"

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

func pipeReceiver(c chan jobStruct) {
	ln, err := winio.ListenPipe(`\\.\pipe\mypipename`, nil)
	defer ln.Close()
	if err != nil {
		fmt.Println("open pipe error")
	}
	for {
		conn, err := ln.Accept()
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
	go pipeReceiver(c)
	worker(c)
}
