package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		fmt.Printf("Error with file: %v", err)
	}

	scanner := bufio.NewScanner(file)
	safe_reports := 0
	safe_reports_with_deletion := 0

	for scanner.Scan() {
		var line []string = strings.Split(scanner.Text(), " ")

		var line_numbers []int

		for i := 0; i < len(line); i++ {
			number, err := strconv.ParseInt(line[i], 10, 8)
			if err != nil {
				return
			}
			line_numbers = append(line_numbers, int(number))
		}

		if isReportSafe(line_numbers) {
			safe_reports += 1
		} else if isReportSafeWithDeletion(line_numbers) {
			safe_reports_with_deletion += 1
		}

	}
	fmt.Printf("safe_reports: %v\n", safe_reports)
	fmt.Printf("safe_reports_with_deletion: %v\n", safe_reports_with_deletion+safe_reports)

}

func isReportSafe(report []int) bool {
	flagIncrease, flagDecrease := false, false

	for i := 0; i < len(report)-1; i++ {

		diff := report[i] - report[i+1]

		if diff < 0 {
			flagIncrease = true
		} else if diff > 0 {
			flagDecrease = true
		} else {
			return false
		}

		if flagIncrease && flagDecrease {
			return false
		}

		if diff > 3 || diff < -3 {
			return false
		}
	}
	return true
}

func isReportSafeWithDeletion(report []int) bool {
	var isSafe bool
	for deleteIndex := 0; deleteIndex < len(report); deleteIndex++ {

		copyReport := make([]int, len(report))
		copy(copyReport, report)

		if deleteIndex == len(report)-1 {
			copyReport = copyReport[:deleteIndex]
		} else {
			copyReport = append(copyReport[:deleteIndex], copyReport[deleteIndex+1:]...)
		}

		isSafe = isReportSafe(copyReport)
		if isSafe {
			return true
		}
	}
	return false
}
