package main

import logServer "github.com/petershen0307/goLogger/logServer"

func main() {
	s := logServer.LogServer{}
	s.Start(`\\.\pipe\mypipename`)
}
