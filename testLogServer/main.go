package main

import (
	"encoding/json"
	"fmt"

	logServer "github.com/petershen0307/goLogger/logServer"
)

func main() {
	//testServer()
	testConfiguration()
}

func testServer() {
	s := logServer.LogServer{}
	s.Start(`\\.\pipe\mypipename`)
}

func testConfiguration() {
	setting := logServer.ServerSetting{}
	setting.Init("")
	b, _ := json.MarshalIndent(setting, "", "\t")
	fmt.Println(string(b))
}
