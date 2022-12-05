package util

import (
	"path/filepath"
	"runtime"
	"strconv"
)

func Cwd() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("Failed to get caller")
	}

	return filepath.Dir(filename)
}

func MustInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return int(i)
}
