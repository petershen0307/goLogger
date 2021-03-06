package logServer

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// LogWriter is a structure to handle with file
type LogWriter struct {
	logFileName string
	logFileDir  string
	writeBuffer bytes.Buffer
}

// SetPath will set dir and file name
func (w *LogWriter) SetPath(dir, fileName string) {
	w.logFileDir, w.logFileName = dir, fileName
}

// BufferSize will get current buffer size
func (w *LogWriter) BufferSize() int {
	return w.writeBuffer.Len()
}

// Flush will flush buffer to file and return file size
func (w *LogWriter) Flush() int64 {
	logFilePath := filepath.Join(w.logFileDir, w.logFileName)
	file, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("Open file:", logFilePath, " error:", err)
		return -1
	}
	defer file.Close()
	_, err = w.writeBuffer.WriteTo(file)
	if err != nil {
		fmt.Println("buffer WriteTo() error:", err)
		return -1
	}
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("get file info error:", err)
		return -1
	}
	return fileInfo.Size()
}

// PushToBuffer will push string to buffer
func (w *LogWriter) PushToBuffer(s string) {
	_, err := w.writeBuffer.WriteString(s)
	if err != nil {
		fmt.Println("buffer WriteString() error:", err)
	}
}

// SplitFile will rename current log file with timestamp
func (w *LogWriter) SplitFile() {
	logFilePath := filepath.Join(w.logFileDir, w.logFileName)
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		return
	}
	t := time.Now().Format(time.RFC3339)
	timestamp := strings.Replace(t, ":", "-", -1)
	newPath := logFilePath + timestamp
	err = os.Rename(logFilePath, newPath)
	if err != nil {
		fmt.Println("os Rename() error:", err)
		return
	}
}
