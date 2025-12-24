package game

import (
	"fmt"
)

type Mark int

const boardSize = 3

const (
	Empty  Mark = iota
	Circle Mark = iota
	Cross  Mark = iota
)

type TicTacToe struct {
	board         [][]Mark
	currentPlayer Mark
	finished      bool
}

func NewGame() *TicTacToe {
	return &TicTacToe{
		board:         [][]Mark{{Empty, Empty, Empty}, {Empty, Empty, Empty}, {Empty, Empty, Empty}},
		currentPlayer: Cross,
		finished:      false,
	}
}

func (t *TicTacToe) nextPlayer() {
	if t.currentPlayer == Circle {
		t.currentPlayer = Cross
	} else {
		t.currentPlayer = Circle
	}
}

func (t *TicTacToe) GetCurrentPlayerName() string {
	if t.currentPlayer == Circle {
		return "Circle"
	} else {
		return "Cross"
	}
}

func (t *TicTacToe) PlayTurn(x, y int) error {
	if x >= boardSize || y >= boardSize || x < 0 || y < 0 {
		return fmt.Errorf("<%d,%d> is an invalid position", x, y)
	}

	if t.board[x][y] > 0 {
		return fmt.Errorf("Tile is not empty")
	}

	t.board[x][y] = t.currentPlayer

	t.checkVictory()

	// if the game is not over yet
	if !t.finished {
		t.nextPlayer()
	}

	return nil
}

func (t *TicTacToe) GetWinner() (string, bool) {
	if t.finished {
		return t.GetCurrentPlayerName(), true
	} else {
		return "", false
	}
}

func (t *TicTacToe) checkVictory() {
	b := t.board
	p := t.currentPlayer

	// verify rows
	for i := 0; i < boardSize; i++ {
		if b[i][0] == p && b[i][1] == p && b[i][2] == p {
			t.finished = true
		}
	}

	// verify columns
	for i := 0; i < boardSize; i++ {
		if b[0][i] == p && b[1][i] == p && b[2][i] == p {
			t.finished = true
		}
	}

	// verify diagonal right
	if b[0][0] == p && b[1][1] == p && b[2][2] == p {
		t.finished = true
	}

	// verify diagonal left
	if b[0][2] == p && b[1][1] == p && b[2][0] == p {
		t.finished = true
	}
}

func (t *TicTacToe) Reset() {
	t = NewGame()
}
