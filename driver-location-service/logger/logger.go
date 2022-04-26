package logger


import (
	"log"
	"os"
)

var logg *log.Logger

func Init() {
	/*
	file ,err := os.OpenFile("Matcing-service.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Faile to open log file", err)
	}
	*/

	logg = log.New(os.Stdout, "Driver-Service LOG ", log.LstdFlags)
}

func Info(msg string) {
	logg.Println("INFO: ", msg)
}

func Infof(format string, v ...interface{}) {
	logg.Printf("INFO: " + format, v...)
}

func Error(msg string) {
	logg.Println("ERROR: ", msg)
}

func Fatal(v ...interface{}) {
	logg.Fatal(v...)
}
