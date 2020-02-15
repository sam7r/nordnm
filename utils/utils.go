package utils

import (
	"io/ioutil"
	"os"
)

// SaveFile saves file at given path
func SaveFile(filepath string, file []byte, permissions int) error {
	return ioutil.WriteFile(filepath, file, os.FileMode(permissions))
}
