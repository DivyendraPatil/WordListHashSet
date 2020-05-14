package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

var hashSet = make(map[string]bool)

func readFiles(file string){
	//File reader

	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error in opening reading file",err)
		panic(err)
	}
	defer readFile.Close()
	reader := bufio.NewReader(readFile)

	fileToBeWritten := "/Users/divyendrapatil/IdeaProjects/wordListHashSet/wordList.txt"

	writeFile, err := os.OpenFile(fileToBeWritten, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error in opening writing file")
		panic(err)
	}
	defer writeFile.Close()
	writer := bufio.NewWriter(writeFile)

	if file != "/Users/divyendrapatil/IdeaProjects/wordListHashSet/testGround" {

		for {
			line := ""
			line, err := reader.ReadString('\n')
			if err != nil {
				//fmt.Println("Error in reading line",err)
				break
			}
			_, present := hashSet[line]
			lineCondition := checkLineCondition(present,line)
			if lineCondition == false {
				hashSet[line] = true
				_, err := writer.WriteString(line)
				if err != nil {
					fmt.Printf("Error in writing",err)
				}
				//fmt.Printf("Line %s ", line)
			}

		}
		err := writer.Flush()
		if err != nil {
			fmt.Printf("Error in flushing",err)
		}
	}
}

func checkLineCondition(present bool,line string) (condition bool){
	length := len(line)
	if present == false && length > 5 && length < 20 {
		return false
	}
	return true
}

func main() {
	var files []string

	// root := "/Users/divyendrapatil/IdeaProjects/wordListHashSet/testGround"
	root := "/Volumes/WD SSD/Passwords"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		fmt.Println("Error in Walking through directory")
		fmt.Println(err)
		panic(err)
	}

	//go func() {
	//	log.Println(http.ListenAndServe("localhost:6060", nil))
	//}()

	for _, file := range files {
		printMemUsage(file)
		readFiles(file)
	}
}
