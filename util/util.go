package util

import (
	"path/filepath"
	"runtime"
)

func Cwd() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("Failed to get caller")
	}

	return filepath.Dir(filename)
}
