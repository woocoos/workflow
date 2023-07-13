package bpmn

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Expression struct {
	BaseElementWithMixedContent
}

type FormalExpression struct {
	Expression
	Language           string `xml:"language,omitempty,attr"`
	EvaluatesToTypeRef string `xml:"evaluatesToTypeRef,attr"`
}

// IsFormal check expression is formal expression
func IsFormal(exp string) bool {
	return strings.HasPrefix(exp, "=")
}

// ConditionExpression is camunda extension
type ConditionExpression struct {
	FormalExpression
	Type     string `xml:"type,attr"`
	Resource string `xml:"resource,attr"`
}

// Evaluate condition expression. Camunda FEEL format is no like golang expression.
//
// Supported operators:
//
//	`=`: FEEL equality, golang expression `==`.. note support in quote string like `="a=b"`
func (c ConditionExpression) Evaluate(input map[string]any) (bool, error) {
	if Convert == nil {
		return false, errors.New("evaluate converter not set")
	}
	return Convert.Match(c.Content, input)
}

var (
	ErrExpressionResultNotString = errors.New("expression result is not string")
	ErrExpressionResultNotInt    = errors.New("expression result is not int")
	ErrExpressionResultNotBool   = errors.New("condition expression result is not bool")
)

func EvaluateExpressionToString(exp string, vars Mappings) (string, error) {
	val, err := Convert.Eval(exp, vars)
	if err != nil {
		return "", err
	}
	s, ok := val.(string)
	if !ok {
		return "", ErrExpressionResultNotString
	}
	return s, nil
}

//--------------------

var ErrISO8601DurationFormat = errors.New("input string is of incorrect format")

// EvaluateExpressionToDuration evaluate iso8601 format to time.Duration
func EvaluateExpressionToDuration(exp string) (du time.Duration, err error) {
	if strings.HasPrefix(exp, "=") {
		err = ErrISO8601DurationFormat
		return
	}
	return ParseDuration(exp)
}

var pattern = regexp.MustCompile(`^P(?:(\d+)Y)?(?:(\d+)M)?(?:(\d+)D)?T(?:(\d+)H)?(?:(\d+)M)?(?:(\d+(?:.\d+)?)S)?$`)

// ParseDuration parses an ISO 8601 string representing a duration,
// and returns the resultant golang time.Duration instance.
func ParseDuration(isoDuration string) (time.Duration, error) {
	matches := pattern.FindStringSubmatch(isoDuration)
	if matches == nil {
		return 0, ErrISO8601DurationFormat
	}

	var seconds int64

	//skipping years and months

	//days
	if matches[3] != "" {
		f, err := strconv.ParseInt(matches[3], 10, 32)
		if err != nil {
			return 0, err
		}

		seconds += f * 24 * 60 * 60
	}
	//hours
	if matches[4] != "" {
		f, err := strconv.ParseInt(matches[4], 10, 32)
		if err != nil {
			return 0, err
		}

		seconds += f * 60 * 60
	}
	//minutes
	if matches[5] != "" {
		f, err := strconv.ParseInt(matches[5], 10, 32)
		if err != nil {
			return 0, err
		}

		seconds += f * 60
	}
	//seconds & milliseconds
	if matches[6] != "" {
		f, err := strconv.ParseInt(matches[6], 10, 32)
		if err != nil {
			return 0, err
		}

		seconds += f
	}
	return time.Duration(seconds) * time.Second, nil
}

// FormatDuration returns an ISO 8601 duration string.
func FormatDuration(dur time.Duration) string {
	return "PT" + strings.ToUpper(dur.Truncate(time.Millisecond).String())
}

func AsStr(val interface{}) interface{} {
	return fmt.Sprintf("%v", val)
}

func AsInt(in interface{}) int64 {
	switch i := in.(type) {
	case float64:
		return int64(i)
	case float32:
		return int64(i)
	case int64:
		return i
	case int32:
		return int64(i)
	case int16:
		return int64(i)
	case int8:
		return int64(i)
	case int:
		return int64(i)
	case uint64:
		return int64(i)
	case uint32:
		return int64(i)
	case uint16:
		return int64(i)
	case uint8:
		return int64(i)
	case uint:
		return int64(i)
	case string:
		inAsInt, err := strconv.ParseInt(i, 10, 64)
		if err == nil {
			return inAsInt
		}
		panic(err)
	}
	panic(fmt.Sprintf("asInt() not supported on %v %v", reflect.TypeOf(in), in))
}

func AsFloat(in interface{}) float64 {
	switch i := in.(type) {
	case float64:
		return i
	case float32:
		return float64(i)
	case int64:
		return float64(i)
	case int32:
		return float64(i)
	case int16:
		return float64(i)
	case int8:
		return float64(i)
	case int:
		return float64(i)
	case uint64:
		return float64(i)
	case uint32:
		return float64(i)
	case uint16:
		return float64(i)
	case uint8:
		return float64(i)
	case uint:
		return float64(i)
	case string:
		inAsFloat, err := strconv.ParseFloat(i, 64)
		if err == nil {
			return inAsFloat
		}
		panic(err)
	}
	panic(fmt.Sprintf("asFloat() not supported on %v %v", reflect.TypeOf(in), in))
}
