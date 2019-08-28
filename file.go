package tools

import (
	"os"

	"github.com/pkg/errors"
)

func MakeDirectory(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(path, 0x755)
		}
		return err
	}
	if !fi.IsDir() {
		return errors.New("specified path is not a directory")
	}
	return nil
}

func PathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}
