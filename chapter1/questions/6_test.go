package main

import "testing"

func TestCompressString(t *testing.T) {
	testCases := map[string]struct{ input, expected string }{
		"returns original when blank string": {
			input:    "",
			expected: "",
		},
		"returns original when compressing does not reduce size": {
			input:    "aaabc",
			expected: "aaabc",
		},
		"returns compressed when compressing reduces size": {
			input:    "aaaaabc",
			expected: "a5b1c1",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			output := CompressString(tc.input)
			if output != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, output)
			}
		})
	}

}
