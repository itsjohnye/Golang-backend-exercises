package main

import (
	"fmt"
	"time"
)

func main() {

	var board [8][8]rune

	board[0] = [...]rune{ //第一行
		'r', 'n', 'b', 'q', 'k', 'b', 'n', 'r', //rune要用单引号
	}

	board[7] = [...]rune{
		'R', 'N', 'B', 'W', 'K', 'B', 'N', 'R',
	}

	for column := range board[1] {
		board[1][column] = 'p'
		board[6][column] = 'P'
	}

	display(board)

}

func display(board [8][8]rune) {
	for _, row := range board { //行
		for _, column := range row { //列
			if column == 0 {
				fmt.Print("  ")
			} else {
				fmt.Printf("%c ", column)
			}
			time.Sleep(1 * time.Second) //这里用sleep函数便于看出实现的步骤
		}
		fmt.Println()
	}
}
