package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"github.com/mincoslav/aoc/utils"
)

func findRegexMatches(input string, regexExpression string) []string {
	regex := regexp.MustCompile(regexExpression)
	return regex.FindAllString(input, -1)
}

func main() {
	var file = utils.GetFileContentAsString("./input.txt")
	var multiplicationSums int64 = 0

	// 	// Part 2:
	var invalidMultiplications = findRegexMatches(file, `don't\(\)(.*?)((do\(\))|(\n.*?do\(\)))`)
	// fmt.Printf("len(invalidMultiplications): %v\n", len(invalidMultiplications))

	for _, invalid_string := range invalidMultiplications {
		file = strings.ReplaceAll(file, invalid_string, "")
	}
	// End of Part 2

	// main part -------------------
	var results = findRegexMatches(file, `mul\(\d{1,3},\d{1,3}\)`)
		fmt.Printf("len(results): %v\n", len(results))
		for _, result := range results {

			var numbers = findRegexMatches(result, `\d{1,3}`)
			left_int, left_error := strconv.ParseInt(numbers[0], 10, 64)
			right_int, right_error := strconv.ParseInt(numbers[1], 10, 64)

			if left_error == nil && right_error == nil {
				multiplicationSums += (left_int * right_int)
			}

		}

	fmt.Printf("multiplicationSums: %v", multiplicationSums)
}
