package log

import (
	"io"
	"log"
	"os"
)

var klog = newLogger()

type logger struct {
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
	fatal *log.Logger
}

type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Fatal(v ...interface{})
}

func newLogger() *logger {
	file, _ := os.OpenFile("logs/log.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)

	mw := io.MultiWriter(os.Stdout, file)

	debugLog := log.New(mw, "< DEBUG^^ > ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(mw, "< INFO<3 > ", log.Ldate|log.Ltime)
	warnLog := log.New(mw, "< WARN(-.-) > ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(mw, "< ERROR:( > ", log.Ldate|log.Ltime|log.Lshortfile)
	fatalLog := log.New(mw, "< FATAL>:(!! > ", log.Ldate|log.Ltime|log.Lshortfile)

	return &logger{
		debug: debugLog,
		info:  infoLog,
		warn:  warnLog,
		error: errorLog,
		fatal: fatalLog,
	}
}

func Debug(v ...interface{}) {
	klog.debug.Println(v...)
}

func Info(v ...interface{}) {
	klog.info.Println(v...)
}

func Warn(v ...interface{}) {
	klog.warn.Println(v...)
}

func Error(v ...interface{}) {
	klog.error.Println(v...)
}

func Fatal(v ...interface{}) {
	klog.fatal.Println(v...)
	os.Exit(1)
}
