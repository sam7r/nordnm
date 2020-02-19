package utils

import (
	"bufio"
	"github.com/google/logger"
	"io"
	"io/ioutil"
	"os"
)

// Logger writes log to stdout
var Logger *logger.Logger

// SaveFile saves file at given path
func SaveFile(filepath string, file []byte, permissions int) error {
	return ioutil.WriteFile(filepath, file, os.FileMode(permissions))
}

// GetStdoutText reads out and concats string from io.Reader
func GetStdoutText(r io.Reader) (stdout []string) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if lineOut := scanner.Text(); lineOut != "" {
			stdout = append(stdout, lineOut)
		}
	}
	return stdout
}

// InitLogger creates and binds instances of InfoLogger and SystemLogger types
func InitLogger(verbose bool) {
	Logger = logger.Init("STDoutLogger", verbose, false, ioutil.Discard)
	defer Logger.Close()
}
