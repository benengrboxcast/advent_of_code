package main

import "testing"

func TestFindNextValue(t *testing.T) {
	t.Run("example1", func(t *testing.T) {
		uut := []float64{0, 3, 6, 9, 12, 15}
		got := findNextValue(uut)
		expected := 18.0

		if got != expected {
			t.Error("Error expected", expected, " got ", got, "with input ", uut)
		}
	})

	t.Run("example2", func(t *testing.T) {
		uut := []float64{1, 3, 6, 10, 15, 21}
		got := findNextValue(uut)
		expected := 28.0

		if got != expected {
			t.Error("Error expected", expected, " got ", got, "with input ", uut)
		}
	})

	t.Run("example3", func(t *testing.T) {
		uut := []float64{10, 13, 16, 21, 30, 45}
		got := findNextValue(uut)
		expected := 68.0

		if got != expected {
			t.Error("Error expected", expected, " got ", got, "with input ", uut)
		}
	})

	t.Run("negatives", func(t *testing.T) {
		uut := []float64{-4, -3, 1, 6, 15, 44, 130, 350, 877, 2122, 5057, 11920, 27747, 63713, 144400, 323457, 716912, 1573135, 716912, 3418095, 7353718, 15663901}
		got := findNextValue(uut)
		expected := 15304826720.0

		if got != expected {
			t.Error("Error expected", expected, " got ", got, "with input ", uut)
		}
	})
}

func TestFindFirstValue(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		uut := []float64{10, 13, 16, 21, 30, 45}
		got := findFirstValue(uut)
		expected := 5.0

		if got != expected {
			t.Error("Error expected", expected, " got ", got, "with input ", uut)
		}
	})
}
