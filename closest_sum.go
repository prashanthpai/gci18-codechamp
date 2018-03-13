package main

import (
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	inputFile  = "input.txt"
	outputFile = "output.txt"
	errStr     = "error"
	maxUint    = ^uint(0)
	maxInt     = int(maxUint >> 1)
)

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
	if len(lines) < 4 {
		fOut.Write([]byte(errStr))
		return
	}

	target, err := strconv.Atoi(lines[0])
	if err != nil {
		fOut.Write([]byte(errStr))
		return
	}

	n, err := strconv.Atoi(lines[1])
	if err != nil {
		fOut.Write([]byte(errStr))
		return
	}

	var ints []int
	for i := 2; i < n+2; i++ {
		tmp, err := strconv.Atoi(lines[i])
		if err != nil {
			fOut.Write([]byte(errStr))
			return
		}
		ints = append(ints, tmp)
	}

	sort.Ints(ints)

	var n1, n2 int
	l := 0
	r := n - 1
	delta := maxInt
	for r > l {
		if int(math.Abs(float64(ints[l]+ints[r]-target))) < delta {
			n1 = ints[l]
			n2 = ints[r]
			delta = int(math.Abs(float64(ints[l] + ints[r] - target)))
		}

		if ints[l]+ints[r] > target {
			r--
		} else {
			l++
		}
	}

	fOut.Write([]byte(strconv.Itoa(n1) + "\n"))
	fOut.Write([]byte(strconv.Itoa(n2)))
}
