package customutils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var EmptyValueError = errors.New("Value should be not empty!")

func GetVal(promnt string) (val string, err error) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(promnt)

	scanner.Scan()
	val = scanner.Text()
	if err := scanner.Err(); err != nil {
		return val, err
	}

	return
}
