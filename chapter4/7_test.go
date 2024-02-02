package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestBuildOrder(t *testing.T) {
	t.Run("returns a valid build order", func(t *testing.T) {
		projects := []string{"a", "b", "c"}
		dependencies := []dependency{
			{
				from: "a",
				to:   "b",
			},
			{
				from: "b",
				to:   "c",
			},
		}

		got, ok := BuildOrder(projects, dependencies)
		if !ok {
			t.Error()
		}

		if strings.Join(got, ",") != "c,b,a" {
			fmt.Println(got)
			t.Error()
		}
	})

	t.Run("returns false when no valid build order", func(t *testing.T) {
		projects := []string{"a", "b", "c"}
		dependencies := []dependency{
			{
				from: "a",
				to:   "b",
			},
			{
				from: "b",
				to:   "c",
			},
			{
				from: "c",
				to:   "a",
			},
		}

		_, ok := BuildOrder(projects, dependencies)
		if ok {
			t.Error()
		}
	})
}
