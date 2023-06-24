package main

import (
	"flag"
	"fmt"

	lab2 "github.com/leagerxd/architecture-lab-2/lab2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	// TODO: Add other flags support for input and output configuration.
)

func main() {
	flag.Parse()

	result, err := lab2.PrefixToPostfix(*inputExpression)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}
