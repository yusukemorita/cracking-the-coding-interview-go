package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestBuildOrder2(t *testing.T) {
	t.Run("returns a valid build order", func(t *testing.T) {
		projects := []string{"a", "b", "c"}
		dependencies := []dependency{
			{
				dependant: "a",
				depended:  "b",
			},
			{
				dependant: "b",
				depended:  "c",
			},
		}

		got, ok := BuildOrder2(projects, dependencies)
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
				dependant: "a",
				depended:  "b",
			},
			{
				dependant: "b",
				depended:  "c",
			},
			{
				dependant: "c",
				depended:  "a",
			},
		}

		_, ok := BuildOrder2(projects, dependencies)
		if ok {
			t.Error()
		}
	})
}
