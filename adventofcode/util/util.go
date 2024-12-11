package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

const solutionsDir = "/Repos/solutions/adventofcode/"

func ReadInput(example bool) []string {
	if example {
		fmt.Println("Reading example data")
	}

	_, callerFile, _, _ := runtime.Caller(1)
	callerFileParts := strings.Split(callerFile, string(os.PathSeparator))

	year, day := callerFileParts[len(callerFileParts)-3], callerFileParts[len(callerFileParts)-2]
	suffix := ".input"
	if example {
		suffix = ".example"
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	inputFilePath := path.Join(homeDir, solutionsDir, year, day, day+suffix)
	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

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

func SliceToSet[T comparable](slice []T) map[T]bool {
	set := make(map[T]bool, len(slice))
	for _, v := range slice {
		set[v] = true
	}
	return set
}

func InsertAt[T any](slice []T, value T, index int) []T {
	newSlice := make([]T, len(slice)+1)
	copy(newSlice, slice)
	newSlice[index] = value
	copy(newSlice[index+1:], slice[index:])
	return newSlice
}

func CopyMap[K comparable, V any](src map[K]V) map[K]V {
	dst := make(map[K]V, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func CopyFilterMap[K comparable, V any](src map[K]V, filter func(K) bool) map[K]V {
	dst := make(map[K]V, len(src))
	for k, v := range src {
		if filter(k) {
			dst[k] = v
		}
	}
	return dst
}

func MakeMatrix[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	for i := range matrix {
		matrix[i] = make([]T, m)
	}
	return matrix
}

func CopyMatrix[T any](src [][]T) [][]T {
	dst := make([][]T, len(src))
	for i := range src {
		row := make([]T, len(src[i]))
		copy(row, src[i])
		dst[i] = row
	}
	return dst
}
