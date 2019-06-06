package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var currLevel = 0

func dirTree(out io.Writer, path string, printFiles bool) error {

	return buildTree(out, path, printFiles)
}

func buildTree(out io.Writer, path string, printFiles bool) error {
	dir, err := os.Open(path)
	defer dir.Close()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return err
	}

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return err
	}

	fileInfo, _ := dir.Stat()
	if fileInfo.IsDir() {
		// it's a directory
		currLevel++
		files, err := dir.Readdir(-1)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		}

		files = SortFiles(files)

		for i, file := range files {
			theLast := i+1 == len(files)
			if file.IsDir() {
				newPath := path + string(filepath.Separator) + file.Name()
				PrintItem(out, file, theLast, currLevel)
				dirTree(out, newPath, printFiles)
			} else {
				PrintItem(out, file, theLast, currLevel)
			}
			if theLast {
				currLevel--
			}
		}
	} else {
		// it's a file
		PrintItem(out, fileInfo, true, 0)
	}

	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

// PRINT
const filePrefix = "├───"
const lastFilePrefix = "└───"
const shiftSpace = "	"
const starterLine = "│"

// PrintItem print into output
func PrintItem(out io.Writer, fileInfo os.FileInfo, theLast bool, level int) {
	isDir := fileInfo.IsDir()
	name := fileInfo.Name()

	shift := strings.Repeat(shiftSpace, level)

	var prefix string
	if theLast {
		prefix = lastFilePrefix
	} else {
		prefix = filePrefix
	}

	if isDir {
		fmt.Fprintf(out, "%s%s%s\n", shift, prefix, name)
	} else {
		fileSize := fileInfo.Size()
		fmt.Fprintf(out, "%s%s%s (%d)\n", shift, prefix, name, fileSize)
	}
}

// SORT
type byFileInfo []os.FileInfo

func (a byFileInfo) Len() int {
	return len(a)
}
func (a byFileInfo) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byFileInfo) Less(i, j int) bool { return a[i].Name() < a[j].Name() }

// SortFiles sort files
func SortFiles(files []os.FileInfo) []os.FileInfo {
	sort.Sort(byFileInfo(files))
	return files
}
