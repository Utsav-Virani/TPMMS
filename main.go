package main

import (
	"TPMMS/config"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	err := recreateDirectory("firstpass")
	if err != nil {
		log.Fatal(err)
	}
	// Remove "secondpass" directory if it exists, and create a new one
	err = recreateDirectory("secondpass")
	if err != nil {
		log.Fatal(err)
	}
	fileSplitter("firstpass/r1_", config.Inputfile1)
	fileSplitter("firstpass/r2_", config.Inputfile2)
}

func recreateDirectory(dirName string) error {
	// Check if directory exists
	if _, err := os.Stat(dirName); err == nil {
		// Directory exists, remove it
		err := os.RemoveAll(dirName)
		if err != nil {
			return err
		}
	}

	// Create the directory
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		return err
	}
	return nil
}

func fileSplitter(OutputfilePrefix string, filename string) {
	inputFileHandle, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFileHandle.Close()
	scanner := bufio.NewScanner(inputFileHandle)
	fileCount := 1
	linesWritten := 0

	if err != nil {
		log.Fatal(err)
	}
	// defer outputFile.Close()
	var outputFile *os.File

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			if linesWritten >= config.NumLinesPerFile {
				if outputFile != nil {
					outputFile.Close()
				}
				fileCount++
				var err error
				outputFile, err = createOutputFile(OutputfilePrefix, fileCount)
				if err != nil {
					log.Fatal(err)
				}
				linesWritten = 0
			}
			if linesWritten == 0 {
				outputFile, err = createOutputFile(OutputfilePrefix, fileCount)
				if err != nil {
					log.Fatal(err)
				}
			}
			_, err := outputFile.WriteString(line + "\n")
			if err != nil {
				log.Fatal(err)
			}
			linesWritten++
		} else {
			break
		}
	}

	if outputFile != nil {
		outputFile.Close()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func createOutputFile(prefix string, count int) (*os.File, error) {
	filename := fmt.Sprintf("%s_%d.txt", prefix, count)
	return os.Create(filename)
}
