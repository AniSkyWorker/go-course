package fizzbuzz

import (
	"bytes"
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

func FizzBuzzBuffer(fizzBuzzRange int) string {
	var resultStr bytes.Buffer
	for i := 1; i <= fizzBuzzRange; i++ {
		firstCheck, secondCheck := i%3 == 0, i%5 == 0
		if firstCheck {
			resultStr.WriteString("Fizz")
			if secondCheck {
				resultStr.WriteString(" Buzz")
			}
		} else if secondCheck {
			resultStr.WriteString("Buzz")
		} else {
			resultStr.WriteString(strconv.Itoa(i))
		}
		if i != fizzBuzzRange {
			resultStr.WriteString(", ")
		}
	}
	return resultStr.String()
}
