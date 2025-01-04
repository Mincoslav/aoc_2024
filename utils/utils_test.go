// utils/utils_test.go
package utils

import (
	"testing"
)

func TestYourFunction(t *testing.T) {
	result := getScannerForFile("../day3/input.txt")
	result.Scan()
	text := result.Text()
	println(text)
	if text == "" {
		t.Error("Expected non-empty string")
	}
}
