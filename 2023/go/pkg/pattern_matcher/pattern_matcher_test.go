package pattern_matcher

import "testing"

func TestSingleValue(t *testing.T) {
	t.Run("single value initialsed", func(t *testing.T) {
		root := PatternNode{Label: ' ', Value: -1}
		testcase := [1]Pattern{{Label: "one", Value: 1}}
		CreateGraph(&root, testcase[:])

		if len(root.Children) != 1 {
			t.Errorf("Got %d childen and expected 1", len(root.Children))
		}
	})

	t.Run("correct value at end of graph", func(t *testing.T) {
		root := PatternNode{Label: ' ', Value: -1}
		testcase := [1]Pattern{{Label: "one", Value: 1}}
		CreateGraph(&root, testcase[:])

		count := 0
		value := -1
		current := &root

		for value < 0 {
			value = current.Value
			if value < 0 {
				if len(current.Children) != 1 {
					t.Errorf("Got %d when 1 was expected on node %d with label %b", len(current.Children), count, current.Label)
				}
				count++
				current = current.Children[0]

			} else {
				if current.Value != 1 {
					t.Errorf("Got %d when expected value was 1", current.Value)
				}

				if count != 3 {
					t.Errorf("It took %d steps to get to the end of the graph, when 3 were expected", count)
				}
			}
		}
	})

	t.Run("value is assigned when the length is 1", func(t *testing.T) {
		root := PatternNode{Label: ' ', Value: -1}
		testcase := [1]Pattern{{Label: "e", Value: 12}}
		CreateGraph(&root, testcase[:])
		if len(root.Children) != 1 {
			t.Errorf("expected root to have 1 child, but there were %d", len(root.Children))
		}

		if root.Children[0].Value != 12 {
			t.Errorf("Expected to get a value of 12, but it was %d", root.Children[0].Value)
		}
	})
}

func TestDualValue(t *testing.T) {
	t.Run("dual path different starts", func(t *testing.T) {
		root := PatternNode{Label: ' ', Value: -1}
		testcase := [2]Pattern{{Label: "two", Value: 2}, {Label: "one", Value: 3}}
		CreateGraph(&root, testcase[:])

		if len(root.Children) != 2 {
			t.Errorf("Expected the root node to have 2 children, but there were %d children", len(root.Children))
		}

		if root.Children[0].Label != 't' {
			t.Errorf("Expected the first child to have a label of 't', gut it was '%c'", root.Children[0].Label)
		}

		if root.Children[1].Label != 'o' {
			t.Errorf("Expected the second child to have a label of 'o', but it was '%c'", root.Children[1].Label)
		}
	})

	t.Run("dual path same start", func(t *testing.T) {
		root := PatternNode{Label: ' ', Value: -1}
		testcase := [2]Pattern{{Label: "two", Value: 2}, {Label: "three", Value: 3}}
		CreateGraph(&root, testcase[:])

		if len(root.Children) != 1 {
			t.Errorf("Expected the root node to have 2 children, but there were %d children", len(root.Children))
		}

		if root.Children[0].Label != 't' {
			t.Errorf("Expected the first child to have a label of 't', gut it was '%c'", root.Children[0].Label)
		}
	})

	t.Run("full", func(t *testing.T) {
		root := PatternNode{Label: ' ', Value: -1}
		testcase := [9]Pattern{
			{Label: "one", Value: 1},
			{Label: "two", Value: 2},
			{Label: "three", Value: 3},
			{Label: "four", Value: 4},
			{Label: "five", Value: 5},
			{Label: "six", Value: 6},
			{Label: "seven", Value: 7},
			{Label: "eight", Value: 8},
			{Label: "nine", Value: 9},
		}
		CreateGraph(&root, testcase[:])

		if len(root.Children) != 6 {
			t.Errorf("Expected there to be 6 children, but there were %d", len(root.Children))
		}
	})
}

func TestNextChar(t *testing.T) {
	t.Run("invalid character", func(t *testing.T) {
		root := PatternNode{Label: ' ', Value: -1}
		testcase := [9]Pattern{
			{Label: "one", Value: 1},
			{Label: "two", Value: 2},
			{Label: "three", Value: 3},
			{Label: "four", Value: 4},
			{Label: "five", Value: 5},
			{Label: "six", Value: 6},
			{Label: "seven", Value: 7},
			{Label: "eight", Value: 8},
			{Label: "nine", Value: 9},
		}

		CreateGraph(&root, testcase[:])

		current := root.Next('b', &root)
		if current != &root {
			t.Errorf("Expected to get root Node, but got %c", current.Label)
		}

	})

	t.Run("valid character", func(t *testing.T) {
		root := PatternNode{Label: ' ', Value: -1}
		testcase := [9]Pattern{
			{Label: "one", Value: 1},
			{Label: "two", Value: 2},
			{Label: "three", Value: 3},
			{Label: "four", Value: 4},
			{Label: "five", Value: 5},
			{Label: "six", Value: 6},
			{Label: "seven", Value: 7},
			{Label: "eight", Value: 8},
			{Label: "nine", Value: 9},
		}

		CreateGraph(&root, testcase[:])

		current := root.Next('o', &root)
		if current == &root {
			t.Errorf("Expected valid Node, but got nil")
			return
		}

		if len(current.Children) != 1 {
			t.Errorf("Current node as %d children when 1 was expected", len(current.Children))
		}
	})

	t.Run("five", func(t *testing.T) {
		root := PatternNode{Label: ' ', Value: -1}
		testcase := [9]Pattern{
			{Label: "one", Value: 1},
			{Label: "two", Value: 2},
			{Label: "three", Value: 3},
			{Label: "four", Value: 4},
			{Label: "five", Value: 5},
			{Label: "six", Value: 6},
			{Label: "seven", Value: 7},
			{Label: "eight", Value: 8},
			{Label: "nine", Value: 9},
		}

		CreateGraph(&root, testcase[:])

		test_string := "five"
		current := &root
		steps := 0

		for i := 0; i < len(test_string); i++ {
			current = current.Next(test_string[i], &root)
			steps++
			if current.Value >= 0 {
				if steps != 4 {
					t.Errorf("Expected 5 steps but got %d", steps)
				}

				if current.Value != 5 {
					t.Errorf("Expected value to be 5, but got %d", current.Value)
				}
			} else {
				if current == &root {
					t.Errorf("Returned to root on step %d", steps)
				}
			}
		}
	})

	t.Run("return_to_root", func(t *testing.T) {
		root := PatternNode{Label: ' ', Value: -1}
		testcase := [9]Pattern{
			{Label: "one", Value: 1},
			{Label: "two", Value: 2},
			{Label: "three", Value: 3},
			{Label: "four", Value: 4},
			{Label: "five", Value: 5},
			{Label: "six", Value: 6},
			{Label: "seven", Value: 7},
			{Label: "eight", Value: 8},
			{Label: "nine", Value: 9},
		}

		CreateGraph(&root, testcase[:])

		test_string := "sevet"
		current := &root
		steps := 0

		for i := 0; i < len(test_string); i++ {
			current = current.Next(test_string[i], &root)
			steps++

			if current.Value >= 0 {
				t.Errorf("Get value of %d, when -1 was expected", current.Value)
				return
			} else if current == &root {
				if steps == 5 {
					return
				} else {
					t.Errorf("Returned to root on step %d", steps)
					return
				}
			}
		}
	})

	t.Run("start new with valid letter", func(t *testing.T) {
		root := PatternNode{Label: ' ', Value: -1}
		testcase := [9]Pattern{
			{Label: "one", Value: 1},
			{Label: "two", Value: 2},
			{Label: "three", Value: 3},
			{Label: "four", Value: 4},
			{Label: "five", Value: 5},
			{Label: "six", Value: 6},
			{Label: "seven", Value: 7},
			{Label: "eight", Value: 8},
			{Label: "nine", Value: 9},
		}

		CreateGraph(&root, testcase[:])
		test_string := "ontwo"
		current := &root
		steps := 0

		for i := 0; i < len(test_string); i++ {
			current = current.Next(test_string[i], &root)
			steps++

			if current.Value >= 0 {
				t.Errorf("Get value of %d, when -1 was expected", current.Value)
				return
			} else if current == &root {
				if steps == 5 && current.Value == 2 {
					return
				}
			}
		}

		t.Errorf("Got to end of loop")
	})
}
