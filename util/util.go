package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strconv"
	"unique"
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

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

const baseURL = "https://adventofcode.com"

func GetInput(year int, day int) ([]byte, error) {
	url, err := url.Parse(fmt.Sprintf("%s/%d/day/%d/input", baseURL, year, day))
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	client.Jar = jar
	client.Jar.SetCookies(url, []*http.Cookie{{
		Name:  "session",
		Value: os.Getenv("AOC_SESSION_ID"),
	}})

	req := &http.Request{
		Method: "GET",
		URL:    url,
		Header: make(map[string][]string),
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func TransposeMatrix[S ~[][]E, E any](s S) S {
	var transposed [][]E
	for i := 0; i < len(s[0]); i++ {
		var row []E
		for j := 0; j < len(s); j++ {
			row = append(row, s[j][i])
		}
		transposed = append(transposed, row)
	}

	return transposed
}

// calculates area of polygon using shoelace formula
func AreaOfPolygon(corners [][2]int) int {
	var sum int
	for i := 0; i < len(corners); i++ {
		sum += corners[i][0] * (corners[CircularIndex(i-1, len(corners))][1] - corners[CircularIndex(i+1, len(corners))][1])
	}

	return sum
}

func CircularIndex(i int, size int) int {
	if i < 0 {
		for i < 0 {
			i += size
		}
	}
	return i % size
}

func UniqueSlice[T comparable, S ~[]T](s S) []unique.Handle[T] {
	var u []unique.Handle[T]

	for _, se := range s {
		u = append(u, unique.Make(se))
	}

	return u
}

func IntersectSlices[S ~[]E, E comparable](a S, b S) S {
	var r S
	for _, e := range a {
		if slices.Contains(b, e) {
			r = append(r, e)
		}
	}
	return r
}

func BoolToByte(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}
