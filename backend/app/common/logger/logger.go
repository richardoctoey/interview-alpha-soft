package logger

import "log"

func Error(err error, obj interface{}) {
	log.Printf("Error: %s || obj %v", err.Error(), obj)
}

func Info(message string, obj interface{}) {
	log.Printf("Error: %s || obj %v", message, obj)
}
