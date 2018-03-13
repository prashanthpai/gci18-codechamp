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

var (
	globalClount int
	globalOrder  int
)

func satisfiesOrder(s string) bool {
	var digit, sum int
	num, _ := strconv.Atoi(s)
	for num > 0 {
		digit = num % 10
		sum = sum + digit
		num /= 10
	}
	if sum == globalOrder {
		return true
	}
	return false
}

func fixedLengthPerms(set []string, prefix string, n, k int) {

	if k == 0 {
		if satisfiesOrder(prefix) {
			globalClount++
		}
		return
	}

	for i := 0; i < n; i++ {
		newPrefix := prefix + set[i]
		fixedLengthPerms(set, newPrefix, n, k-1)
	}
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
	if len(lines) > 2 {
		fOut.Write([]byte(errStr))
		return
	}

	globalOrder, err = strconv.Atoi(lines[0])
	if err != nil {
		fOut.Write([]byte(errStr))
		return
	}

	set := []string{"1", "2"}
	minLen := int(math.Ceil(float64(globalOrder / 2)))
	for i := minLen; i <= globalOrder; i++ {
		fixedLengthPerms(set, "", len(set), i)
	}

	fOut.Write([]byte(strconv.Itoa(globalClount)))
}
