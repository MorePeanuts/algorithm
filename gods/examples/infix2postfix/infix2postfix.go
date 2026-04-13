// Package infix2postfix provides functions to convert infix expressions to postfix (Reverse Polish Notation)
// and evaluate postfix expressions. It supports basic arithmetic operations, exponentiation, and parentheses.
package infix2postfix

import (
	"fmt"
	"math"
	"strconv"

	"github.com/MorePeanuts/algorithm/gods/stacks/arraystack"
)

// operator represents a mathematical operator with its properties.
type operator struct {
	symbol    string                          // Operator symbol, e.g., "+", "-", "*", "/", "^"
	priority  int                             // Operator precedence (higher number means higher precedence
	leftFirst bool                            // Whether left associativity: true for left-associative, false for right-associative
	calculate func(float64, float64) float64 // Function to perform the calculation
}

// operatorMap maps operator symbols to their corresponding operator definitions.
var operatorMap = map[string]operator{
	"+": {"+", 5, true, func(l, r float64) float64 {
		return l + r
	}},
	"-": {"-", 5, true, func(l, r float64) float64 {
		return l - r
	}},
	"*": {"*", 7, true, func(l, r float64) float64 {
		return l * r
	}},
	"/": {"/", 7, true, func(l, r float64) float64 {
		return l / r
	}},
	"^": {"^", 9, false, func(l, r float64) float64 {
		return math.Pow(l, r)
	}},
}

// EvalRPN evaluates a Reverse Polish Notation (postfix) expression.
// It takes a slice of tokens where each token is either a number or an operator.
// Returns the result of the evaluation or an error if the expression is invalid.
func EvalRPN(tokens []string) (float64, error) {
	stack := arraystack.New[float64]()
	for _, token := range tokens {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			// Push numbers onto the stack
			stack.Push(num)
		} else {
			// Pop two operands for the operator
			num2, _ := stack.Pop()
			num1, ok := stack.Pop()
			if !ok {
				return math.NaN(), fmt.Errorf("invalid suffix expression: %v", tokens)
			}
			// Apply the operator and push the result back
			op := operatorMap[token]
			res := op.calculate(num1, num2)
			stack.Push(res)
		}
	}
	result, _ := stack.Pop()
	return result, nil
}

// Infix2Postfix converts an infix expression to postfix (Reverse Polish Notation) using the Shunting-yard algorithm.
// It takes a slice of tokens (numbers, operators, parentheses) and returns the postfix form.
// Returns an error if there are mismatched parentheses or unknown operators.
func Infix2Postfix(tokens []string) (result []string, err error) {
	stack := arraystack.New[string]()
	for _, token := range tokens {
		if _, err := strconv.ParseFloat(token, 64); err == nil {
			// Operands are added directly to the output
			result = append(result, token)
		} else {
			switch token {
			case ")":
				// Pop until matching opening parenthesis is found
				for elem, ok := stack.Pop(); elem != "("; elem, ok = stack.Pop() {
					if !ok {
						return nil, fmt.Errorf("infix expression parentheses do not match: %v", tokens)
					}
					result = append(result, elem)
				}
			case "(":
				// Push opening parenthesis onto the stack
				stack.Push(token)
			default:
				// Handle operators
				currOp, ok := operatorMap[token]
				if !ok {
					return nil, fmt.Errorf("unknown operator: %s", token)
				}
				// Pop operators with higher or equal precedence from the stack
				for {
					elem, ok := stack.Peek()
					if topOp := operatorMap[elem]; !ok || elem == "(" || topOp.priority < currOp.priority || topOp.priority == currOp.priority && !currOp.leftFirst {
						// Current operator has higher precedence or is right-associative with equal precedence
						break
					}
					// Top operator has higher precedence, pop it to output
					elem, _ = stack.Pop()
					result = append(result, elem)
				}
				// Push current operator onto the stack
				stack.Push(token)
			}
		}
	}
	// Pop remaining operators from the stack
	for elem, ok := stack.Pop(); ok; elem, ok = stack.Pop() {
		if elem == "(" {
			return nil, fmt.Errorf("unclosed parentheses: %v", tokens)
		}
		result = append(result, elem)
	}
	return result, nil
}

// String2Infix converts an infix expression string into a slice that separates operators and operands.
// It handles multi-digit numbers, decimal points, and skips whitespace.
// This function can only handle single-character operators.
// Returns an error if invalid numbers or operators are encountered.
func String2Infix(expr string) ([]string, error) {
	var res []string
	var l, r int
	for r < len(expr) {
		currByte := expr[r]
		if currByte >= '0' && currByte <= '9' || currByte == '.' {
			// Continue building a number token
			r++
		} else {
			// Push any accumulated number token
			if l != r {
				_, err := strconv.ParseFloat(expr[l:r], 64)
				if err != nil {
					return nil, fmt.Errorf("invalid expression %s: %v", expr, err)
				}
				res = append(res, expr[l:r])
			}
			// Process operator or parenthesis
			if currByte != ' ' {
				op, ok := operatorMap[expr[r:r+1]]
				if !ok && currByte != '(' && currByte != ')' {
					return nil, fmt.Errorf("invalid operator %s", expr[r:r+1])
				}
				// Handle unary plus/minus by inserting a leading zero
				if (op.symbol == "+" || op.symbol == "-") && leftNonSpace(expr, r) == '(' {
					res = append(res, "0")
				}
				res = append(res, expr[r:r+1])
			}
			l, r = r+1, r+1
		}
	}
	// Push any remaining number token
	if l != len(expr) {
		res = append(res, expr[l:r])
	}
	return res, nil
}

// leftNonSpace finds the first non-space character to the left of the given index.
// Returns '(' if no non-space character is found (indicating start of expression).
func leftNonSpace(s string, idx int) byte {
	for i := idx - 1; i >= 0; i-- {
		if s[i] != ' ' {
			return s[i]
		}
	}
	return '('
}
