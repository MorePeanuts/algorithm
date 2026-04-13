package main

import (
	"fmt"
	"math"

	"github.com/MorePeanuts/algorithm/gods/examples/infix2postfix"
)

func calc(expr string) (float64, error) {
	infix := infix2postfix.String2Infix(expr)
	postfix, err := infix2postfix.Infix2Postfix(infix)
	if err != nil {
		return math.NaN(), err
	}
	result, err := infix2postfix.EvalRPN(postfix)
	return result, err
}

func main() {
	fmt.Println("Hello from calc!")
	expr := "5*-3"
	result, err := calc(expr)
	if err != nil {
		fmt.Printf("error occurred while evaluating the expression %s: %v\n", expr, err)
		return
	}
	fmt.Printf("%s=%v\n", expr, result)
}
