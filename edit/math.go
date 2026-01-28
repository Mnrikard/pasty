package edit

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func (e *EditorArgs) SolveMath(input string) (string, error) {
	input = strings.TrimSpace(input)

	validChars := regexp.MustCompile(`^[\d\s\+\-\*\/\(\)\=\.]+$`)
	if !validChars.MatchString(input) {
		return input, nil
	}

	expr := strings.Split(input, "=")[0]
	expr = strings.TrimSpace(expr)

	result, err := evaluateExpression(expr)
	if err != nil {
		return input, nil
	}

	if result == float64(int64(result)) {
		return fmt.Sprintf("%d", int64(result)), nil
	}
	return fmt.Sprintf("%g", result), nil
}

func evaluateExpression(expr string) (float64, error) {
	expr = strings.ReplaceAll(expr, " ", "")
	return parseAddSub(expr)
}

func parseAddSub(expr string) (float64, error) {
	result, rest, err := parseMulDiv(expr)
	if err != nil {
		return 0, err
	}

	for len(rest) > 0 {
		op := rest[0]
		if op != '+' && op != '-' {
			break
		}
		next, newRest, err := parseMulDiv(rest[1:])
		if err != nil {
			return 0, err
		}
		if op == '+' {
			result += next
		} else {
			result -= next
		}
		rest = newRest
	}
	return result, nil
}

func parseMulDiv(expr string) (float64, string, error) {
	result, rest, err := parseFactor(expr)
	if err != nil {
		return 0, "", err
	}

	for len(rest) > 0 {
		op := rest[0]
		if op != '*' && op != '/' {
			break
		}
		next, newRest, err := parseFactor(rest[1:])
		if err != nil {
			return 0, "", err
		}
		if op == '*' {
			result *= next
		} else {
			if next == 0 {
				return 0, "", fmt.Errorf("division by zero")
			}
			result /= next
		}
		rest = newRest
	}
	return result, rest, nil
}

func parseFactor(expr string) (float64, string, error) {
	if len(expr) == 0 {
		return 0, "", fmt.Errorf("unexpected end of expression")
	}

	if expr[0] == '(' {
		depth := 1
		i := 1
		for i < len(expr) && depth > 0 {
			if expr[i] == '(' {
				depth++
			} else if expr[i] == ')' {
				depth--
			}
			i++
		}
		if depth != 0 {
			return 0, "", fmt.Errorf("mismatched parentheses")
		}
		result, err := parseAddSub(expr[1 : i-1])
		if err != nil {
			return 0, "", err
		}
		return result, expr[i:], nil
	}

	if expr[0] == '-' {
		result, rest, err := parseFactor(expr[1:])
		if err != nil {
			return 0, "", err
		}
		return -result, rest, nil
	}

	i := 0
	for i < len(expr) && (isDigit(expr[i]) || expr[i] == '.') {
		i++
	}
	if i == 0 {
		return 0, "", fmt.Errorf("expected number")
	}
	num, err := strconv.ParseFloat(expr[:i], 64)
	if err != nil {
		return 0, "", err
	}
	return num, expr[i:], nil
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
