package edit

import "testing"

func TestSolveMath(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple addition", "2 + 3", "5"},
		{"simple subtraction", "10 - 4", "6"},
		{"simple multiplication", "3 * 4", "12"},
		{"simple division", "20 / 5", "4"},
		{"order of operations", "2 + 3 * 4", "14"},
		{"parentheses", "(2 + 3) * 4", "20"},
		{"nested parentheses", "((2 + 3) * 2) + 1", "11"},
		{"negative result", "5 - 10", "-5"},
		{"decimal result", "5 / 2", "2.5"},
		{"with equals sign", "2 + 2 =", "4"},
		{"complex expression", "10 + 20 / 4 - 3", "12"},
		{"spaces around operators", "  10  +  5  ", "15"},
		{"leading negative", "-5 + 10", "5"},
		{"decimal input", "2.5 + 2.5", "5"},
		{"whole number from decimals", "2.5 * 2", "5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := EditorArgs{}
			result, err := args.SolveMath(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertEqual(t, tt.expected, result, "")
		})
	}
}

func TestSolveMath_InvalidInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"contains letters", "2 + abc"},
		{"contains special chars", "2 + 3 @ 4"},
		{"contains hash", "2 # 3"},
		{"contains percent", "50%"},
		{"contains dollar", "$100 + 50"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := EditorArgs{}
			result, err := args.SolveMath(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertEqual(t, tt.input, result, "invalid input should be returned unchanged")
		})
	}
}

func TestSolveMath_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"just a number", "42", "42"},
		{"negative number", "-42", "-42"},
		{"zero", "0", "0"},
		{"division resulting in integer", "10 / 2", "5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := EditorArgs{}
			result, err := args.SolveMath(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertEqual(t, tt.expected, result, "")
		})
	}
}

func TestEvaluateExpression_Errors(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"division by zero", "10 / 0"},
		{"mismatched parentheses open", "(2 + 3"},
		{"unexpected end", "2 +"},
		{"expected number", "2 + * 3"},
		{"empty expression", ""},
		{"only operator", "+"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := evaluateExpression(tt.input)
			if err == nil {
				t.Errorf("expected error for %q, but got nil", tt.input)
			}
		})
	}
}
