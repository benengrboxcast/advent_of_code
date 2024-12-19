package main

import "testing"

func TestCreateMapRange(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		dut := "50 98 2"
		got := stringToMapRange(dut)

		if got.Dest != 50 || got.Source != 98 || got.Size != 2 {
			t.Errorf("unexpected result got (source %d, dest %d, size %d", got.Source, got.Dest, got.Size)
		}
	})
}

func TestCombineMaps(t *testing.T) {

	t.Run("separate sources", func(t *testing.T) {
		first := []MapRange{{Source: 15, Dest: 0, Size: 37}}
		second := []MapRange{{Source: 52, Dest: 37, Size: 2}}

		got := combineMaps(first, second)
		expected := []MapRange{
			{Source: 15, Dest: 0, Size: 37},
			{Source: 52, Dest: 37, Size: 2},
		}
		if len(got) != len(expected) {
			t.Errorf("Size is wrong, it should be %d, but it is %d", len(expected), len(got))
		}

		for i := 0; i < len(got); i++ {
			if expected[i] != got[i] {
				t.Errorf("Wrong map")
			}
		}
	})

	t.Run("D1 overlaps S1", func(t *testing.T) {
		first := []MapRange{{Source: 15, Dest: 0, Size: 2}}
		second := []MapRange{{Source: 0, Dest: 37, Size: 2}}

		got := combineMaps(first, second)
		expected := []MapRange{
			{Source: 15, Dest: 37, Size: 2},
		}
		if len(got) != len(expected) {
			t.Errorf("Size is wrong, it should be %d, but it is %d", len(expected), len(got))
			return
		}

		for i := 0; i < len(got); i++ {
			if expected[i] != got[i] {
				t.Errorf("Wrong map")
			}
		}
	})

	t.Run("D1 partial overlap S1", func(t *testing.T) {
		first := []MapRange{{Source: 15, Dest: 0, Size: 2}}
		second := []MapRange{{Source: 1, Dest: 37, Size: 5}}

		got := combineMaps(first, second)
		expected := []MapRange{
			{Source: 15, Dest: 0, Size: 1},
			{Source: 16, Dest: 37, Size: 1},
			{Source: 2, Dest: 38, Size: 3},
		}
		if len(got) != len(expected) {
			t.Errorf("Size is wrong, it should be %d, but it is %d", len(expected), len(got))
			return
		}

		for i := 0; i < len(got); i++ {
			if expected[i] != got[i] {
				t.Errorf("Wrong map")
			}
		}
	})
	//t.Run("second fully inside", func(t *testing.T) {
	//	first := []MapRange{{Source: 50, Dest: 52, Size: 48}}
	//	second := []MapRange{{Source: 52, Dest: 37, Size: 2}}
	//
	//	got := combineMaps(first, second)
	//	expected := []MapRange{
	//		{Source: 50, Dest: 52, Size: 2},
	//		{Source: 52, Dest: 35, Size: 2},
	//		{Source: 54, Dest: 56, Size: 43},
	//	}
	//	if len(got) != len(expected) {
	//		t.Errorf("Size is wrong, it should be %d, but it is %d", len(expected), len(got))
	//	}
	//
	//	for i := 0; i < len(got); i++ {
	//		if expected[i] != got[i] {
	//			t.Errorf("Wrong map")
	//		}
	//	}
	//})
}
