package utils

import (
	"io"
	"log"
	"os"
)

const (
	Reset      = "\033[0m"
	Red        = "\033[31m"
	Green      = "\033[32m"
	Yellow     = "\033[33m"
	Blue       = "\033[34m"
	Cyan       = "\033[36m"
	CustomGray = "\033[38;5;248m"
)
const (
	INFO  = Green
	WARN  = Yellow
	DEBUG = Blue
	ERROR = Red
)

type Logger struct {
	info  *log.Logger // INFO logger
	warn  *log.Logger // WARN logger
	debug *log.Logger // DEBUG logger
	error *log.Logger // ERROR logger
	file  *os.File    // Logger file
}

func (logger Logger) Info(msg string) {
	logger.info.Output(2, Green+msg+Reset)
}

func (logger Logger) Warn(msg string) {
	logger.warn.Output(2, Yellow+msg+Reset)
}
func (logger Logger) Debug(msg string) {
	logger.debug.Output(2, Blue+msg+Reset)
}

func (logger Logger) Error(msg string) {
	logger.error.Output(2, Red+msg+Reset)
}

func (logger Logger) Close() error {
	return logger.file.Close()
}

func NewLogger() (*Logger, error) {
	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		return nil, err
	}

	multiWriter := io.MultiWriter(os.Stdout, file)
	flags := log.Ldate | log.Ltime | log.Lshortfile

	return &Logger{
		info:  log.New(multiWriter, Green+"INFO: "+CustomGray, flags),
		debug: log.New(multiWriter, Blue+"DEBUG: "+CustomGray, flags),
		warn:  log.New(multiWriter, Yellow+"WARN: "+CustomGray, flags),
		error: log.New(multiWriter, ERROR+"ERROR: "+CustomGray, flags),
	}, nil
}
