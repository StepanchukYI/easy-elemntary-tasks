package main

import (
	"fmt"
	"log"
	cu "customutils"
)

func reverseSting(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func findFullStringPolindrome(number string) bool {
	var s1, s2 string
	if len(number)%2 != 0 {
		lenPart := len(number) / 2

		s1 = number[:lenPart+1]
		s2 = number[lenPart:]
	} else {
		lenPart := len(number) / 2

		s1 = number[:lenPart]
		s2 = number[lenPart:]
	}

	return isPolindrome(s1, reverseSting(s2))
}

func findFirstStringPolindrome(number string) []string {
	var s1, s2 string
	startFrom := 0
	lens := len(number)

	fmt.Println(lens)
	for i := 1; i < len(number)/2; i++ {
		lenth := i - startFrom

		s1 = number[startFrom:lenth]
		s2 = number[lenth : lenth+lenth]
		if isPolindrome(s1, reverseSting(s2)) {

			return []string{s1, s2}
		}
	}

	return []string{}
}

func findAllStringPolindrome(number string) (res [][]string) {
	var s1, s2 string

	for startFrom := 0; startFrom < len(number)-2; startFrom++ {
		for i := 1; i < len(number); i++ {
			lenth := i + startFrom

			if startFrom+lenth >= len(number) {
				continue
			}

			if lenth+i >= len(number) {
				continue
			}

			s1 = number[startFrom:lenth]
			s2 = number[lenth : lenth+i]
			fmt.Println("S1:" + s1)
			fmt.Println("S2:" + s2)
			if isPolindrome(s1, reverseSting(s2)) {
				res = append(res, []string{s1, s2})
			}
		}
		fmt.Println("End of iter")
	}

	return
}

func isPolindrome(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}
	return false
}

func easyLevel(number string) {
	fmt.Println("Is Polindrome inside " + number + "?")
	fmt.Println(findFullStringPolindrome(number))
}

func midLevel(number string) {
	fmt.Println("First Polindrome inside " + number + "?")
	ret := findAllStringPolindrome(number)
	if len(ret) == 0 {
		fmt.Println("no any Polindrome inside")
		return
	}
	fmt.Println(ret)
}

func hardLevel(number string) {
	fmt.Println("All Polindrome inside " + number + "?")
	ret := findAllStringPolindrome(number)
	if len(ret) == 0 {
		fmt.Println("no any Polindrome inside")
		return
	}
	fmt.Println(ret)
}

func validateInput(num int) (res bool, err []string) {
	res = true

	if num < 10 {
		err = append(err, "invalid input")
	}

	if len(err) > 0 {
		res = false
	}

	return
}

func main() {
	inputString, err := cu.GetVal("Please, enter your number:")
	if err != nil {
		log.Fatal(err)
	}

	if inputString == "" {
		log.Fatal(cu.EmptyValueError)
	}

	easyLevel(inputString)

	midLevel(inputString)

	hardLevel(inputString)
}
