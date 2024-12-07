package main

import "testing"

func TestCreateMapRange(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		expect := 10
		got := 1

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})
}
