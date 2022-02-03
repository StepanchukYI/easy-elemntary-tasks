package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

var exampleSting = "4539 1488 0343 6464"

func validateCardNumber(card string) (valid bool, err error) {

	// Validate the Card number length
	if len(card) != 16 {
		return false, errors.New("Card number must be 16 characters long")
	}

    for _, c := range card {
        if c < '0' || c > '9' {
            return false, errors.New("Card number must contains numbers only")
        }
    }

	return true, nil
}

func lastFour(card string) (string) {
	return card[len(card)-4 : len(card)]
}

func printLastFour(promnt string,card string){
	fmt.Println(promnt)
	
	res := "**** **** **** " + lastFour(card)

	fmt.Println(res)
}

func main() {
	numbers := strings.ReplaceAll(exampleSting, " ", "")

	valid, error := validateCardNumber(numbers)
	if !valid {
		log.Fatal(error)
	}

	printLastFour("Card Is Valid", numbers)
}
