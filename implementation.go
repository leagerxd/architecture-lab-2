package lab2

import (
	"errors"
	"strings"
)

// PrefixToPostfix converts a prefix expression to a postfix expression.
func PrefixToPostfix(input string) (string, error) {
	tokens := strings.Fields(input)
	stack := []string{}

	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]

		if isOperator(token) {
			if len(stack) < 2 {
				return "", errors.New("invalid expression: insufficient operands")
			}

			operand2 := stack[len(stack)-1] // Correct order
			operand1 := stack[len(stack)-2] // Correct order
			stack = stack[:len(stack)-2]

			result := operand2 + " " + operand1 + " " + token // Correct order
			stack = append(stack, result)
		} else {
			stack = append(stack, token)
		}
	}

	if len(stack) != 1 {
		return "", errors.New("invalid expression: too many operands")
	}

	return stack[0], nil
}

func isOperator(token string) bool {
	operators := []string{"+", "-", "*", "/", "^"}
	for _, op := range operators {
		if op == token {
			return true
		}
	}
	return false
}
