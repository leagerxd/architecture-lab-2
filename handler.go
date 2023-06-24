package lab2

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type ComputeHandler struct {
	input  io.Reader
	output io.Writer
}

func NewComputeHandler(input io.Reader, output io.Writer) *ComputeHandler {
	return &ComputeHandler{
		input:  input,
		output: output,
	}
}

func (ch *ComputeHandler) Compute() error {
	scanner := bufio.NewScanner(ch.input)
	for scanner.Scan() {
		expression := scanner.Text()
		result, err := CheckPrefixExpression(expression)
		if err != nil {
			return err
		}

		output := fmt.Sprintf("Result: %d\n", result)
		_, err = ch.output.Write([]byte(output))
		if err != nil {
			return err
		}
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	return nil
}

func CheckPrefixExpression(input string) (int, error) {
	tokens := strings.Fields(input)
	stack := []int{}

	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]

		if ifOperand(token) {
			operand, err := strconv.Atoi(token)
			if err != nil {
				return 0, errors.New("invalid expression: unable to parse operand")
			}
			stack = append(stack, operand)
		} else if ifOperator(token) {
			if len(stack) < 2 {
				return 0, errors.New("invalid expression: insufficient operands")
			}

			operand1 := stack[len(stack)-1]
			operand2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			result := PerformOperation(token, operand1, operand2)
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

func ifOperand(token string) bool {
	_, err := strconv.Atoi(token)
	return err == nil
}

func ifOperator(token string) bool {
	operators := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
		"^": true,
	}

	return operators[token]
}

func PerformOperation(operator string, operand1, operand2 int) int {
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
		return Power(operand1, operand2)
	default:
		return 0
	}
}

func Power(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}
