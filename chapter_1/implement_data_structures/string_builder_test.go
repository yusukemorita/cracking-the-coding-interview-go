package main

import "testing"

func TestStringBuilder_AddString(t *testing.T) {
	builder := NewStringBuilder()
	builder.addString("hello")
	builder.addString("hi")
	builder.addString("goodmorning")

	expected := "hellohigoodmorning"
	got := builder.toString()
	if builder.toString() != expected {
		t.Errorf("expected: %s, got: %v", expected, got)
	}
}
