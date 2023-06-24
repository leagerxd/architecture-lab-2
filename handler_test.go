package lab2_test

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	lab2 "github.com/leagerxd/architecture-lab-2/lab2"
)

func TestComputeHandler_Compute(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedError  error
		expectedOutput string
	}{
		{
			name:           "ValidExpression",
			input:          "+ 2 3\n",
			expectedError:  nil,
			expectedOutput: "Result: 5\n",
		},
		{
			name:           "InvalidExpression",
			input:          "+ 2\n",
			expectedError:  errors.New("invalid expression: insufficient operands"),
			expectedOutput: "",
		},
		{
			name:           "InvalidSyntax",
			input:          "+ 2a 3\n",
			expectedError:  errors.New("invalid expression: unrecognized token"),
			expectedOutput: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			input := strings.NewReader(test.input)
			output := bytes.Buffer{}

			handler := lab2.NewComputeHandler(input, &output)
			err := handler.Compute()

			if err != nil && test.expectedError == nil {
				t.Errorf("unexpected error: %v", err)
			}

			if err == nil && test.expectedError != nil {
				t.Errorf("expected error: %v, but got none", test.expectedError)
			}

			if err != nil && test.expectedError != nil && err.Error() != test.expectedError.Error() {
				t.Errorf("expected error: %v, but got: %v", test.expectedError, err)
			}

			outputString := output.String()
			if outputString != test.expectedOutput {
				t.Errorf("expected output: %q, but got: %q", test.expectedOutput, outputString)
			}
		})
	}
}
