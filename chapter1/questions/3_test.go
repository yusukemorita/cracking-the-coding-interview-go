package main

import "testing"

func TestURLify(t *testing.T) {
	t.Run("returns URLified string when string length is just right", func(t *testing.T) {
		got := URLify("hello hi good morning      ", 21)
		want := "hello%20hi%20good%20morning"

		if got != want {
			t.Errorf("wanted: %v, got: %v", want, got)
		}
	})

	t.Run("returns URLified string when string length is longer than necessary", func(t *testing.T) {
		got := URLify("hello hi good morning         ", 21)
		want := "hello%20hi%20good%20morning"

		if got != want {
			t.Errorf("wanted: %v, got: %v", want, got)
		}
	})
}
