package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/leagerxd/architecture-lab-2/lab2"
)

type ComputeHandler struct {
	input  io.Reader
	output io.Writer
}

func (c *ComputeHandler) Compute() error {
	expr, err := readExpression(c.input)
	if err != nil {
		return fmt.Errorf("failed to read expression: %w", err)
	}

	result, err := lab2.EvaluatePrefixExpression(expr)
	if err != nil {
		return fmt.Errorf("failed to convert expression: %w", err)
	}

	_, err = fmt.Fprintln(c.output, result)
	if err != nil {
		return fmt.Errorf("failed to write result: %w", err)
	}

	return nil
}

func readExpression(input io.Reader) (string, error) {
	buf := make([]byte, 1024)
	n, err := input.Read(buf)
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("failed to read expression: %w", err)
	}

	expr := string(buf[:n])
	return expr, nil
}

func main() {
	exprFlag := flag.String("e", "", "expression to evaluate")
	fileFlag := flag.String("f", "", "file containing the expression")
	outputFlag := flag.String("o", "", "file to write the result")

	flag.Parse()

	var input io.Reader
	if *exprFlag != "" {
		input = strings.NewReader(*exprFlag)
	} else if *fileFlag != "" {
		file, err := os.Open(*fileFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open input file: %s\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		fmt.Fprintln(os.Stderr, "either -e or -f flag must be provided")
		os.Exit(1)
	}

	var output io.Writer
	if *outputFlag != "" {
		file, err := os.Create(*outputFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to create output file: %s\n", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	} else {
		output = os.Stdout
	}

	handler := ComputeHandler{
		input:  input,
		output: output,
	}

	err := handler.Compute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
