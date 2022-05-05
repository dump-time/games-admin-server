package log

import "log"

func Info(v any) {
	log.Println("[\033[1;32mINFO\033[0m] ", v)
}

func Warn(v any) {
	log.Println("[\033[1;33mWARN\033[0m] ", v)
}

func Error(v any) {
	log.Println("[\033[1;31mERROR\033[0m] ", v)
}

func Fatal(v any) {
	log.Fatalln("[\033[1;30;41mFATAL\033[0m] ", v)
}
