package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
)

var hashSet = make(map[string]bool)

func readFileLineByLine(file string){
	if file != "/Volumes/WD SSD/Passwords"  {
		// Open passed file
		readFile, err := os.Open(file)
		if err != nil {
			panic(err)
		}

		// Open file to be written
		writeFile, err := os.Create("/Volumes/Work/wordList.txt")
		if err != nil {
			panic(err)
		}
		// close writeFile on exit and check for its returned error
		defer func() {
			if err := readFile.Close(); err != nil {
				panic(err)
			}
			if err := writeFile.Close(); err != nil {
				panic(err)
			}
		}()

		// Initialize reader and writer from bufio
		reader := bufio.NewReader(readFile)
		writer := bufio.NewWriter(writeFile)

		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				break
			}
			stringLine := string(line)
			_, present := hashSet[stringLine]
			if present != true && len(stringLine) > 5 {
				// fmt.Println(stringLine)
				hashSet[stringLine] = true

				// Write to file
				if _, err := writer.WriteString(stringLine+"\n"); err != nil {
					panic(err)
				}
			}
		}
		if err = writer.Flush(); err != nil {
			panic(err)
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

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	for _, file := range files {
		fmt.Println("Reading => ",file)
		readFileLineByLine(file)
		printMemUsage()
	}
}