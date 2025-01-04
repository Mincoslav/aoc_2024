package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		fmt.Printf("Error with file: %v", err)
	}

	var left_column_array []int64
	var right_column_array []int64

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")
		left_int, left_error := strconv.ParseInt(line[0], 10, 64)
		right_int, right_error := strconv.ParseInt(line[1], 10, 64)

		if left_error == nil && right_error == nil {
			left_column_array = append(left_column_array, left_int)
			right_column_array = append(right_column_array, right_int)
		}
	}

	slices.Sort(left_column_array)
	slices.Sort(right_column_array)

	// Part 1 solution:
	// var total_difference int64
	// for i := 0; i < 1000; i++ {
	// 	var difference int64 = left_column_array[i] - right_column_array[i]
	// 	total_difference += int64(math.Abs(float64(difference)))
	// }
	// fmt.Printf("total_difference:%v", total_difference)

	// Part 2:
	var occurences_map = make(map[int64]int64)

	for i := 0; i < 1000; i++ {
		value, ok := occurences_map[right_column_array[i]]
		if ok {
			occurences_map[right_column_array[i]] = value + 1
		} else {
			occurences_map[right_column_array[i]] = 1
		}
	}
	// fmt.Printf("%v", occurences_map)

	var similarity_score int64 = 0

	for i := 0; i < len(left_column_array); i++ {
		similarity_score += (left_column_array[i] * occurences_map[left_column_array[i]])
	}

	fmt.Printf("similarity_score: %v\n", similarity_score)
}
