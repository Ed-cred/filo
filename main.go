package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var (
	re            = regexp.MustCompile("^(.+?) ([0-9]{4}) [(]([0-9]+) of ([0-9]+)[)][.](.+?)$")
	replaceString = "$2 - $1 - $3 of $4.$5"
)

func main() {
	dryFlag := flag.Bool("dry", true, "This flag tells the program to execute the renaming of the files, if not provided it will do a dry run to showcase the expected output behaviour.")
	flag.Parse()
	walkDir := "./sample"
	var toRename []string
	filepath.Walk(walkDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if _, err := match(info.Name()); err == nil {
			toRename = append(toRename, path)
		}
		// currDir := filepath.Dir(path)
		// if m, err := match(info.Name()); err == nil {
		// 	key := filepath.Join(currDir, fmt.Sprintf("%s.%s", m.base, m.ext))
		// 	toRename[key] = append(toRename[key], info.Name())
		// }
		return nil
	})
	for _, oldPath := range toRename {
		dir := filepath.Dir(oldPath)
		filename := filepath.Base(oldPath)
		newFilename, err := match(filename)
		if err != nil {
			log.Println(err)
		}
		newPath := filepath.Join(dir, newFilename)
		fmt.Printf("mv %s => %s\n", oldPath, newPath)
		if !*dryFlag {
			err := os.Rename(oldPath, newPath)
			if err != nil {
				log.Println("Error renaming:", oldPath, newPath, err)
			}
		}
	}
	// for _, files := range toRename {
	// 	for _, f := range files {
	// 		fmt.Printf("%q\n", f)
	// 	}
	// }
	// for key, files := range toRename {
	// 	dir := filepath.Dir(key)
	// 	n := len(files)
	// 	sort.Strings(files)
	// 	for fidx, filename := range files {
	// 		res, _ := match(filename)
	// 		newFilename := fmt.Sprintf("%s - %d of %d.%s", res.base, fidx+1, n, res.ext)
	// 		oldPath := filepath.Join(dir, filename)
	// 		newPath := filepath.Join(dir, newFilename)
	// 	}
	// }
}

// type matchResult struct {
// 	base  string
// 	ext   string
// 	index int
// }

// Regex version of match
func match(filename string) (string, error) {
	if !re.MatchString(filename) {
		return "", fmt.Errorf("%s did not match the pattern", filename)
	}
	return re.ReplaceAllString(filename, replaceString), nil
}

// Match returns the new filename or an error
// func match(filename string) (*matchResult, error) {
// pieces := strings.Split(filename, ".")
// 	ext := pieces[len(pieces)-1]
// 	temp := strings.Join(pieces[0:len(pieces)-1], ".")
// 	pieces = strings.Split(temp, "_")
// 	name := strings.Join(pieces[0:len(pieces)-1], "_")
// 	number, err := strconv.Atoi(pieces[len(pieces)-1])
// 	if err != nil {
// 		return nil, fmt.Errorf("%s did not match the pattern", filename)
// 	}
// 	return &matchResult{
// 		base:  name,
// 		ext:   ext,
// 		index: number,
// 	}, nil
// }
