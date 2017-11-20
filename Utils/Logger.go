package Utils

import (
	"log"
	"os"
	"time"
)

const Info = 100
const Fatal = 101

type ILogger interface {
	WriteLog(str string, t int)
}

type Logger struct {
}

func (l *Logger) WriteLog(str string, t int) {
	settings := GetSettings()
	logDate, fileDate := setDates()

	dir := settings.Log.Folder
	file := dir + "/" + settings.Log.File + fileDate + settings.Log.Ext

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModeDir)
	}

	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Println(err)
	}

	f.WriteString(logDate + ": " + str + "\n")
	defer f.Close()

	switch t {
	case Info:
		log.Println(str)
	case Fatal:
		log.Fatal(str)
	default:
		log.Println(str)
	}
}

func setDates() (string, string) {
	now := time.Now()
	return now.Format("01-02-2006 15:04:05"), now.Format("20060102")
}
