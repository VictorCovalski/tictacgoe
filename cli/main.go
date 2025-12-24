package main

import (
	"fmt"

	"github.com/VictorCovalski/tictacgoe/game"
)

func PrintBoard(board [][]game.Mark) {
	fmt.Printf(" -----------\n")
	for i := range board {
		fmt.Printf("| %s | %s | %s |\n", board[i][0], board[i][1], board[i][2])
	}
	fmt.Printf(" -----------\n")
}

func PlayGame() {
	game := game.NewGame()
	fmt.Println("TicTacToe Game")
	var x, y int
	for {
		fmt.Println("Game Board")
		PrintBoard(game.GetBoard())
		fmt.Printf("[%s] Player turn: ", game.GetCurrentPlayerName())
		fmt.Scanf("%d %d", &x, &y)

		if err := game.PlayTurn(x, y); err != nil {
			fmt.Printf("\nPlay again! %s\n", err)
			continue
		}

		if winner, ok := game.GetWinner(); ok {
			fmt.Printf("Player %s is the winner!\n", winner)
			break
		}
	}
}

func main() {
	PlayGame()
}
