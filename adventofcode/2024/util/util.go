package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const solutionsDir = "/Repos/solutions/adventofcode/2024"

func ReadInput(fileName string) []string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	var dir string
	if strings.Contains(fileName, "_") {
		dir = strings.Split(fileName, "_")[0]
	} else {
		dir = strings.Split(fileName, ".")[0]
	}

	file, err := os.Open(homeDir + solutionsDir + string(os.PathSeparator) + dir + string(os.PathSeparator) + fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	text := make([]string, 0)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	return text
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ParseToIntArray(str string, sep string) []int {
	var res []int

	values := strings.Split(str, sep)
	for i := range values {
		res = append(res, Atoi(values[i]))
	}

	return res
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func MergeSlices(a, b []int) []int {
	res := make([]int, 0, len(a)+len(b))

	for i := range a {
		res = append(res, a[i])
	}
	for i := range b {
		res = append(res, b[i])
	}

	return res
}
