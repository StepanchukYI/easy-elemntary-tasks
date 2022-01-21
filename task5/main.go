package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func calcLackyCount(min, max int) (count int) {

	for i := min; i <= max; i++ {
		ex := fmt.Sprintf("%d", i)
		p1 := strings.Split(ex[:len(ex)/2], "")
		p2 := strings.Split(ex[len(ex)/2:], "")

		sum1 := sumStringArr(p1)
		sum2 := sumStringArr(p2)

		if sum1 == sum2 {
			count++
		}
	}

	return
}

func sumStringArr(array []string) (sum int) {
	for _, num := range array {
		val, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		sum += val
	}
	return
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

func validateMin(num int) (res bool) {
	res = true
	if num < 99999 {
		return false
	}

	return true
}

func validateMax(num int) (res bool) {
	res = true
	if num > 999999 {
		return false
	}

	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please, enter min:")
	min := getIntFromScan(scanner)

	if !validateMin(min) {
		log.Fatal("invalid input")
	}

	fmt.Println("Please, enter max:")
	max := getIntFromScan(scanner)

	if !validateMax(max) {
		log.Fatal("invalid input")
	}

	fmt.Println(calcLackyCount(min, max))
}
