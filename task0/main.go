package main

import (
	"bufio"
	"fmt"
	"os"
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
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please, enter width:")
	scanner.Scan()
	x := scanner.Text()
	xVal, err := strconv.Atoi(x)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println("Please, enter Height:")
	scanner.Scan()
	y := scanner.Text()

	yVal, err := strconv.Atoi(y)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println("Please, enter Symbol to write:")
	scanner.Scan()
	symbol := scanner.Text()

	printBoard(xVal, yVal, symbol)
}
