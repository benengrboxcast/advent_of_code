package main

import "testing"

func TestParseLine(t *testing.T) {
	t.Run("no symbols", func(t *testing.T) {
		test_line := "467..114.."
		s, n := parse_line(test_line)

		for idx, val := range s {
			if val {
				t.Errorf("Unexpected symbol was found at %d", idx)
			}
		}

		if len(n) != 2 {
			t.Errorf("There were %d numbers when 2 were expected", len(n))
			return
		}

		uut := n[0]
		if uut.Value != 467 && uut.StartIndex != 0 && uut.EndIndex != 2 {
			t.Errorf("First Number was unexpected Got %d at %d-%d", uut.Value, uut.StartIndex, uut.EndIndex)

		}

		uut = n[1]
		if uut.Value != 114 && uut.StartIndex != 5 && uut.EndIndex != 7 {
			t.Errorf("First Number was unexpected Got %d at %d-%d", uut.Value, uut.StartIndex, uut.EndIndex)
		}
	})

	t.Run("only symbols", func(t *testing.T) {
		test_line := "...*......"
		s, n := parse_line(test_line)

		if len(s) != 10 {
			t.Errorf("There were %d symbols when 10 was expected", len(s))
			return
		}

		for idx, val := range s {
			if val != (idx == 3) {
				t.Errorf("Invalid symbol at index %d, Got %t Expected %t", idx, val, idx == 3)
			}
		}
		if len(n) != 0 {
			t.Errorf("There were %d numbers when 0 were expected", len(n))
			return
		}
	})

	t.Run("adjacent symbol and number", func(t *testing.T) {
		test_line := "617*......"
		s, n := parse_line(test_line)

		if len(s) != 10 {
			t.Errorf("There were %d symbols when 10 was expected", len(s))
			return
		}

		for idx, val := range s {
			if val != (idx == 3) {
				t.Errorf("Invalid symbol at index %d, Got %t Expected %t", idx, val, idx == 3)
			}
		}

		if len(n) != 1 {
			t.Errorf("There were %d numbers when 1 were expected", len(n))
			return
		}

		uut := n[0]
		if uut.Value != 617 && uut.StartIndex != 0 && uut.EndIndex != 2 {
			t.Errorf("Invalid number expected 617 0-2 and got %d %d-%d", uut.Value, uut.StartIndex, uut.EndIndex)
		}
	})
}
