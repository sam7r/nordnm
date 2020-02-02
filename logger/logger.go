package logger

import (
	"github.com/google/logger"
	"io/ioutil"
)

// STDout writes log to stdout
var STDout *logger.Logger

// init creates and binds instances of InfoLogger and SystemLogger types
func init() {
	STDout = logger.Init("STDoutLogger", true, false, ioutil.Discard)
	defer STDout.Close()
}
