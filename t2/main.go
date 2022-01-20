package main

import (
	"fmt"
	"strings"
)

var exampleSting = "4539 1488 0343 6464"

func main() {
	var errors []string

	numbers := strings.ReplaceAll(exampleSting, " ", "")

	// Validate the Card number length
	if len(numbers) != 16 {
		errors = append(errors, "Invalid credit card number")
	}

	for _, error := range errors {
		fmt.Println(error)
	}

}
