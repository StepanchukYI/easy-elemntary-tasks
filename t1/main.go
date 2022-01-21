package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var exampleSting = "30,-1,-6,90,-6,22,52,123,2,35,6"

func posOrNeg(numbers []int)

func main() {
	var negative []int
	var positive []int
	numbers := strings.Split(exampleSting, ",")

	for _, value := range numbers {
		v, err := strconv.Atoi(value)

		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		if v > 0 {
			positive = append(positive, v)
		} else {
			negative = append(negative, v)
		}
	}

	fmt.Println("Positive Numbers:")
	printSlice(positive)
	fmt.Println("Negative Numbers:")
	printSlice(negative)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
