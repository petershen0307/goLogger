package logClient

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func getTimeStr() string {
	t := time.Now()
	return t.Format(time.RFC3339)
}

func getFileNameLineStr() string {
	// the grandparent is the actual caller we interest
	pc, filePath, line, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	_, fileName := path.Split(filePath)
	return fmt.Sprintf("(%v, %v, %v)", fileName, funcName, line)
}

func getPIDStr() string {
	return fmt.Sprintf("{%v}", os.Getpid())
}

func getProcessName() string {
	_, fileName := filepath.Split(os.Args[0])
	return fmt.Sprintf("<%v>", fileName)
}

// Log is to generate log string
func Log(level LogLevel, format string, parameters ...interface{}) {
	fmtTmp := fmt.Sprintf(format, parameters...)
	logStr := strings.Join([]string{getTimeStr(), getProcessName(), getPIDStr(), level.toString(), fmtTmp, getFileNameLineStr()}, " ")
	logStr += "\n"
	switch ModeSetting {
	case ModePipe:
	case ModePrint:
		fmt.Print(logStr)
	}
}
