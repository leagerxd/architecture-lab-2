package lab2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluatePrefixExpression(t *testing.T) {
	input := "+ 2 3"
	expected := 5

	result, err := EvaluatePrefixExpression(input)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestEvaluatePrefixExpression_ComplexExpression(t *testing.T) {
	input := "* + 3 + 4 2 5"
	expected := 45

	result, err := EvaluatePrefixExpression(input)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestEvaluatePrefixExpression_InvalidExpression(t *testing.T) {
	input := "+ 2"
	expected := 0

	result, err := EvaluatePrefixExpression(input)
	assert.NotNil(t, err)
	assert.Equal(t, expected, result)
}
