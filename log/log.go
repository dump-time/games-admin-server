package log

import "log"

func Info(v interface{}) {
	log.Println("[\033[1;32mINFO\033[0m] ", v)
}

func Warn(v interface{}) {
	log.Println("[\033[1;33mWARN\033[0m] ", v)
}

func Error(v interface{}) {
	log.Println("[\033[1;31mERROR\033[0m] ", v)
}

func Fatal(v interface{}) {
	log.Fatalln("[\033[1;30;41mFATAL\033[0m] ", v)
}
