package main

import (
	"bufio"
	"fmt"

	npipe "gopkg.in/natefinch/npipe.v2"
)

func main() {
	ln, err := npipe.Listen(`\\.\pipe\mypipename`)
	if err != nil {
		fmt.Println("open pipe error")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}

		msg, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print(msg)
	}
}
