package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrossWinsHorizontalGame(t *testing.T) {
	game := NewGame()

	moves := [][2]int{{0, 0}, {1, 0}, {0, 1}, {2, 0}, {0, 2}}

	// final board (X = Cross, O = Circle, E = Empty)
	// X X X
	// O E E
	// O E E

	for _, m := range moves {
		assert.NoError(t, game.PlayTurn(m[0], m[1]))
	}

	if winner, ok := game.GetWinner(); ok {
		assert.Equal(t, winner, "Cross")
	} else {
		t.Error("Cross didnt win the game.")
	}
}

func TestCrossWinsVerticalGame(t *testing.T) {
	game := NewGame()

	// final board (X = Cross, O = Circle, E = Empty)
	// X O E
	// X O E
	// X E E
	moves := [][2]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}, {2, 0}}

	for _, m := range moves {
		assert.NoError(t, game.PlayTurn(m[0], m[1]))
	}

	if winner, ok := game.GetWinner(); ok {
		assert.Equal(t, winner, "Cross")
	} else {
		t.Error("Cross didnt win the game.")
	}
}

func TestCrossWinsLeftDiagonalGame(t *testing.T) {
	game := NewGame()

	// final board (X = Cross, O = Circle, E = Empty)
	// X O E
	// O X E
	// E E X
	moves := [][2]int{{0, 0}, {0, 1}, {1, 1}, {1, 0}, {2, 2}}

	for _, m := range moves {
		assert.NoError(t, game.PlayTurn(m[0], m[1]))
	}

	if winner, ok := game.GetWinner(); ok {
		assert.Equal(t, winner, "Cross")
	} else {
		t.Error("Cross didnt win the game.")
	}
}

func TestCrossWinsRightDiagonalGame(t *testing.T) {
	game := NewGame()

	// final board (X = Cross, O = Circle, E = Empty)
	// E O X
	// O X E
	// X E E
	moves := [][2]int{{0, 2}, {0, 1}, {1, 1}, {1, 0}, {2, 0}}

	for _, m := range moves {
		assert.NoError(t, game.PlayTurn(m[0], m[1]))
	}

	if winner, ok := game.GetWinner(); ok {
		assert.Equal(t, winner, "Cross")
	} else {
		t.Error("Cross didnt win the game.")
	}
}

func TestInvalidMoves(t *testing.T) {
	game := NewGame()

	moves := [][2]int{{-1, 2}, {3, 1}, {1, 4}, {1, -2}, {-2, 0}}

	for _, m := range moves {
		assert.ErrorContains(t, game.PlayTurn(m[0], m[1]), "is an invalid position")
	}
}

func TestOccupiedTile(t *testing.T) {
	game := NewGame()

	assert.NoError(t, game.PlayTurn(1, 1))
	assert.ErrorContains(t, game.PlayTurn(1, 1), "not empty")
}
