package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func parsePatchFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file %s: %w", filePath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var outFile *os.File
	var outFileName string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "diff --git") {
			if outFile != nil {
				outFile.Close()
			}
			fileNames := strings.Split(line, " ")
			filename := fileNames[2][2:]
			outFileName := strings.ReplaceAll(filename, "/", "_") + ".patch"
			outFile, err = os.Create(outFileName)
			if err != nil {
				return fmt.Errorf("error creating file %s: %w", outFileName, err)
			}
		}

		if outFile != nil {
			_, err = outFile.WriteString(line + "\n")
			if err != nil {
				return fmt.Errorf("error writing to file %s: %w", outFileName, err)
			}
		}
	}
	if outFile != nil {
		outFile.Close()
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	return nil
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Usage: patch_split <file_path>")
		os.Exit(1)
	}
	filePath := flag.Arg(0)

	err := parsePatchFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
