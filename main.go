package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type rename struct {
	filename string
	path     string
}

func main() {
	// fileName := "birthday_001.txt"
	// newName, err := match(fileName, 1)
	// if err != nil {
	// 	fmt.Println("no match found")
	// 	os.Exit(1)
	// }
	// fmt.Println(newName)
	dir := "./sample"
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	count := 0
	var toRename []rename
	for _, file := range files {
		if file.IsDir() {
		} else {
			_, err := match(file.Name(), 4)
			if err == nil {
				count++
			}
			toRename = append(toRename, rename{
				filename: file.Name(),
				path:     filepath.Join(dir, file.Name()),
			})
		}
	}
	for _, orig := range toRename {
		newFilename, err := match(orig.filename, count)
		if err != nil {
			log.Fatal(err)
		}
		newPath := filepath.Join(dir, newFilename)
		fmt.Printf("mv %s => %s\n", orig.path, newPath)
		err = os.Rename(orig.path, newPath)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Match returns the new filename or an error
func match(filename string, total int) (string, error) {
	pieces := strings.Split(filename, ".")
	ext := pieces[len(pieces)-1]
	temp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(temp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return "", fmt.Errorf("%s did not match the pattern", filename)
	}
	return fmt.Sprintf("%s - %d of %d.%s", name, number, total, ext), nil
}
