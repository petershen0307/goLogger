package main

import (
	"bufio"
	"fmt"

	winio "github.com/Microsoft/go-winio"
)

func main() {
	ln, err := winio.ListenPipe(`\\.\pipe\mypipename`, nil)
	defer ln.Close()
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
