package logger


import (
	"log"
	"os"
)

type builtInLogger struct {
	l	*log.logger,
}


func NewBuiltInLogger(out, io.Writer, name string) *Logger {
	newLog := l.New(out, name, log.LstdFlags)
	return &builtInLogger{
		l:	newLog,
	}
}

func (ll *builInLogger) Println(v ...any) {
	ll.l.Println("LOG: ", v)
}

func (ll *builtInLogger) Errorf(v ...any) {
	ll.l.Errorf("ERROR: %v\n",v)
}
