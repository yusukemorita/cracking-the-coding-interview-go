package main

import "testing"

func TestIsOneEditAway1(t *testing.T) {
	testCases := map[string]struct {
		expected bool
		s1, s2   string
	}{
		"returns false when more than one edit away":      {false, "hel", "hello"},
		"returns true when same string":                   {true, "hello", "hello"},
		"returns true when first character is missing":    {true, "hello", "ello"},
		"returns true when middle character is missing":   {true, "hello", "hllo"},
		"returns true when last character is missing":     {true, "hello", "hell"},
		"returns true when first character is different":  {true, "bello", "hello"},
		"returns true when middle character is different": {true, "heclo", "hello"},
		"returns true when last character is different":   {true, "hello", "hella"},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := IsOneEditAway1(testCase.s1, testCase.s2)
			if got != testCase.expected {
				t.Errorf("expected %v, got %v", testCase.expected, got)
			}
		})
	}
}

func TestIsOneEditAway2(t *testing.T) {
	testCases := map[string]struct {
		expected bool
		s1, s2   string
	}{
		"returns false when more than one edit away":      {false, "hel", "hello"},
		"returns true when same string":                   {true, "hello", "hello"},
		"returns true when first character is missing":    {true, "hello", "ello"},
		"returns true when middle character is missing":   {true, "hello", "hllo"},
		"returns true when last character is missing":     {true, "hello", "hell"},
		"returns true when first character is different":  {true, "bello", "hello"},
		"returns true when middle character is different": {true, "heclo", "hello"},
		"returns true when last character is different":   {true, "hello", "hella"},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := IsOneEditAway2(testCase.s1, testCase.s2)
			if got != testCase.expected {
				t.Errorf("expected %v, got %v", testCase.expected, got)
			}
		})
	}
}
