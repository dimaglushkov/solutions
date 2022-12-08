package util

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const solutionsDir = "/Repos/solutions/adventofcode/2022"

func ReadInput(fileName string) []string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	dir := strings.Split(fileName, "_")[0]
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
