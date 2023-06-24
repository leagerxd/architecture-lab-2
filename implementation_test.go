package lab2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixToPostfix(t *testing.T) {
	input := "+ 2 3"
	expected := "2 3 +"

	result, err := PrefixToPostfix(input)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestPrefixToPostfix_ComplexExpression(t *testing.T) {
	input := "+ * 3 + 4 2 5"
	expected := "4 2 + 3 * 5 +"

	result, err := PrefixToPostfix(input)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestPrefixToPostfix_InvalidExpression(t *testing.T) {
	input := "+ 2"
	expected := ""

	result, err := PrefixToPostfix(input)
	assert.NotNil(t, err)
	assert.Equal(t, expected, result)
}
