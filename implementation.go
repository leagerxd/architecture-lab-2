package lab2

import (
	"errors"
	"strconv"
	"strings"
)

// EvaluatePrefixExpression evaluates a prefix expression and returns the result.
func EvaluatePrefixExpression(input string) (int, error) {
	tokens := strings.Fields(input)
	stack := []int{}

	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]

		if isOperand(token) {
			operand, err := strconv.Atoi(token)
			if err != nil {
				return 0, errors.New("invalid expression: unable to parse operand")
			}
			stack = append(stack, operand)
		} else if isOperator(token) {
			if len(stack) < 2 {
				return 0, errors.New("invalid expression: insufficient operands")
			}

			operand1 := stack[len(stack)-1]
			operand2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			result := performOperation(token, operand1, operand2)
			stack = append(stack, result)
		} else {
			return 0, errors.New("invalid expression: unrecognized token")
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("invalid expression: too many operands")
	}

	return stack[0], nil
}

func isOperand(token string) bool {
	_, err := strconv.Atoi(token)
	return err == nil
}

func isOperator(token string) bool {
	operators := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
		"^": true,
	}

	return operators[token]
}

func performOperation(operator string, operand1, operand2 int) int {
	switch operator {
	case "+":
		return operand1 + operand2
	case "-":
		return operand1 - operand2
	case "*":
		return operand1 * operand2
	case "/":
		return operand1 / operand2
	case "^":
		return power(operand1, operand2)
	default:
		return 0
	}
}

func power(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}
