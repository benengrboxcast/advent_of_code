package numerics

import "testing"

func TestGetNumeric(t *testing.T) {
	t.Run("number at start", func(t *testing.T) {
		expected := 123
		expected_idx := 3
		got, index := GetNumeric("123 asdf", 0)

		if got != expected {
			t.Errorf("Expected %d, but got %d", expected, got)
		}

		if index != expected_idx {
			t.Errorf("Index should have been %d, but it was %d", expected_idx, index)
		}

	})

	t.Run("number in the middel", func(t *testing.T) {
		expected := 987
		expected_idx := 5
		got, index := GetNumeric("ab987 123 asdf", 0)

		if got != expected {
			t.Errorf("Expected %d, but got %d", expected, got)
		}

		if index != expected_idx {
			t.Errorf("Index should have been %d, but it was %d", expected_idx, index)
		}
	})

	t.Run("start in the middel", func(t *testing.T) {
		expected := 654
		expected_idx := 9
		got, index := GetNumeric("ab987 654 asdf", 5)

		if got != expected {
			t.Errorf("Expected %d, but got %d", expected, got)
		}

		if index != expected_idx {
			t.Errorf("Index should have been %d, but it was %d", expected_idx, index)
		}
	})

	t.Run("number at end 1023", func(t *testing.T) {
		expected := 1023
		test_str := "ab987 123 asdf 1023"
		expected_idx := len(test_str)
		got, index := GetNumeric("ab987 123 asdf 1023", 9)

		if got != expected {
			t.Errorf("Expected %d, but got %d", expected, got)
		}

		if index != expected_idx {
			t.Errorf("Index should have been %d, but it was %d", expected_idx, index)
		}
	})
}
