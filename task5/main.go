package main

import (
	"errors"
	"fmt"
	"log"
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

func validateMin(num int) (res bool, err error) {
	res = true
	if num < 99999 {
		return false, errors.New("invalid input")
	}

	return true, nil
}

func validateMax(num int) (res bool, err error) {
	res = true
	if num > 999999 {
		return false, errors.New("invalid input")
	}

	return true, nil
}

func main() {
	minString, err := getVal("Please, enter min:")
	if err != nil {
		log.Fatal(err)
	}
	if minString == "" {
		log.Fatal(EmptyValueError)
	}
	min, err := strconv.Atoi(minString)
	if err != nil {
		log.Fatal(err)
	}
	minValid, minErr := validateMin(min)
	if !minValid {
		log.Fatal(minErr)
	}

	maxString, err := getVal("Please, enter max:")
	if err != nil {
		log.Fatal(err)
	}
	if maxString == "" {
		log.Fatal(EmptyValueError)
	}
	max, err := strconv.Atoi(maxString)
	if err != nil {
		log.Fatal(err)
	}
	maxValid, maxErr := validateMax(max)
	if !maxValid {
		log.Fatal(maxErr)
	}

	fmt.Println(calcLackyCount(min, max))
}
