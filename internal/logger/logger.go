package logger

import (
	"log"
	"os"
)

func New(name string) *log.Logger {
	return log.New(os.Stdout, name+":", log.Ldate|log.Ltime)
}
