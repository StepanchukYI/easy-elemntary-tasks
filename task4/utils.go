package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var EmptyValueError = errors.New("Value should be not empty!")

func getVal(promnt string) (val string, err error) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(promnt)

	scanner.Scan()
	val = scanner.Text()
	if err := scanner.Err(); err != nil {
		return val, err
	}

	return
}
