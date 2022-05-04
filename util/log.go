package util

import "log"

func InfoLog(logString string) {
	log.Printf("[INFO] %v\n", logString)
}

func ErrorLog(logString string) {
	log.Printf("[Error] %v\n", logString)
}
