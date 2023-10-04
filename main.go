package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := "birthday_001.txt"
	newName, err := match(fileName, 1)
	if err != nil {
		fmt.Println("no match found")
		os.Exit(1)
	}
	fmt.Println(newName)
}

// Match returns the new filename or an error
func match(fileName string, total int) (string, error) {
	pieces := strings.Split(fileName, ".")
	ext := pieces[len(pieces)-1]
	temp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(temp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return "", fmt.Errorf("%s did not match the pattern", fileName)
	}
	return fmt.Sprintf("%s - %d of %d.%s", name, number, total, ext), nil
}
