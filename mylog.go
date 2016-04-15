package mylog

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	// Trace ...
	Trace *log.Logger
	// Info ...
	Info *log.Logger
	// Warning ...
	Warning *log.Logger
	// Error ...
	Error *log.Logger
)

func init() {
	logformat := log.Ldate | log.Ltime | log.Lshortfile
	traceHandle := ioutil.Discard
	infoHandle := ioutil.Discard
	warningHandle := ioutil.Discard
	errorHandle := ioutil.Discard

	switch os.Getenv("MYLOGLEVEL") {
	case "TRACE":
		traceHandle = os.Stdout
		infoHandle = os.Stdout
		warningHandle = os.Stdout
		errorHandle = os.Stdout
		break
	case "INFO":
		infoHandle = os.Stdout
		warningHandle = os.Stdout
		errorHandle = os.Stdout
		break
	case "WARNING":
		warningHandle = os.Stdout
		errorHandle = os.Stdout
		break
	case "ERROR":
		errorHandle = os.Stdout
		break
	}

	Trace = log.New(traceHandle,
		"TRACE: ",
		logformat)

	Info = log.New(infoHandle,
		"INFO: ",
		logformat)

	Warning = log.New(warningHandle,
		"WARNING: ",
		logformat)

	Error = log.New(errorHandle,
		"ERROR: ",
		logformat)
}
