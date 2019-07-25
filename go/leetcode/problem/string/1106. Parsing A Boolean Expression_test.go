package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/parsing-a-boolean-expression/

// Return the result of evaluating a given boolean expression, represented as a string.

// An expression can either be:
// "t", evaluating to True;
// "f", evaluating to False;
// "!(expr)", evaluating to the logical NOT of the inner expression expr;
// "&(expr1,expr2,...)", evaluating to the logical AND of 2 or more inner expressions expr1, expr2, ...;
// "|(expr1,expr2,...)", evaluating to the logical OR of 2 or more inner expressions expr1, expr2, ...

// ref: https://leetcode.com/problems/parsing-a-boolean-expression/discuss/323402/ChineseC%2B%2B-1106.

func parseBoolExpr(expression string) bool {
	start := 0
	return helper1106(expression, &start)
}

func helper1106(expression string, start *int) bool {
	*start = *start + 1
	switch expression[*start-1] {
	case 't':
		return true
	case 'f':
		return false
	case '!':
		return !helper1106More(expression, start, true)
	case '|':
		return helper1106More(expression, start, true)
	case '&':
		return helper1106More(expression, start, false)
	}
	return true
}

func helper1106More(expression string, start *int, flag bool) bool {
	*start = *start + 1 // skip (
	ret := helper1106(expression, start)
	for expression[*start] == ',' {
		*start = *start + 1 //skip ,
		tmp := helper1106(expression, start)
		if flag {
			ret = ret || tmp
		} else {
			ret = ret && tmp
		}
	}
	*start = *start + 1 //skip )
	return ret
}

func TestParseBoolExpr(t *testing.T) {
	tests := []struct {
		expression string
		output     bool
	}{
		{"!(f)", true},
		{"|(f,t)", true},
		{"&(t,f)", false},
		{"|(&(t,f,t),!(t))", false},
		{"&(t,t,f)", false},
		{"!(&(&(!(&(f)),&(t),|(f,f,t)),&(t),&(t,t,f)))", true},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, parseBoolExpr(ts.expression))
		})
	}
}
