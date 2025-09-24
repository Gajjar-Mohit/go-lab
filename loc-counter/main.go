package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var currentWorkingDir = GetCurrentDirectory()
	var files = ListAllFoldersInDirectory(currentWorkingDir)
	fmt.Println(files)
}

func ReadFile(fileName string) {
	fmt.Println("Reading file")
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println("file name " + fileName)
	fmt.Printf("file size %d\n", len(data))
}

func GetCurrentDirectory() string {
	var currentDir, error = os.Getwd()
	if error != nil {
		fmt.Println("Error:", error)
	}
	return currentDir
}

func ListAllFoldersInDirectory(directory string) []string {
	var files []string
	var entries, error = os.ReadDir(directory)

	if error != nil {
		fmt.Println("Error:", error)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			// fmt.Printf("DIR: %s\n", entry.Name())
			files = append(files, entry.Name())
		} else {
			// fmt.Printf("FILE: %s\n", entry.Name())
		}
	}
	return files
}

func CountNoOfLines(file string) (int, error) {
	var data, error = os.ReadFile(file)

	if error != nil {
		fmt.Println("Error:", error)
	}
	var lineCount = bytes.Count(data, []byte{'\n'})

	if len(data) > 0 && data[len(data)-1] != '\n' {
		lineCount++
	}

	return lineCount, nil
}
