package main

import (
	"math"
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		name    string
		expr    string
		want    float64
		wantErr bool
	}{
		// Basic arithmetic
		{
			name: "simple addition",
			expr: "1+2",
			want: 3,
		},
		{
			name: "simple subtraction",
			expr: "5-3",
			want: 2,
		},
		{
			name: "simple multiplication",
			expr: "4*5",
			want: 20,
		},
		{
			name: "simple division",
			expr: "10/2",
			want: 5,
		},
		{
			name: "simple exponentiation",
			expr: "2^3",
			want: 8,
		},

		// Operator precedence
		{
			name: "multiplication before addition",
			expr: "1+5*7",
			want: 36,
		},
		{
			name: "division before subtraction",
			expr: "20-10/2",
			want: 15,
		},
		{
			name: "exponentiation before multiplication",
			expr: "3*2^3",
			want: 24,
		},
		{
			name: "mixed precedence",
			expr: "2+3*4-6/2",
			want: 11,
		},
		{
			name: "full precedence chain",
			expr: "2+3*4^2/6",
			want: 10,
		},

		// Parentheses
		{
			name: "simple parentheses",
			expr: "(1+2)*3",
			want: 9,
		},
		{
			name: "nested parentheses",
			expr: "((1+2)*3+4)*5",
			want: 65,
		},
		{
			name: "multiple parentheses groups",
			expr: "(1+2)*(3+4)",
			want: 21,
		},
		{
			name: "parentheses changing precedence",
			expr: "(2+3)*4",
			want: 20,
		},
		{
			name: "deeply nested parentheses",
			expr: "((((1+2)+3)+4)+5)",
			want: 15,
		},
		{
			name: "parentheses with exponentiation",
			expr: "(2+3)^2",
			want: 25,
		},

		// Left associativity
		{
			name: "left associative addition",
			expr: "1+2+3",
			want: 6,
		},
		{
			name: "left associative subtraction",
			expr: "10-5-2",
			want: 3,
		},
		{
			name: "left associative multiplication",
			expr: "2*3*4",
			want: 24,
		},
		{
			name: "left associative division",
			expr: "100/10/2",
			want: 5,
		},
		{
			name: "mixed left associative",
			expr: "100-20-10+5",
			want: 75,
		},

		// Right associativity (exponentiation)
		{
			name: "right associative exponentiation",
			expr: "2^3^2",
			want: 512, // 2^(3^2) = 2^9 = 512, not (2^3)^2 = 64
		},
		{
			name: "exponentiation with parentheses to force left",
			expr: "(2^3)^2",
			want: 64,
		},

		// Decimal numbers
		{
			name: "decimal addition",
			expr: "1.5+2.5",
			want: 4,
		},
		{
			name: "decimal multiplication",
			expr: "2.5*4",
			want: 10,
		},
		{
			name: "decimal division",
			expr: "10/2.5",
			want: 4,
		},
		{
			name: "decimal exponentiation",
			expr: "4^0.5",
			want: 2,
		},
		{
			name: "complex decimal expression",
			expr: "10.5+3.5*2-8.0/4",
			want: 15.5,
		},

		// Negative results
		{
			name: "subtraction with negative result",
			expr: "3-5",
			want: -2,
		},
		{
			name: "complex expression with negative result",
			expr: "5-3*4",
			want: -7,
		},

		// Zero handling
		{
			name: "addition with zero",
			expr: "0+5",
			want: 5,
		},
		{
			name: "multiplication with zero",
			expr: "5*0",
			want: 0,
		},
		{
			name: "zero exponent",
			expr: "5^0",
			want: 1,
		},
		{
			name: "exponent of zero",
			expr: "0^5",
			want: 0,
		},
		{
			name: "division resulting in zero",
			expr: "0/5",
			want: 0,
		},

		// Large numbers
		{
			name: "large multiplication",
			expr: "1000000*1000000",
			want: 1e12,
		},
		{
			name: "large exponentiation",
			expr: "2^20",
			want: 1048576,
		},

		// Edge cases
		{
			name: "single number",
			expr: "42",
			want: 42,
		},
		{
			name: "single decimal",
			expr: "3.14159",
			want: 3.14159,
		},

		// Whitespace handling (String2Infix uses TrimSpace but not internal spaces)
		{
			name: "expression with only leading/trailing spaces",
			expr: " 1+2*3 ",
			want: 7,
		},

		// Complex combinations
		{
			name: "complex expression 1",
			expr: "(2+3*4)^2-10/2",
			want: 191,
		},
		{
			name: "complex expression 2",
			expr: "100-((2+3)*4)^2/2",
			want: -100,
		},
		{
			name: "complex expression 3",
			expr: "2^3+4*5-(6+7)/13",
			want: 27,
		},
		{
			name: "complex expression with multiple exponents",
			expr: "2^2+3^2*4^2",
			want: 148,
		},

		// Error cases
		{
			name:    "mismatched parentheses - missing closing",
			expr:    "(1+2",
			wantErr: true,
		},
		{
			name:    "mismatched parentheses - missing opening",
			expr:    "1+2)",
			wantErr: true,
		},
		{
			name:    "unknown operator",
			expr:    "1$2",
			wantErr: true,
		},
		{
			name:    "deeply nested mismatched",
			expr:    "((1+2)+3",
			wantErr: true,
		},
		{
			name:    "multiplication with unary minus with out parentheses",
			expr:    "5*-3",
			wantErr: true,
		},
		{
			name:    "division with unary minus with out parentheses",
			expr:    "5/-2",
			wantErr: true,
		},
		{
			name:    "exponentiation with unary minus with out parentheses",
			expr:    "2^-3",
			wantErr: true,
		},

		// Division by zero (Note: EvalRPN doesn't check for division by zero, it will return Inf)
		{
			name: "division by zero",
			expr: "5/0",
			want: math.Inf(1),
		},

		// Unary operators (+ or - without preceding number)
		{
			name: "unary minus at start",
			expr: "-5",
			want: -5,
		},
		{
			name: "unary plus at start",
			expr: "+5",
			want: 5,
		},
		{
			name: "unary minus followed by addition",
			expr: "-5+3",
			want: -2,
		},
		{
			name: "unary plus followed by subtraction",
			expr: "+5-3",
			want: 2,
		},
		{
			name: "multiplication with unary minus",
			expr: "5*(-3)",
			want: -15,
		},
		{
			name: "division with unary minus",
			expr: "5/(-2)",
			want: -2.5,
		},
		{
			name: "exponentiation with unary minus",
			expr: "2^(-3)",
			want: 0.125,
		},
		{
			name: "unary minus inside parentheses",
			expr: "(-5+3)*2",
			want: -4,
		},
		{
			name: "unary plus inside parentheses",
			expr: "(+5-3)*2",
			want: 4,
		},
		{
			name: "unary minus after operator",
			expr: "2*(-3+5)",
			want: 4,
		},
		{
			name: "unary plus after operator",
			expr: "2*(+3+5)",
			want: 16,
		},
		{
			name: "double unary minus",
			expr: "-(-5)",
			want: 5,
		},
		{
			name: "unary minus with unary plus",
			expr: "-(+5)",
			want: -5,
		},
		{
			name: "unary plus with unary minus",
			expr: "+(-5)",
			want: -5,
		},
		{
			name: "double unary plus",
			expr: "+(+5)",
			want: 5,
		},
		{
			name: "complex expression with unary operators",
			expr: "-2^2+5*(-3+10)",
			want: 31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calc(tt.expr)
			if (err != nil) != tt.wantErr {
				t.Errorf("calc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			// Handle NaN and Inf comparisons specially
			if math.IsNaN(tt.want) {
				if !math.IsNaN(got) {
					t.Errorf("calc() = %v, want %v", got, tt.want)
				}
			} else if math.IsInf(tt.want, 1) {
				if !math.IsInf(got, 1) {
					t.Errorf("calc() = %v, want %v", got, tt.want)
				}
			} else if math.IsInf(tt.want, -1) {
				if !math.IsInf(got, -1) {
					t.Errorf("calc() = %v, want %v", got, tt.want)
				}
			} else {
				// Use a small epsilon for floating point comparisons
				const epsilon = 1e-9
				if math.Abs(got-tt.want) > epsilon {
					t.Errorf("calc() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

