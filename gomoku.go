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

func main() {

	fmt.Println("gomoku here")
	initBoard()
	printBoard(board)

	var x int
	//var y int

	fmt.Scan(&x)

}
