package main

import "testing"

func TestIsRotation(t *testing.T) {
	t.Run("returns true when is rotation", func(t *testing.T) {
		got := IsRotation("erbottlewat", "waterbottle")
		if !got {
			t.Error("expected true")
		}
	})

	t.Run("returns true when same string", func(t *testing.T) {
		got := IsRotation("waterbottle", "waterbottle")
		if !got {
			t.Error("expected true")
		}
	})

	t.Run("returns false when not a rotation", func(t *testing.T) {
		got := IsRotation("waterbottl", "waterbottle")
		if got {
			t.Error("expected false")
		}
	})

	t.Run("returns false when different lengths", func(t *testing.T) {
		got := IsRotation("waterbottl", "t")
		if got {
			t.Error("expected false")
		}
	})
}
