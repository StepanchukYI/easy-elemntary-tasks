package main

import (
	"fmt"
	"log"
	"strconv"
)

func printBoard(x int, y int, symbol string) {
	for xRange := 0; xRange <= x; xRange++ {
		for yRange := 0; yRange <= y; yRange++ {
			fmt.Printf("%v ", symbol)
		}
		fmt.Printf("\n")
	}
}

func main() {
	x, err := getVal("Please, enter width: ")
	if err != nil {
		log.Fatal(err)
	}
	if x == "" {
		log.Fatal(EmptyValueError)
	}
	xVal, err := strconv.Atoi(x)
	if err != nil {
		log.Fatal(err)
	}

	y, err := getVal("Please, enter height: ")
	if err != nil {
		log.Fatal(err)
	}
	if x == "" {
		log.Fatal(EmptyValueError)
	}
	yVal, err := strconv.Atoi(y)
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := getVal("Please, enter Symbol to write: ")
	if err != nil {
		log.Fatal(err)
	}
	if symbol == "" {
		log.Fatal(EmptyValueError)
	}

	printBoard(xVal, yVal, symbol)
}
