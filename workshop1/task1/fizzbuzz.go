package main

import (
	"fmt"
	"strconv"
)

func main() {
	resultStr := ""
	for i := 1; i < 100; i++ {
		firstCheck, secondCheck := i%3 == 0, i%5 == 0
		if firstCheck {
			resultStr += "Fizz"
			if secondCheck {
				resultStr += " Buzz"
			}
		} else if secondCheck {
			resultStr += "Buzz"
		} else {
			resultStr += strconv.Itoa(i)
		}
		resultStr += ", "
	}
	fmt.Println(resultStr)
}
