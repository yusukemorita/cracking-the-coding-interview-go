package main

import "testing"

func TestIsPermutation1(t *testing.T) {
	t.Run("returns true when two strings are permutations", func(t *testing.T) {
		got := IsPermutation1("abcc", "cbac")
		want := true

		if got != want {
			t.Errorf("wanted: %v, got: %v", want, got)
		}
	})

	t.Run("returns false when two strings are not permutations", func(t *testing.T) {
		got := IsPermutation1("abcc", "defg")
		want := false

		if got != want {
			t.Errorf("wanted: %v, got: %v", want, got)
		}
	})

	t.Run("returns false when one strings characters are a subset of the other strings characters", func(t *testing.T) {
		got := IsPermutation1("abcc", "abccd")
		want := false

		if got != want {
			t.Errorf("wanted: %v, got: %v", want, got)
		}
	})
}

func TestIsPermutation2(t *testing.T) {
	t.Run("returns true when two strings are permutations", func(t *testing.T) {
		got := IsPermutation2("abcc", "cbac")
		want := true

		if got != want {
			t.Errorf("wanted: %v, got: %v", want, got)
		}
	})

	t.Run("returns false when two strings are not permutations", func(t *testing.T) {
		got := IsPermutation2("abcc", "defg")
		want := false

		if got != want {
			t.Errorf("wanted: %v, got: %v", want, got)
		}
	})

	t.Run("returns false when one strings characters are a subset of the other strings characters", func(t *testing.T) {
		got := IsPermutation2("abcc", "abccd")
		want := false

		if got != want {
			t.Errorf("wanted: %v, got: %v", want, got)
		}
	})
}
