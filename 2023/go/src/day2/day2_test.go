package main

import "testing"

func TestLineToGame(t *testing.T) {
	t.Run("first example", func(t *testing.T) {
		test_str := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
		game := stringToGame(test_str)

		if game.Id != 1 {
			t.Errorf("Expected id to be 1, but it was %d", game.Id)
		}
		if len(game.Draws) != 3 {
			t.Errorf("Expected 3 draws but got %d", len(game.Draws))
		}

		if game.MaxBlue != 6 {
			t.Errorf("MaxBlue was %d but expected 6", game.MaxBlue)
		}

		if game.MaxGreen != 2 {
			t.Errorf("MaxGreen was %d but expected 2", game.MaxGreen)
		}

		if game.MaxRed != 4 {
			t.Errorf("MaxRed was %d but expected 4", game.MaxRed)
		}

		if game.Draws[0].Blue != 3 {
			t.Errorf("Expected draw 0 blue to be 3, but it was %d", game.Draws[0].Blue)
		}
		if game.Draws[0].Red != 4 {
			t.Errorf("Expected draw 0 red to be 4, but it was %d", game.Draws[0].Blue)
		}
		if game.Draws[0].Green != 0 {
			t.Errorf("Expected draw 0 green to be 0, but it was %d", game.Draws[0].Blue)
		}

		if game.Draws[1].Blue != 6 {
			t.Errorf("Expected draw 1 blue to be 6, but it was %d", game.Draws[0].Blue)
		}
		if game.Draws[1].Red != 1 {
			t.Errorf("Expected draw 1 red to be 1, but it was %d", game.Draws[0].Blue)
		}
		if game.Draws[1].Green != 2 {
			t.Errorf("Expected draw 1 green to be 2, but it was %d", game.Draws[0].Blue)
		}

		if game.Draws[2].Blue != 0 {
			t.Errorf("Expected draw 2 blue to be 0, but it was %d", game.Draws[0].Blue)
		}
		if game.Draws[2].Red != 0 {
			t.Errorf("Expected draw 2 red to be 0, but it was %d", game.Draws[0].Blue)
		}
		if game.Draws[2].Green != 2 {
			t.Errorf("Expected draw 0 green to be 2, but it was %d", game.Draws[0].Blue)
		}
	})
}
