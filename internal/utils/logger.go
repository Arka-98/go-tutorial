package utils

import (
	"io"
	"log"
	"os"
)

const (
	reset      = "\033[0m"
	red        = "\033[31m"
	green      = "\033[32m"
	yellow     = "\033[33m"
	blue       = "\033[34m"
	customGray = "\033[38;5;248m"
)
const (
	INFO  = green
	WARN  = yellow
	DEBUG = blue
	ERROR = red
)

type logger struct {
	info  *log.Logger // INFO logger
	warn  *log.Logger // WARN logger
	debug *log.Logger // DEBUG logger
	error *log.Logger // ERROR logger
	file  *os.File    // Logger file
}

func NewLogger() (*logger, error) {
	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		return nil, err
	}

	multiWriter := io.MultiWriter(os.Stdout, file)
	flags := log.Ldate | log.Ltime | log.Lshortfile

	return &logger{
		info:  log.New(multiWriter, INFO+"INFO: "+customGray, flags),
		debug: log.New(multiWriter, DEBUG+"DEBUG: "+customGray, flags),
		warn:  log.New(multiWriter, WARN+"WARN: "+customGray, flags),
		error: log.New(multiWriter, ERROR+"ERROR: "+customGray, flags),
		file: file,
	}, nil
}

func (logger logger) Info(msg string) {
	logger.info.Output(2, green+msg+reset)
}

func (logger logger) Warn(msg string) {
	logger.warn.Output(2, yellow+msg+reset)
}
func (logger logger) Debug(msg string) {
	logger.debug.Output(2, blue+msg+reset)
}

func (logger logger) Error(msg string) {
	logger.error.Output(2, red+msg+reset)
}

func (logger logger) Close() error {
	return logger.file.Close()
}
