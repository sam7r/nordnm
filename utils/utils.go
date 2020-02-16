package utils

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

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
