package main

import (
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile  = "input.txt"
	outputFile = "output.txt"
	errStr     = "error"
)

func isSorted(n int) bool {
	next := n % 10
	n = n / 10
	for n > 0 {
		digit := n % 10
		if digit > next {
			return false
		}
		next = digit
		n = n / 10
	}
	return true
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

	lines := strings.Split(string(content), "\n")
	if len(lines) != 3 {
		fOut.Write([]byte(errStr))
		return
	}

	start, err := strconv.Atoi(lines[0])
	if err != nil {
		fOut.Write([]byte(errStr))
		return
	}

	end, err := strconv.Atoi(lines[1])
	if err != nil {
		fOut.Write([]byte(errStr))
		return
	}

	lowest := int(math.Ceil(math.Sqrt(float64(start))))
	highest := int(math.Ceil(math.Sqrt(float64(end))))

	count := 0
	for i := lowest; i <= highest; i++ {
		if isSorted(i*i) && isSorted(i) {
			count++
		}
	}

	fOut.Write([]byte(strconv.Itoa(count)))
}
