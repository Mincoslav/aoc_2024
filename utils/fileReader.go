package utils

import (
	"bufio"
	"fmt"
	"os"
)

func GetScannerForFile(filePath string) *bufio.Scanner {
	
	file, err := os.Open(filePath)
	
	if err != nil {
		fmt.Printf("Error with file: %v", err)
	}
	
	scanner := bufio.NewScanner(file)
	return scanner
}


func GetFileContentAsString(filePath string) string {

	file, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Printf("Error with file: %v", err)
	}

	contentString := string(file)
	return contentString
}
