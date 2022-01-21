package main

import (
	"bufio"
	"fmt"
	"log"
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

func getVal(scan *bufio.Scanner) (val string) {
	scan.Scan()
	val = scan.Text()

	if val == "" {
		log.Fatal("Empty Imput string")
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func convertStringToInt(val string) (number int) {
	number, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func getIntFromScan(scan *bufio.Scanner) (val int) {
	val = convertStringToInt(getVal(scan))
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please, enter width:")
	xVal := getIntFromScan(scanner)

	fmt.Println("Please, enter Height:")
	yVal := getIntFromScan(scanner)

	fmt.Println("Please, enter Symbol to write:")
	symbol := getVal(scanner)

	printBoard(xVal, yVal, symbol)
}
