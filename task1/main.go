package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var exampleSting = "30,-1,-6,90,-6,22,52,123,2,35,6"

func posOrNeg(numbers []string) (positive, negative []int) {
	for _, value := range numbers {
		v, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}

		if v > 0 {
			positive = append(positive, v)
		} else {
			negative = append(negative, v)
		}
	}
	return
}

func main() {
	numbers := strings.Split(exampleSting, ",")

	positive, negative := posOrNeg(numbers)

	fmt.Println("Positive Numbers:")
	printSlice(positive)
	fmt.Println("Negative Numbers:")
	printSlice(negative)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
