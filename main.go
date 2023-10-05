package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type file struct {
	name string
	path string
}
func main() {
	dir := "./sample"
	var toRename []file
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if _, err := match(info.Name()); err == nil {
			toRename = append(toRename, file{
				name: info.Name(),
				path: path,
			})	
		}
		return nil
	})
	for _, orig := range toRename {
		var n file
		var err error
		n.name, err= match(orig.name)
		if err != nil {
			log.Println("Error matching:", orig.path, err)
		}
		n.path = filepath.Join(dir, n.name)
		fmt.Printf("mv %s => %s\n", orig.path, n.path)
		// err = os.Rename(orig.path, n.path)
		// if err != nil {
			// log.Println("Error renaming:", orig.path, err)
		// }
	}
}

// Match returns the new filename or an error
func match(filename string) (string, error) {
	pieces := strings.Split(filename, ".")
	ext := pieces[len(pieces)-1]
	temp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(temp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return "", fmt.Errorf("%s did not match the pattern", filename)
	}
	return fmt.Sprintf("%s - %d.%s", name, number, ext), nil
}
