package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func renameFile(oldPath, newPath string) {
	err := os.Rename(oldPath, newPath)
	if err != nil {
		fmt.Printf("Error renaming file %s to %s: %v\n", oldPath, newPath, err)
	} else {
		fmt.Printf("File renamed from %s to %s\n", oldPath, newPath)
	}
}

func renameFilesInDirectory(directory string) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			oldPath := filepath.Join(directory, file.Name())
			newFileName := strings.ReplaceAll(file.Name(), "_", " ")
			newPath := filepath.Join(directory, newFileName)
			renameFile(oldPath, newPath)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the directory as an argument.")
		os.Exit(1)
	}

	directory := os.Args[1]

	renameFilesInDirectory(directory)
}