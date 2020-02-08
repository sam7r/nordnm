package utils

import (
	"io/ioutil"
	"os"
)

// SaveFileToTempLocation saves file at given path
func SaveFileToTempLocation(filepath string, file []byte, permissions int) error {
	return ioutil.WriteFile(filepath, file, os.FileMode(permissions))
}
