package strings

import "testing"

func TestCountUnicodeChars(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"a", 1},
		{"abc", 3},
		{"ილიაა", 5},
	}

	for _, tc := range testCases {
		actual := CountUnicodeChars(tc.input)
		if actual != tc.expected {
			t.Errorf("countUnicodeChars(%q) = %d; expected %d", tc.input, actual, tc.expected)
		}
	}
}
