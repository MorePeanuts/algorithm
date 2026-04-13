package infix2postfix

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"

	"github.com/MorePeanuts/algorithm/gods/stacks/arraystack"
)

type operator struct {
	symbol    string
	priority  int
	leftFirst bool
	calculate func(float64, float64) float64
}

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

func EvalRPN(tokens []string) (float64, error) {
	stack := arraystack.New[float64]()
	for _, token := range tokens {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			stack.Push(num)
		} else {
			num2, _ := stack.Pop()
			num1, ok := stack.Pop()
			if !ok {
				return math.NaN(), fmt.Errorf("invalid suffix expression: %v", tokens)
			}
			op := operatorMap[token]
			res := op.calculate(num1, num2)
			stack.Push(res)
		}
	}
	result, _ := stack.Pop()
	return result, nil
}

func Infix2Postfix(tokens []string) (result []string, err error) {
	stack := arraystack.New[string]()
	for _, token := range tokens {
		if _, err := strconv.ParseFloat(token, 64); err == nil {
			result = append(result, token)
		} else {
			switch token {
			case ")":
				for elem, ok := stack.Pop(); elem != "("; elem, ok = stack.Pop() {
					if !ok {
						return nil, fmt.Errorf("infix expression parentheses do not match: %v", tokens)
					}
					result = append(result, elem)
				}
			case "(":
				stack.Push(token)
			default:
				currOp, ok := operatorMap[token]
				if !ok {
					return nil, fmt.Errorf("unknown operator: %s", token)
				}
				for {
					elem, ok := stack.Peek()
					if topOp := operatorMap[elem]; !ok || elem == "(" || topOp.priority < currOp.priority || topOp.priority == currOp.priority && !currOp.leftFirst {
						// currOp participates in the calculation first
						break
					}
					// topOp participates in the calculation first
					elem, _ = stack.Pop()
					result = append(result, elem)
				}
				stack.Push(token)
			}
		}
	}
	for elem, ok := stack.Pop(); ok; elem, ok = stack.Pop() {
		if elem == "(" {
			return nil, fmt.Errorf("unclosed parentheses: %v", tokens)
		}
		result = append(result, elem)
	}
	return result, nil
}

// String2Infix converts infix string into a slice that separates operators and operands.
// This function can only handle single-character operators
func String2Infix(expr string) []string {
	cleanExpr := strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, expr)
	opPos := make([]int, 0)
	for i, r := range cleanExpr {
		if _, ok := operatorMap[string(r)]; ok || r == '(' || r == ')' {
			opPos = append(opPos, i)
		}
	}
	infix := make([]string, 0)
	lastPos := 0
	for _, pos := range opPos {
		num := cleanExpr[lastPos:pos]
		// if len(num) > 0 {
		// 	infix = append(infix, num)
		// }
		op := cleanExpr[pos : pos+1]
		if len(num) > 0 {
			infix = append(infix, num)
		} else {
			if (op == "+" || op == "-") && (pos-1 < 0 || cleanExpr[pos-1:pos] == "(") {
				infix = append(infix, "0")
			}
		}
		infix = append(infix, op)
		lastPos = pos + 1
	}
	if lastPos != len(cleanExpr) {
		infix = append(infix, cleanExpr[lastPos:])
	}
	return infix
}
