package main

import (
	"fmt"
	"strings"
)

func main() {
	boardSize := 8
	fmt.Println("gomoku here")

	// print empty board
	var emptyLine string
	emptyLine = strings.Repeat(".", boardSize)
	for i := 0; i < boardSize; i++ {
		fmt.Println(emptyLine)
	}
}
