package edit

import "testing"

func TestEncodeBase64(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple text", "hello", "aGVsbG8="},
		{"with spaces", "hello world", "aGVsbG8gd29ybGQ="},
		{"empty string", "", ""},
		{"special chars", "hello\nworld", "aGVsbG8Kd29ybGQ="},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := EditorArgs{}
			result, err := args.EncodeBase64(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertEqual(t, tt.expected, result, "")
		})
	}
}

func TestDecodeBase64(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  string
		expectErr bool
	}{
		{"simple text", "aGVsbG8=", "hello", false},
		{"with spaces", "aGVsbG8gd29ybGQ=", "hello world", false},
		{"empty string", "", "", false},
		{"invalid base64", "not-valid-base64!", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := EditorArgs{}
			result, err := args.DecodeBase64(tt.input)
			if tt.expectErr {
				if err == nil {
					t.Fatal("expected error but got none")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertEqual(t, tt.expected, result, "")
		})
	}
}

func TestEncodeForUrl(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple text", "hello", "hello"},
		{"with spaces", "hello world", "hello+world"},
		{"special chars", "a=b&c=d", "a%3Db%26c%3Dd"},
		{"empty string", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := EditorArgs{}
			result, err := args.EncodeForUrl(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertEqual(t, tt.expected, result, "")
		})
	}
}

func TestDecodeFromUrl(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  string
		expectErr bool
	}{
		{"simple text", "hello", "hello", false},
		{"with plus", "hello+world", "hello world", false},
		{"percent encoded", "a%3Db%26c%3Dd", "a=b&c=d", false},
		{"empty string", "", "", false},
		{"invalid percent", "%ZZ", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := EditorArgs{}
			result, err := args.DecodeFromUrl(tt.input)
			if tt.expectErr {
				if err == nil {
					t.Fatal("expected error but got none")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertEqual(t, tt.expected, result, "")
		})
	}
}

func TestEncodeForXml(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple text", "hello", "hello"},
		{"angle brackets", "<div>hello</div>", "&lt;div&gt;hello&lt;/div&gt;"},
		{"ampersand", "a & b", "a &amp; b"},
		{"quotes", `"hello"`, "&#34;hello&#34;"},
		{"empty string", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := EditorArgs{}
			result, err := args.EncodeForXml(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertEqual(t, tt.expected, result, "")
		})
	}
}

func TestDecodeFromXml(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple text", "hello", "hello"},
		{"angle brackets", "&lt;div&gt;hello&lt;/div&gt;", "<div>hello</div>"},
		{"ampersand", "a &amp; b", "a & b"},
		{"quotes", "&#34;hello&#34;", `"hello"`},
		{"empty string", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := EditorArgs{}
			result, err := args.DecodeFromXml(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertEqual(t, tt.expected, result, "")
		})
	}
}
