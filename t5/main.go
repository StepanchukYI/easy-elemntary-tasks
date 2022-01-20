package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calcLackyCount(min, max int) (count int) {

	for i := min; i <= max; i++ {
		ex := fmt.Sprintf("%d", i)
		p1 := strings.Split(ex[:len(ex)/2], "")
		p2 := strings.Split(ex[len(ex)/2:], "")

		sum1 := 0
		sum2 := 0

		for _, num := range p1 {
			val, err := strconv.Atoi(num)
			if err != nil {
				// handle error
				fmt.Println("P1")
				fmt.Println(i)
				fmt.Println(err)
				fmt.Println(num)
				fmt.Println(p1)
				os.Exit(2)
			}
			sum1 += val
		}

		for _, num := range p2 {
			val, err := strconv.Atoi(num)
			if err != nil {
				// handle error
				fmt.Println("P2")
				fmt.Println(i)
				fmt.Println(err)
				fmt.Println(num)
				fmt.Println(p2)
				os.Exit(2)
			}
			sum2 += val
		}

		if sum1 == sum2 {
			count++
		}
	}

	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please, enter min:")
	scanner.Scan()
	min := scanner.Text()
	if len(min) != 6 {
		fmt.Println("Input numbers contains exactly 6 digits. Not more or less")
		return
	}
	minVal, err := strconv.Atoi(min)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	if minVal < 99999 {
		fmt.Println("invalid input")
		return
	}

	fmt.Println("Please, enter max:")
	scanner.Scan()
	max := scanner.Text()
	if len(max) != 6 {
		fmt.Println("Input numbers contains exactly 6 digits. Not more or less")
		return
	}
	maxVal, err := strconv.Atoi(max)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	if maxVal > 999999 {
		fmt.Println("invalid input")
		return
	}

	fmt.Println(calcLackyCount(minVal, maxVal))
}
