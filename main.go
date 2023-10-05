package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	dryFlag := flag.Bool("dry", true, "This flag tells the program to execute the renaming of the files, if not provided it will do a dry run to showcase the expected output behaviour.")
	flag.Parse()
	walkDir := "./sample"
	toRename := make(map[string][]string)
	filepath.Walk(walkDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		currDir := filepath.Dir(path)
		if m, err := match(info.Name()); err == nil {
			key := filepath.Join(currDir, fmt.Sprintf("%s.%s", m.base, m.ext))
			toRename[key] = append(toRename[key], info.Name())
		}
		return nil
	})
	// for _, files := range toRename {
	// 	for _, f := range files {
	// 		fmt.Printf("%q\n", f)
	// 	}
	// }
	for key, files := range toRename {
		dir := filepath.Dir(key)
		n := len(files)
		sort.Strings(files)
		for fidx, filename := range files {
			res, _ := match(filename)
			newFilename := fmt.Sprintf("%s - %d of %d.%s", res.base, fidx+1, n, res.ext)
			oldPath := filepath.Join(dir, filename)
			newPath := filepath.Join(dir, newFilename)
			fmt.Printf("mv %s => %s\n", oldPath, newPath)
			if !*dryFlag {
				err := os.Rename(oldPath, newPath)
				if err != nil {
					log.Println("Error renaming:", oldPath, newPath, err)
				}
			}
		}
	}
}

type matchResult struct {
	base  string
	ext   string
	index int
}

// Match returns the new filename or an error
func match(filename string) (*matchResult, error) {
	pieces := strings.Split(filename, ".")
	ext := pieces[len(pieces)-1]
	temp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(temp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return nil, fmt.Errorf("%s did not match the pattern", filename)
	}
	return &matchResult{
		base:  name,
		ext:   ext,
		index: number,
	}, nil
}
