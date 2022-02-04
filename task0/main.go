package main

import (
	"fmt"
	"log"
	"strconv"
	cu "customutils"
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
	x, err := cu.GetVal("Please, enter width: ")
	if err != nil {
		log.Fatal(err)
	}
	if x == "" {
		log.Fatal(cu.EmptyValueError)
	}
	xVal, err := strconv.Atoi(x)
	if err != nil {
		log.Fatal(err)
	}

	y, err := cu.GetVal("Please, enter height: ")
	if err != nil {
		log.Fatal(err)
	}
	if x == "" {
		log.Fatal(cu.EmptyValueError)
	}
	yVal, err := strconv.Atoi(y)
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := cu.GetVal("Please, enter Symbol to write: ")
	if err != nil {
		log.Fatal(err)
	}
	if symbol == "" {
		log.Fatal(cu.EmptyValueError)
	}

	printBoard(xVal, yVal, symbol)
}
