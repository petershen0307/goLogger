package logServer

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"

	winio "github.com/Microsoft/go-winio"
)

type command int

const (
	cmdFlush command = iota
	cmdEnqueue
	cmdSplit
	cmdExit
)

type jobStruct struct {
	jobCommand command
	message    string
}

// LogServer is the structure collecting log from pipe
type LogServer struct {
	pipeName  string
	jobQueue  chan jobStruct
	wg        sync.WaitGroup
	exitEvent chan struct{}
}

// Start is the LogServer entry function
func (server *LogServer) Start(pipeName string) {
	if pipeName == "" {
		fmt.Println("empty pipe name")
		return
	}
	server.pipeName = pipeName
	// 1. pre work
	workFuncs := []func(){server.receiver, server.worker}
	server.wg.Add(len(workFuncs))
	server.jobQueue = make(chan jobStruct, 300)
	server.exitEvent = make(chan struct{})

	osEvent := make(chan os.Signal, 1)
	signal.Notify(osEvent, os.Interrupt)
	go func() {
		// wait interrupt signal
		<-osEvent
		// sig is a ^C, handle it
		close(server.exitEvent)
	}()
	// 2. create receiver, command worker
	for _, wFunc := range workFuncs {
		go wFunc()
	}
	server.wg.Wait()
}

func (server *LogServer) receiver() {
	defer server.wg.Done()
	listener, err := winio.ListenPipe(server.pipeName, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// handle pipe message
	go func(l net.Listener) {
		for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Println("exit with error:", err)
				return
			}
			bufioReader := bufio.NewReader(conn)
			msg, _ := bufioReader.ReadString('\n')
			job := jobStruct{jobCommand: cmdEnqueue, message: msg}
			server.jobQueue <- job
			conn.Close()
		}
	}(listener)

	// wait exit event
	<-server.exitEvent
	listener.Close()
}

func (server *LogServer) worker() {
	defer server.wg.Done()
	bExit := false
	for {
		if bExit && 0 == len(server.jobQueue) {
			return
		}
		// listen exit command and pipe listener
		select {
		case <-server.exitEvent:
			bExit = true
		case job := <-server.jobQueue:
			switch job.jobCommand {
			case cmdFlush:
			case cmdEnqueue:
				fmt.Print(job.message)
			case cmdSplit:
			}
		}
	}
}
