package log

import (
	"fmt"
)

type LogMes string
type Logger interface {
	INFO(args ...interface{})
	DEBUG(args ...interface{})
	ERROR(args ...interface{})
	WARN(args ...interface{})
}

func (LogInfo LogMes) INFO() {
	s := "INFO: " + LogInfo + "\r\n"
	//log_s := +s
	fmt.Println(s)
	Savefile("INFO", string(s))
}
func (LogInfo LogMes) DEBUG() {
	s := "DEBUG: " + LogInfo + "\r\n"
	//log_s := +s
	fmt.Println(s)
	Savefile("DEBUG", string(s))
}
func (LogInfo LogMes) ERROR() {
	s := "ERROR: " + LogInfo + "\r\n"
	//log_s := +s
	fmt.Println(s)
	Savefile("ERROR", string(s))
}
func (LogInfo LogMes) WARN() {
	s := "WARN: " + LogInfo + "\r\n"
	//log_s := +s
	fmt.Println(s)
	Savefile("WARN", string(s))
}
