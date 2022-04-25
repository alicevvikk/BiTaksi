package logger

import(
	"log"
	"os"
)

var Log *log.Logger

func init() {
	file ,err := os.OpenFile("Matcing-service.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Faile to open log file", err)
	}
	Log = log.New(file, "Matcing-service INFO", log.LstdFlags)

}
