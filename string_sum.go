package string_sum

import (
	"errors"
	"fmt"
	"strconv"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	var sum int
	var i int = 0
	var num string
	cnt := 0
	if len(input) == 0 {
		return "", fmt.Errorf("error happened: %w", errorEmptyInput)
	}
	for i < len(input) {
		if string(input[i]) == " " || string(input[i]) == "+" {
			i++
		} else {
			if string(input[i]) == "-" {
				num = num + string("-")
				i++
				for string(input[i]) != " " && string(input[i]) != "-" && string(input[i]) != "+" {
					_, err := strconv.Atoi(string(input[i]))
					if err != nil {
						return "", fmt.Errorf("error happened: %w", errorNotTwoOperands)
					}
					num = num + string(input[i])
					i++
					if i >= len(input) {
						break
					}
				}
				cnt++
			} else {
				for string(input[i]) != " " && string(input[i]) != "-" && string(input[i]) != "+" {
					_, err := strconv.Atoi(string(input[i]))
					if err != nil {
						return "", fmt.Errorf("error happened: %w", errorNotTwoOperands)
					}
					num = num + string(input[i])
					i++
					if i >= len(input) {
						break
					}
				}
				cnt++
			}
			int_num, _ := strconv.Atoi(num)
			sum += int_num
		}
		num = ""
	}
	if cnt != 2 {
		return "", fmt.Errorf("error happened: %w", errorNotTwoOperands)
	}
	return strconv.Itoa(sum), nil
}
