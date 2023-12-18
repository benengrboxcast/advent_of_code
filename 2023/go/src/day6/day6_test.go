package main

import "testing"

func TestWaysToWin(t *testing.T) {
	t.Run("time odd", func(t *testing.T) {
		got := waysToWin(7, 9)
		expected := 4
		if got != expected {
			t.Errorf("Test failed, expected %d and got %d", expected, got)
		}
	})

	t.Run("time even", func(t *testing.T) {
		got := waysToWin(30, 200)
		expected := 9
		if got != expected {
			t.Errorf("Test failed, expected %d and got %d", expected, got)
		}
	})
}
