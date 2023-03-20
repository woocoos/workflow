package conv

import (
	"fmt"
	"github.com/antonmedv/expr"
	"strings"
)

//---------------------- evaluate expression ----------------------

// Eval simulate to expr.Eval, but support more types
func Eval(exp string, env any, opts ...expr.Option) (any, error) {
	opts = append(opts, expr.Env(env))
	program, err := expr.Compile(exp, opts...)
	if err != nil {
		return nil, err
	}
	return expr.Run(program, env)
}

// EvaluateExpression try evaluating an expression,if expression is a static value,vars should nil.
//
// Expressions always start with an equals sign (=). For example, = order.amount > 100.
// The text following the equal sign is the actual expression. For example, order.amount > 100 checks
// if the amount of the order is greater than 100.
//
// If the element does not start with the prefix, it is used as a static value.
// A static value is used either as a string (e.g. job type) or as a number (e.g. job retries).
// A string value must not be enclosed in quotes.
func EvaluateExpression(exp string, vars any, opts ...expr.Option) (any, error) {
	if strings.HasPrefix(exp, "=") {
		exp = strings.TrimPrefix(exp, "=")
		val, err := Eval(exp, vars, opts...)
		return val, err
	}
	return exp, nil
}

func EvaluateExpressionToInt(exp string, vars any) (int, error) {
	val, err := EvaluateExpression(exp, vars, expr.AsInt())
	if err != nil {
		return 0, fmt.Errorf("[EvaluateExpressionToInt]:%w", err)
	}
	return val.(int), nil
}
