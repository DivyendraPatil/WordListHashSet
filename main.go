package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

var hashSet = make(map[string]bool)

func readFileLineByLine(file string){
	if file != "/Volumes/WD SSD/Passwords"  {
		// Open passed file
		currentFile, err := os.Open(file)
		if err != nil {
			panic(err)
		}
		defer currentFile.Close()

		// Initialize reader and read File Line By Libe
		reader := bufio.NewReader(currentFile)
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				break
			}
			_, present := hashSet[string(line)]
			if present != true {
				hashSet[string(line)] = true
			}
		}
	}
}

func main() {
	var files []string

	root := "/Volumes/WD SSD/Passwords"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println("Reading => ",file)
		readFileLineByLine(file)
	}
}