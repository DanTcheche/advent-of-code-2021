package utils

import (
	"bufio"
	"fmt"
	"os"
)

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func GetFileData(filePath string) []string {
	var file, err = os.Open(filePath)
	if isError(err) {
		panic("File couldn't be read")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text
}
