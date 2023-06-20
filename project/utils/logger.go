package utils

import (
	"io"
	"os"
)

type Logger struct {
	File string
	Fw   *os.File
}

func NewLogger(file string) *Logger {
	l := &Logger{
		File: file,
	}

	f, err := os.Create(file)
	if err != nil {
		return nil
	}

	defer f.Close()

	l.Fw = f

	return l
}

func (l *Logger) WriteLog() (wf io.Writer) {
	wf = io.MultiWriter(os.Stdout, l.Fw)

	return
}
