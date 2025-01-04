package utils

import (
	"bufio"
	"fmt"
	"os"
)

func getScannerForFile(filePath string) *bufio.Scanner {
	
	file, err := os.Open(filePath)
	
	if err != nil {
		fmt.Printf("Error with file: %v", err)
	}
	
	scanner := bufio.NewScanner(file)
	return scanner
}