package main

import (
	"fmt"
	"os"
	"os/signal"
	"path"
	"runtime"
	"sync"
	"time"

	myLog "github.com/petershen0307/goLogger/logClient"
)

func main() {
	testLog()
	// testGoRoutine()
	// testExitSignal()
}

func testLog() {
	//myLog.ModeSetting = myLog.ModePrint
	myLog.Log(myLog.LevelDebug, "this is test %d, %s", 123, "hello")
}

func testExitSignal() {
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	s := <-c
	fmt.Println("Got signal:", s)
	s = <-c
	fmt.Println("Got signal:", s)
}

func testGoRoutine() {
	stopchan := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			time.Sleep(time.Millisecond)
			select {
			case <-stopchan:
				fmt.Println("exit 1")
				return
			default:
				fmt.Println("1")
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			time.Sleep(time.Millisecond)
			select {
			case <-stopchan:
				fmt.Println("exit 2")
				return
			default:
				fmt.Println("2")
			}
		}
	}()
	time.Sleep(time.Millisecond * 10)
	close(stopchan)
	wg.Wait()
}

func testCaller() {
	pc, filePath, line, _ := runtime.Caller(1)
	fnCaller := runtime.FuncForPC(pc)
	pcFile, pcLine := fnCaller.FileLine(pc)
	_, fileName := path.Split(filePath)
	fmt.Printf("file name: %v, line: %v, function name: %v, pc file %v, pc line %v\n", fileName, line, fnCaller.Name(), pcFile, pcLine)
}

func testTime() {
	timeObj := time.Now()
	fmt.Println(timeObj.Format(time.RFC3339))
}

func multiInput(format string, other ...interface{}) {
	fmt.Printf(format, other...)
}
