package fizzbuzz

import (
	"strconv"
)

func FizzBuzz(fizzBuzzRange int) string {
	resultStr := ""
	for i := 1; i <= fizzBuzzRange; i++ {
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
		if i != fizzBuzzRange {
			resultStr += ", "
		}
	}
	return resultStr
}
