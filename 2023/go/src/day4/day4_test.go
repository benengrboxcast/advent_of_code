package main

import "testing"

func TestParseWinningNumbers(t *testing.T) {
	t.Run("all two digit", func(t *testing.T) {
		uut := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"

		values, index := parseWinningNumbers(uut)

		if index != 25 {
			t.Errorf("Expected to be on index 25, but got %d", index)
		}

		if len(values) != 5 {
			t.Errorf("Expected to get 5 values, but got %d", len(values))
		}

		expected := [5]int{17, 41, 48, 83, 86}
		for i, val := range expected {
			if val != values[i] {
				t.Errorf("Expected index %d to be %d, but it was %d", i, expected[i], values[i])
			}
		}
	})

	t.Run("single digit", func(t *testing.T) {
		uut := "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1"

		values, index := parseWinningNumbers(uut)

		if index != 25 {
			t.Errorf("Expected to be on index 25, but got %d", index)
		}

		if len(values) != 5 {
			t.Errorf("Expected to get 5 values, but got %d", len(values))
		}

		expected := [5]int{1, 21, 44, 53, 59}
		for i, val := range expected {
			if val != values[i] {
				t.Errorf("Expected index %d to be %d, but it was %d", i, expected[i], values[i])
			}
		}
	})
}
