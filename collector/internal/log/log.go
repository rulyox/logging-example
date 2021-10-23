package log

import (
	"os"
	"path"
	"time"

	"collector/internal/utility"
)

var logDir = "logs"
var logFile = "collector.log"

func InitLogPath() {
	currentDir := utility.CurrentDir()

	// create logs path if not exists
	err := os.MkdirAll(path.Join(currentDir, logDir), os.ModePerm)
	if err != nil {
		panic("Log path creation failed")
	}
}

func Log(text string) {
	currentDir := utility.CurrentDir()

	// create file if not exist and append
	f, _ := os.OpenFile(path.Join(currentDir, logDir, logFile), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	datetime := time.Now()
	formattedDatetime := datetime.Format("2006-01-02 15:04:05 ")

	f.WriteString(formattedDatetime + text + "\n")
}
