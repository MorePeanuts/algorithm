package main

import (
	"fmt"
	"math"
	"os"

	"github.com/MorePeanuts/algorithm/gods/examples/infix2postfix"
)

func calc(expr string) (float64, error) {
	infix, err := infix2postfix.String2Infix(expr)
	if err != nil {
		return math.NaN(), err
	}
	postfix, err := infix2postfix.Infix2Postfix(infix)
	if err != nil {
		return math.NaN(), err
	}
	result, err := infix2postfix.EvalRPN(postfix)
	return result, err
}

func main() {
	fmt.Println("Hello from calc!")
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "calc accepts an expression as an argument. (got %v)", len(os.Args))
		os.Exit(1)
	}
	expr := os.Args[1]
	result, err := calc(expr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error occurred while evaluating the expression %s: %v\n", expr, err)
		os.Exit(1)
	}
	fmt.Printf("%s=%v\n", expr, result)
}
