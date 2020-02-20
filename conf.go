package logger

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// logger definition. Their behavior is defined by two env variable. $SERVICE_NAME, which can be used to customize
// the output and $LOG_LEVEL which may have five value : empty "debug", "info", "warn" and "error".
// If $LOG_LEVEL is not set, the default value will be "info". The critical level may not be disabled.
var (
	Debug    *log.Logger = log.New(debugWriter(), fmt.Sprintf("%s - DEBUG : ", os.Getenv("SERVICE_NAME")), log.Ldate|log.Ltime|log.Lshortfile)
	Info     *log.Logger = log.New(infoWriter(), fmt.Sprintf("%s - INFO : ", os.Getenv("SERVICE_NAME")), log.Ldate|log.Ltime|log.Lshortfile)
	Warn     *log.Logger = log.New(warnWriter(), fmt.Sprintf("%s - WARN : ", os.Getenv("SERVICE_NAME")), log.Ldate|log.Ltime|log.Lshortfile)
	Error    *log.Logger = log.New(errorWriter(), fmt.Sprintf("%s - ERROR : ", os.Getenv("SERVICE_NAME")), log.Ldate|log.Ltime|log.Llongfile)
	Critical *log.Logger = log.New(os.Stderr, fmt.Sprintf("%s - CRITICAL : ", os.Getenv("SERVICE_NAME")), log.Ldate|log.Ltime|log.Llongfile)
)

func debugWriter() io.Writer {
	if debugEnable() {
		return os.Stdout
	}
	return ioutil.Discard
}

func infoWriter() io.Writer {
	if infoEnable() {
		return os.Stdout
	}
	return ioutil.Discard
}

func warnWriter() io.Writer {
	if warnEnable() {
		return os.Stdout
	}
	return ioutil.Discard
}

func errorWriter() io.Writer {
	if errorEnable() {
		return os.Stdout
	}
	return ioutil.Discard
}

func debugEnable() bool {
	return os.Getenv("LOG_LEVEL") == "debug"
}

func infoEnable() bool {
	return debugEnable() || os.Getenv("LOG_LEVEL") == "" || os.Getenv("LOG_LEVEL") == "info"
}

func warnEnable() bool {
	return infoEnable() || os.Getenv("LOG_LEVEL") == "warn"
}

func errorEnable() bool {
	return warnEnable() || os.Getenv("LOG_LEVEL") == "error"
}
