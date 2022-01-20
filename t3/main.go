package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	lastNum := 0
	beforeLastNum := 0
	return func() int {
		ret := lastNum + beforeLastNum
		if beforeLastNum == 0 {
			lastNum = 1
		}
		beforeLastNum = lastNum
		lastNum = ret
		return ret
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
