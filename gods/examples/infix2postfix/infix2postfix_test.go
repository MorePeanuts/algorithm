package infix2postfix

import (
	"math"
	"testing"
)

func TestString2InfixExamples(t *testing.T) {
	testCases := []struct {
		name string
		expr string
		want []string
	}{
		{
			name: "simple addition",
			expr: "3+5",
			want: []string{"3", "+", "5"},
		},
		{
			name: "multiple operators",
			expr: "3+5*2",
			want: []string{"3", "+", "5", "*", "2"},
		},
		{
			name: "with parentheses",
			expr: "(3+5)*2",
			want: []string{"(", "3", "+", "5", ")", "*", "2"},
		},
		{
			name: "with exponentiation",
			expr: "2^3",
			want: []string{"2", "^", "3"},
		},
		{
			name: "with spaces",
			expr: " 3 + 5 * 2 ",
			want: []string{"3", "+", "5", "*", "2"},
		},
		{
			name: "complex expression",
			expr: "10+3*5/(16-4)",
			want: []string{"10", "+", "3", "*", "5", "/", "(", "16", "-", "4", ")"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := String2Infix(tc.expr)
			if len(got) != len(tc.want) {
				t.Errorf("String2Infix(%q) length = %v, want %v", tc.expr, len(got), len(tc.want))
				t.Errorf("got = %v, want = %v", got, tc.want)
				return
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Errorf("String2Infix(%q) = %v, want %v", tc.expr, got, tc.want)
					break
				}
			}
		})
	}
}

func TestInfix2PostfixExamples(t *testing.T) {
	testCases := []struct {
		name    string
		infix   []string
		want    []string
		wantErr bool
	}{
		{
			name:    "simple addition",
			infix:   []string{"3", "+", "5"},
			want:    []string{"3", "5", "+"},
			wantErr: false,
		},
		{
			name:    "operator precedence",
			infix:   []string{"3", "+", "5", "*", "2"},
			want:    []string{"3", "5", "2", "*", "+"},
			wantErr: false,
		},
		{
			name:    "parentheses override precedence",
			infix:   []string{"(", "3", "+", "5", ")", "*", "2"},
			want:    []string{"3", "5", "+", "2", "*"},
			wantErr: false,
		},
		{
			name:    "exponentiation right associative",
			infix:   []string{"2", "^", "3", "^", "2"},
			want:    []string{"2", "3", "2", "^", "^"},
			wantErr: false,
		},
		{
			name:    "division and multiplication same precedence",
			infix:   []string{"10", "/", "2", "*", "5"},
			want:    []string{"10", "2", "/", "5", "*"},
			wantErr: false,
		},
		{
			name:    "complex expression",
			infix:   []string{"10", "+", "3", "*", "5", "/", "(", "16", "-", "4", ")"},
			want:    []string{"10", "3", "5", "*", "16", "4", "-", "/", "+"},
			wantErr: false,
		},
		{
			name:    "unclosed parentheses error",
			infix:   []string{"(", "3", "+", "5"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "mismatched parentheses error",
			infix:   []string{"3", "+", "5", ")"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "unknown operator error",
			infix:   []string{"3", "%", "5"},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Infix2Postfix(tc.infix)
			if (err != nil) != tc.wantErr {
				t.Errorf("Infix2Postfix() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !tc.wantErr {
				if len(got) != len(tc.want) {
					t.Errorf("Infix2Postfix() length = %v, want %v", len(got), len(tc.want))
					t.Errorf("got = %v, want = %v", got, tc.want)
					return
				}
				for i := range got {
					if got[i] != tc.want[i] {
						t.Errorf("Infix2Postfix() = %v, want %v", got, tc.want)
						break
					}
				}
			}
		})
	}
}

func TestEvalRPNExamples(t *testing.T) {
	testCases := []struct {
		name    string
		tokens  []string
		want    float64
		wantErr bool
	}{
		{
			name:    "simple addition",
			tokens:  []string{"3", "5", "+"},
			want:    8,
			wantErr: false,
		},
		{
			name:    "simple subtraction",
			tokens:  []string{"10", "4", "-"},
			want:    6,
			wantErr: false,
		},
		{
			name:    "simple multiplication",
			tokens:  []string{"3", "5", "*"},
			want:    15,
			wantErr: false,
		},
		{
			name:    "simple division",
			tokens:  []string{"10", "2", "/"},
			want:    5,
			wantErr: false,
		},
		{
			name:    "exponentiation",
			tokens:  []string{"2", "3", "^"},
			want:    8,
			wantErr: false,
		},
		{
			name:    "operator precedence example",
			tokens:  []string{"3", "5", "2", "*", "+"},
			want:    13,
			wantErr: false,
		},
		{
			name:    "parentheses example",
			tokens:  []string{"3", "5", "+", "2", "*"},
			want:    16,
			wantErr: false,
		},
		{
			name:    "complex expression",
			tokens:  []string{"10", "3", "5", "*", "16", "4", "-", "/", "+"},
			want:    11.25,
			wantErr: false,
		},
		{
			name:    "right associative exponent",
			tokens:  []string{"2", "3", "2", "^", "^"},
			want:    512,
			wantErr: false,
		},
		{
			name:    "decimal numbers",
			tokens:  []string{"3.5", "2.5", "+"},
			want:    6,
			wantErr: false,
		},
		{
			name:    "invalid expression: not enough operands",
			tokens:  []string{"3", "+"},
			want:    math.NaN(),
			wantErr: true,
		},
		{
			name:    "invalid expression: too many operands",
			tokens:  []string{"3", "5", "2", "+"},
			want:    7,
			wantErr: false,
		},
		{
			name:    "invalid expression: starts with operator",
			tokens:  []string{"+", "3", "5"},
			want:    math.NaN(),
			wantErr: true,
		},
		{
			name:    "empty expression",
			tokens:  []string{},
			want:    math.NaN(),
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := EvalRPN(tc.tokens)
			if (err != nil) != tc.wantErr {
				t.Errorf("EvalRPN() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !tc.wantErr {
				if math.Abs(got-tc.want) > 1e-9 {
					t.Errorf("EvalRPN() = %v, want %v", got, tc.want)
				}
			}
		})
	}
}

func TestFullWorkflow(t *testing.T) {
	testCases := []struct {
		name string
		expr string
		want float64
	}{
		{
			name: "simple expression",
			expr: "3+5*2",
			want: 13,
		},
		{
			name: "with parentheses",
			expr: "(3+5)*2",
			want: 16,
		},
		{
			name: "complex expression",
			expr: "10+3*5/(16-4)",
			want: 11.25,
		},
		{
			name: "exponentiation",
			expr: "2^3+5",
			want: 13,
		},
		{
			name: "nested parentheses",
			expr: "((2+3)*4)/(1+1)",
			want: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			infix := String2Infix(tc.expr)
			postfix, err := Infix2Postfix(infix)
			if err != nil {
				t.Fatalf("Infix2Postfix() error = %v", err)
			}
			got, err := EvalRPN(postfix)
			if err != nil {
				t.Fatalf("EvalRPN() error = %v", err)
			}
			if math.Abs(got-tc.want) > 1e-9 {
				t.Errorf("Full workflow for %q = %v, want %v", tc.expr, got, tc.want)
			}
		})
	}
}