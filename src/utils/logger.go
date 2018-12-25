package utils

import (
	"github.com/gin-gonic/gin"
	"log"
	"sync"
)

var logger *log.Logger
var once sync.Once

func Error(str string) {

}

func LogInstance() *log.Logger {
	once.Do(func() {
		if logger == nil {
			logger = log.New(gin.DefaultWriter, "", log.Ldate|log.Ltime|log.Llongfile)
		}
	})
	return logger
}
