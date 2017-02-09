package main

import (
	"fmt"
	"strconv"
	"strings"
)

var boardSize int = 8
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
	fmt.Print("Enter column:")
	fmt.Scan(&x)
	fmt.Print("Enter row:")
	fmt.Scan(&y)
	index := xyToIndex(x, y)
	board = board[:index] + "X" + board[index+1:]

	printBoard(board)

}
