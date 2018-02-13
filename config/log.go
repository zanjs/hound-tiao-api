package config

import (
	"os"
	"time"

	"github.com/houndgo/houndgo/logfile"
)

func init() {
	logfile.Mkdir(logfile.LogFIlePath)

	f := NewLogFile()
	defer f.Close()
}

// get a filename based on the date, file logs works that way the most times
// but these are just a sugar.
func todayFilename() string {
	today := time.Now().Format("2006-01-02")
	return logfile.LogFIlePath + "/" + today + ".log"
}

// NewLogFile is
func NewLogFile() *os.File {
	filename := todayFilename()
	// open an output file, this will append to the today's file if server restarted.
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return f
}
