package main

import "testing"

func TestAreRunesUnique1(t *testing.T) {
	t.Run("returns false when runes are not unique", func(t *testing.T) {
		got := AreRunesUnique1("abcc")
		expected := false

		if got != expected {
			t.Errorf("expected: %v, got: %v", expected, got)
		}
	})

	t.Run("returns true when runes are unique", func(t *testing.T) {
		got := AreRunesUnique1("abcd")
		expected := true

		if got != expected {
			t.Errorf("expected: %v, got: %v", expected, got)
		}
	})
}

func TestAreRunesUnique2(t *testing.T) {
	t.Run("returns false when runes are not unique", func(t *testing.T) {
		got := AreRunesUnique2("abcc")
		expected := false

		if got != expected {
			t.Errorf("expected: %v, got: %v", expected, got)
		}
	})

	t.Run("returns true when runes are unique", func(t *testing.T) {
		got := AreRunesUnique2("abcd")
		expected := true

		if got != expected {
			t.Errorf("expected: %v, got: %v", expected, got)
		}
	})
}

func TestAreRunesUnique3(t *testing.T) {
	t.Run("returns false when runes are not unique", func(t *testing.T) {
		got := AreRunesUnique3("abcc")
		expected := false

		if got != expected {
			t.Errorf("expected: %v, got: %v", expected, got)
		}
	})

	t.Run("returns true when runes are unique", func(t *testing.T) {
		got := AreRunesUnique3("abcd")
		expected := true

		if got != expected {
			t.Errorf("expected: %v, got: %v", expected, got)
		}
	})
}
