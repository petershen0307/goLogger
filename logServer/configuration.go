package logServer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// ServerSetting is the setting for log server
type ServerSetting struct {
	SplitSizeMB      int64
	PipeName         string
	FlushFrequencyMS time.Duration
	LogFileName      string
	LogFileDir       string
}

// Init the log server setting
func (s *ServerSetting) Init(settingPath string) {
	cwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	// assign default value
	*s = ServerSetting{
		SplitSizeMB:      10,
		PipeName:         `\\.\pipe\mypipename`,
		FlushFrequencyMS: 1000,
		LogFileName:      "debug.log",
		LogFileDir:       cwd,
	}
	// check file exist or not
	if _, err := os.Stat(settingPath); os.IsNotExist(err) {
		return
	}
	file, _ := os.Open(settingPath)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(s)
	if err != nil {
		fmt.Println("Decode setting error:", err)
	}
}

// MB to byte, MS to nano second
func (s *ServerSetting) normalizeValue() {
	s.FlushFrequencyMS *= time.Millisecond
	s.SplitSizeMB *= 1024 * 1024
}
