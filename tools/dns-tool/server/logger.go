package dnsserver

import (
	"fmt"
	"log"
	"os"
)

var (
	logger = log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	infof  = func(msg string) {
		err := logger.Output(2, fmt.Sprintf("[INFO] %s", msg))
		if err != nil {
			return
		}
	}
	debugf = func(msg string) {
		err := logger.Output(2, fmt.Sprintf("[DEBUG] %s", msg))
		if err != nil {
			return
		}
	}
)
