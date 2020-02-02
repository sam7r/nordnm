package logger

import (
	"github.com/google/logger"
	"io/ioutil"
)

// Stdout writes log to stdout
var Stdout *logger.Logger

// Init creates and binds instances of InfoLogger and SystemLogger types
func Init(verbose bool) {
	Stdout = logger.Init("STDoutLogger", verbose, false, ioutil.Discard)
	defer Stdout.Close()
}
