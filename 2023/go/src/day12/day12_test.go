package main

import "testing"

func TestCountContiguousUnknowns(t *testing.T) {
	t.Run("one group", func(t *testing.T) {
		count := 6
		groups := []int{2}
		got := CountContiguousUnknowns(count, groups)
		expect := 5
		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("two groups", func(t *testing.T) {
		count := 6
		groups := []int{2, 1}
		got := CountContiguousUnknowns(count, groups)
		expect := 6
		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})
	t.Run("three groups", func(t *testing.T) {
		count := 10
		groups := []int{3, 2, 1}
		got := CountContiguousUnknowns(count, groups)
		expect := 10

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("??????????????? 1,3,1", func(t *testing.T) {
		count := 15
		groups := []int{1, 3, 1}
		got := CountContiguousUnknowns(count, groups)
		expect := 165

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

}

func TestHowManyWays(t *testing.T) {
	t.Run("groups=required", func(t *testing.T) {
		lut := ".??..??...?##. 1,1,3"
		groups, counts := parseLine(lut)
		got := HowManyWays(groups, counts)
		expect := 4

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("groups=required take 2", func(t *testing.T) {
		lut := "????.#...#... 4,1,1"
		groups, counts := parseLine(lut)
		got := HowManyWays(groups, counts)
		expect := 1

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("groups=required take 3", func(t *testing.T) {
		lut := "????.######..#####. 1,6,5"
		groups, counts := parseLine(lut)
		got := HowManyWays(groups, counts)
		expect := 4

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("just unknowns", func(t *testing.T) {
		lut := "???????????? 3,2,1"
		groups, counts := parseLine(lut)
		got := HowManyWays(groups, counts)
		expect := 35

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("split group into multiple", func(t *testing.T) {
		lut := "?###???????? 3,2,1"
		groups, counts := parseLine(lut)
		got := HowManyWays(groups, counts)
		expect := 10

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("split group into multiple wildcard", func(t *testing.T) {
		lut := "??##???????? 3,2,1"
		groups, counts := parseLine(lut)
		got := HowManyWays(groups, counts)
		expect := 16

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("split group into multiple wildcard", func(t *testing.T) {
		lut := "?#?#?#?#?#?#?#? 1,3,1,6"
		groups, counts := parseLine(lut)
		got := HowManyWays(groups, counts)
		expect := 1

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("split last group", func(t *testing.T) {
		lut := "???.### 1,1,3"
		groups, counts := parseLine(lut)
		got := HowManyWays(groups, counts)
		expect := 1

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("first failure", func(t *testing.T) {
		lut := "?#????#?#.. 2,1,3"
		groups, counts := parseLine(lut)
		got := HowManyWays(groups, counts)
		expect := 3

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("second failure", func(t *testing.T) {
		lut := "??##???????### 1,6,4"
		groups, counts := parseLine(lut)
		got := HowManyWays(groups, counts)
		expect := 1

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("group that must be all broke", func(t *testing.T) {
		lut := "???#??.?.#?#??### 4,8"
		groups, counts := parseLine(lut)
		got := HowManyWays(groups, counts)
		expect := 3

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("one i got wrong I think", func(t *testing.T) {
		lut := "#??..??#.?#? 1,1,1,3"
		groups, counts := parseLine(lut)
		got := HowManyWays(groups, counts)
		expect := 2

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("group that must be all broke", func(t *testing.T) {
		lut := "???#??.?.#?#??### 4,8"
		groups, counts := parseLine(lut)
		got := HowManyWays(groups, counts)
		expect := 3

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("another test", func(t *testing.T) {
		lut := ".#?????..???????.? 6,7"
		groups, counts := parseLine(lut)
		got := HowManyWays(groups, counts)
		expect := 1

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("another test", func(t *testing.T) {
		lut := "???#???.?#?????? 1,4,2,2,1"
		groups, counts := parseLine(lut)
		got := HowManyWays(groups, counts)
		expect := 12

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("#??.???..? 1,1,1", func(t *testing.T) {
		lut := "#??.???..? 1,1,1"
		groups, counts := parseLine(lut)
		got := HowManyWays(groups, counts)
		expect := 8

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("???#?# 3", func(t *testing.T) {
		lut := "???#?# 3"
		groups, counts := parseLine(lut)
		got := HowManyWays(groups, counts)
		expect := 1

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("??#?#????#?# 6,3", func(t *testing.T) {
		lut := "??#?#????#?# 6,3"
		groups, counts := parseLine(lut)
		got := HowManyWays(groups, counts)
		expect := 3

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("#??.#??.??#?#????#?# 2,1,6,3", func(t *testing.T) {
		lut := "#??.#??.??#?#????#?# 2,1,6,3"
		groups, counts := parseLine(lut)
		got := HowManyWays(groups, counts)
		expect := 3

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("????#??? 3,3", func(t *testing.T) {
		lut := "????#??? 3,3"
		groups, counts := parseLine(lut)
		got := HowManyWays(groups, counts)
		expect := 1

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("????..??#?#?#?? 3,8", func(t *testing.T) {
		lut := "????..??#?#?#??.???? 3,8"
		groups, counts := parseLine(lut)
		got := HowManyWays(groups, counts)
		expect := 4

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})
}

func TestBruteForce(t *testing.T) {
	t.Run("groups=required", func(t *testing.T) {
		lut := ".??..??...?##. 1,1,3"
		groups, counts := parseLine(lut)
		got := bruteForce(groups, counts)
		expect := 4

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("groups=required take 2", func(t *testing.T) {
		lut := "????.#...#... 4,1,1"
		groups, counts := parseLine(lut)
		got := bruteForce(groups, counts)
		expect := 1

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("groups=required take 3", func(t *testing.T) {
		lut := "????.######..#####. 1,6,5"
		groups, counts := parseLine(lut)
		got := bruteForce(groups, counts)
		expect := 4

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("just unknowns", func(t *testing.T) {
		lut := "???????????? 3,2,1"
		groups, counts := parseLine(lut)
		got := bruteForce(groups, counts)
		expect := 35

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("split group into multiple", func(t *testing.T) {
		lut := "?###???????? 3,2,1"
		groups, counts := parseLine(lut)
		got := bruteForce(groups, counts)
		expect := 10

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("split group into multiple wildcard", func(t *testing.T) {
		lut := "??##???????? 3,2,1"
		groups, counts := parseLine(lut)
		got := bruteForce(groups, counts)
		expect := 16

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("split group into multiple wildcard", func(t *testing.T) {
		lut := "?#?#?#?#?#?#?#? 1,3,1,6"
		groups, counts := parseLine(lut)
		got := bruteForce(groups, counts)
		expect := 1

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("split last group", func(t *testing.T) {
		lut := "???.### 1,1,3"
		groups, counts := parseLine(lut)
		got := bruteForce(groups, counts)
		expect := 1

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("first failure", func(t *testing.T) {
		lut := "?#????#?#.. 2,1,3"
		groups, counts := parseLine(lut)
		got := bruteForce(groups, counts)
		expect := 3

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("second failure", func(t *testing.T) {
		lut := "??##???????### 1,6,4"
		groups, counts := parseLine(lut)
		got := bruteForce(groups, counts)
		expect := 1

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("group that must be all broke", func(t *testing.T) {
		lut := "???#??.?.#?#??### 4,8"
		groups, counts := parseLine(lut)
		got := bruteForce(groups, counts)
		expect := 3

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("another test", func(t *testing.T) {
		lut := ".#?????..???????.? 6,7"
		groups, counts := parseLine(lut)
		got := bruteForce(groups, counts)
		expect := 1

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("another test", func(t *testing.T) {
		lut := "???#???.?#?????? 1,4,2,2,1"
		groups, counts := parseLine(lut)
		got := bruteForce(groups, counts)
		expect := 12

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("#??.???..? 1,1,1", func(t *testing.T) {
		lut := "#??.???..? 1,1,1"
		groups, counts := parseLine(lut)
		got := bruteForce(groups, counts)
		expect := 8

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("???#?# 3", func(t *testing.T) {
		lut := "???#?# 3"
		groups, counts := parseLine(lut)
		got := bruteForce(groups, counts)
		expect := 1

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("??#?#????#?# 6,3", func(t *testing.T) {
		lut := "??#?#????#?# 6,3"
		groups, counts := parseLine(lut)
		got := bruteForce(groups, counts)
		expect := 3

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("#??.#??.??#?#????#?# 2,1,6,3", func(t *testing.T) {
		lut := "#??.#??.??#?#????#?# 2,1,6,3"
		groups, counts := parseLine(lut)
		got := bruteForce(groups, counts)
		expect := 3

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("????#??? 3,3", func(t *testing.T) {
		lut := "????#??? 3,3"
		groups, counts := parseLine(lut)
		got := bruteForce(groups, counts)
		expect := 1

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})

	t.Run("????..??#?#?#?? 3,8", func(t *testing.T) {
		lut := "????..??#?#?#??.???? 3,8"
		groups, counts := parseLine(lut)
		got := bruteForce(groups, counts)
		expect := 4

		if got != expect {
			t.Errorf("Expected %d, but got %d", expect, got)
		}
	})
}
