package strings

import "testing"

func TestReverseString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"a", "a"},
		{"hello", "olleh"},
		{"ილია", "აილი"},
	}

	for _, tc := range testCases {
		actual := ReverseString(tc.input)
		if actual != tc.expected {
			t.Errorf("reverseString(%q) = %q; expected %q", tc.input, actual, tc.expected)
		}
	}
}
