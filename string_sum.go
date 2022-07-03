package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
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

	if len(input) == 0 {
		return "", fmt.Errorf(errorEmptyInput.Error(), input)
	}
	/*var parts []string = strings.Split(input, "+")
	if len(parts) < 2 {
		parts []string = strings.Split(input, "-")
	}*/
	//if len( parts) != 2 return "", Errorf()
	var sum int = 0
	var sign string = "+"
	var currendDig string = "0"
	var operandCount int = 0
	var meetDigit bool = false
	for i := 0; i < len(input); i++ {
		if input[i] == '-' || input[i] == '+' {
			var val, _ = strconv.Atoi(sign + currendDig)
			sum += val
			currendDig = "0"
			sign = string(input[i])

			if meetDigit {
				operandCount++
			}
			meetDigit = false
		} else if unicode.IsDigit(rune(input[i])) {
			meetDigit = true
			currendDig += string(input[i])
		} else if input[i] != byte(' ') {
			_, err = strconv.Atoi(string(input[i]))
			return "", fmt.Errorf(err.Error(), input)
		}
	}
	if meetDigit {
		operandCount++
	}
	//fmt.Println(operandCount)
	if operandCount != 2 {
		return "", fmt.Errorf(errorNotTwoOperands.Error(), input)
	}
	var val, _ = strconv.Atoi(sign + currendDig)
	sum += val
	output = strconv.Itoa(sum)
	/*	a, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		b, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	*/
	//output = strconv.Itoa(a + b)
	return output, nil

}
