package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need one argument")
		os.Exit(1)
	}

	path := os.Args[1]
	fullPath, _ := filepath.Abs(path)
	filesMap := findMatchedFiles(fullPath)
	renameFiles(filesMap)
}

// renameFiles renames files in the map
func renameFiles(filesMap map[string][]string) {
	for basename, fullpaths := range filesMap {
		total := len(filesMap[basename])
		for i, file := range fullpaths {
			targetDir := filepath.Dir(file)
			newfile := fmt.Sprintf("%s/%s_%d_of_%d.txt\n", targetDir, basename, i+1, total)
			fmt.Printf("moving %s to %s\n", file, newfile)
			os.Rename(file, newfile)
		}
	}

}

// findMatchedFiles walks a directory, selects files with names matching xxxx_nnn.txt
// and returns a map of files grouped by their name.
// Example: myfile_123.txt myfile_456.txt yourfile_123.txt yourfile_456.txt
// Return:   myfile : [myfile_123.txt myfile_456.txt]
//         yourfile : [yourfile_123.txt yourfile_456.txt]
func findMatchedFiles(path string) map[string][]string {
	var m = make(map[string][]string)

	_ = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %s \n", path)
			return err
		}
		if !d.IsDir() {
			re := regexp.MustCompile(`(\w+)_(\d+)\.(txt)`)
			found := re.FindStringSubmatch(d.Name())
			if found != nil {
				name := found[1]
				m[name] = append(m[name], path)
			}
		}
		return err
	})
	return m
}
