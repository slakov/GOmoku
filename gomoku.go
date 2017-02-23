package main

import (
	"fmt"
	"strconv"
	"strings"
)

var boardSize int = 9
var boardArea int = boardSize * boardSize
var board string

func initBoard() {
	board = strings.Repeat(".", boardArea)
}

func printBoard(b string) {

	columnLabels := make([]int, boardSize)

	for i := range columnLabels {
		columnLabels[i] = 1 + i
	}

	valuesText := []string{}
	valuesText = append(valuesText, " ")

	for i := range columnLabels {
		number := columnLabels[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}

	fmt.Println(strings.Join(valuesText, ""))

	for i := 0; i < boardSize; i++ {
		firstIndex := i * boardSize
		lastIndex := i*boardSize + boardSize
		fmt.Print(i + 1)
		fmt.Println(b[firstIndex:lastIndex])
	}

}

func xyToIndex(x, y int) int {
	y = y - 1
	x = x - 1
	index := y*boardSize + x
	return index
}

func main() {

	fmt.Println("gomoku here")
	initBoard()
	printBoard(board)

	var x, y int
	var x_turn bool
	x_turn = true
	var gEnd bool = false

	for gEnd == false {

		// Who's turn
		if x_turn {
			fmt.Println("X is aan de beurt..")
		} else {
			fmt.Println("O is aan de beurt..")
		}

		// Input move
		fmt.Print("Enter column:")
		fmt.Scan(&x)
		fmt.Print("Enter row:")
		fmt.Scan(&y)
		index := xyToIndex(x, y)

		// 46 is ASCII for dot
		if board[index] != 46 {
			fmt.Println("illegal move")
			continue
		}

		if x_turn {
			board = board[:index] + "X" + board[index+1:]
		} else {
			board = board[:index] + "O" + board[index+1:]
		}

		x_turn = !x_turn
		printBoard(board)

		// check board status
		if gStatus() == 1 {
		}
		if gStatus() == 2 {
		}

	}
}
