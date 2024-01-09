package main

import "testing"

func TestIsPalindromePermutation1(t *testing.T) {
	t.Run("returns true when is palindrome permutation without spaces", func(t *testing.T) {
		got := IsPalindromePermutation1("abcba")
		if got != true {
			t.Errorf("expected true")
		}
	})

	t.Run("returns true when is palindrome permutation with spaces", func(t *testing.T) {
		got := IsPalindromePermutation1("tact coa")
		if got != true {
			t.Errorf("expected true")
		}
	})

	t.Run("returns false when is not palindrome permutation", func(t *testing.T) {
		got := IsPalindromePermutation1("taco bell")
		if got != false {
			t.Errorf("expected false")
		}
	})
}

func TestIsPalindromePermutation2(t *testing.T) {
	t.Run("returns true when is palindrome permutation without spaces", func(t *testing.T) {
		got := IsPalindromePermutation2("abcba")
		if got != true {
			t.Errorf("expected true")
		}
	})

	t.Run("returns true when is palindrome permutation with spaces", func(t *testing.T) {
		got := IsPalindromePermutation2("tact coa")
		if got != true {
			t.Errorf("expected true")
		}
	})

	t.Run("returns false when is not palindrome permutation", func(t *testing.T) {
		got := IsPalindromePermutation2("taco bell")
		if got != false {
			t.Errorf("expected false")
		}
	})
}
