package util

import (
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
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
