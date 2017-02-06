package main

import (
	"fmt"
	"strings"
)

var boardSize int = 8
var boardArea int = boardSize * boardSize
var board string

func initBoard() {
	board = strings.Repeat(".", boardArea)
}

func printBoard(b string) {

	for i := 0; i < boardSize; i++ {
		firstIndex := i * boardSize
		lastIndex := i*boardSize + boardSize
		fmt.Println(b[firstIndex:lastIndex])
	}

}

func main() {

	fmt.Println("gomoku here")
	initBoard()
	printBoard(board)

	var x int
	//var y int

	fmt.Scan(&x)

}
