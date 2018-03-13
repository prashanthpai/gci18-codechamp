package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	inputFile     = "input.txt"
	outputFile    = "output.txt"
	errStr        = "error"
	invalidErrStr = "invalid"
	format        = "02/01/2006"
)

var sentinelDate = time.Date(2018, time.Month(3), 10, 0, 0, 0, 0, time.UTC)

func parseDate(value string) (time.Time, error) {

	t, err := time.Parse(format, value)
	if err != nil {
		return t, err
	}

	if t.Format(format) != value {
		return t, fmt.Errorf("invalid time")
	}

	return t, nil
}

func main() {

	fOut, err := os.Create(outputFile)
	if err != nil {
		return
	}
	defer fOut.Close()

	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fOut.Write([]byte(errStr))
		return
	}

	dates := strings.Split(string(content), " ")
	if len(dates) != 2 {
		fOut.Write([]byte(errStr))
		return
	}

	t1, err := parseDate(dates[0])
	if err != nil {
		fOut.Write([]byte(invalidErrStr))
		return
	}

	t2, err := parseDate(strings.TrimSpace(dates[1]))
	if err != nil {
		fOut.Write([]byte(invalidErrStr))
		return
	}

	if t1.Before(sentinelDate) || t2.Before(sentinelDate) {
		fOut.Write([]byte(invalidErrStr))
		return
	}

	if !t1.Before(t2) {
		fOut.Write([]byte(invalidErrStr))
		return
	}

	count := 0
	for !t1.After(t2) {
		if t1.Weekday() != time.Saturday && t1.Weekday() != time.Sunday {
			count++
		}
		t1 = t1.AddDate(0, 0, 1)
	}

	fOut.Write([]byte(strconv.Itoa(count)))
}
