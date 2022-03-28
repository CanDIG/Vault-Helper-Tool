package configSettings

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Does error handling
func ValidateToken() error {
	absPath, _ := filepath.Abs("../cli/configSettings/token.txt")
	_, err := ioutil.ReadFile(absPath) // just pass the file name
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}
	fi, err := os.Stat(absPath)
	if err != nil {
		return err
	}
	size := fi.Size()
	if size == 0 {
		return errors.New("file is empty")
	}
	return nil

}
