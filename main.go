package main

import "fmt"

var secondPlayer string
var firstPlayer string
var boardSize [3][3]string

func main() {
	firstPlayer, secondPlayer := names()
	fmt.Printf("Today's game is between %v and %v\n", firstPlayer, secondPlayer)
	moves()
}

func moves() {
	var row int
	var column int
	for i := 0; i < 9; i++ {
		if i%2 == 1 {
			fmt.Printf("It's the turn of %v\n", firstPlayer)
			fmt.Println("Please enter the row:")
			fmt.Scanln(&row)
			fmt.Println("Please enter the column:")
			fmt.Scanln(&column)
			fmt.Println("Please enter the column")
			if row > 3 || column > 3 || boardSize[row-1][column-1] != "" {
				fmt.Printf("Invalid input\n")
				i--
				fmt.Printf("Give again row and column, row and column need to be between 1-3 and cannot be in the same position as any other sign x or o \n")
				continue
			} else {
				boardSize[row-1][column-1] = "x"
			}
			printBoard(boardSize)
			if checkWinning() == true {
				break
			}
		} else {
			fmt.Printf("It's the turn of %v\n", secondPlayer)
			fmt.Println("Please enter the row:")
			fmt.Scanln(&row)
			fmt.Println("Please enter the column:")
			fmt.Scanln(&column)
			fmt.Println("Please enter the column")
			if row > 3 || column > 3 || boardSize[row-1][column-1] != "" {
				fmt.Printf("Invalid input\n")
				i--
				fmt.Printf("Give again row and column, row and column need to be between 1-3 and cannot be in the same position as any other sign x or o \n")
				continue
			} else {
				boardSize[row-1][column-1] = "o"
			}
			printBoard(boardSize)
			if checkWinning() == true {
				break
			}

		}

	}

}

func names() (string, string) {

	fmt.Println("Please enter your first player nickname:")
	fmt.Scanln(&firstPlayer)
	fmt.Println("Please enter your second player nickname:")
	fmt.Scanln(&secondPlayer)
	return firstPlayer, secondPlayer
}

func checkWinning() bool {
	if (boardSize[0][0] == boardSize[0][1] && boardSize[0][1] == boardSize[0][2] && boardSize[0][0] != "") ||
		(boardSize[1][0] == boardSize[1][1] && boardSize[1][1] == boardSize[1][2] && boardSize[1][0] != "") ||
		(boardSize[2][0] == boardSize[2][1] && boardSize[2][1] == boardSize[2][2] && boardSize[2][0] != "") ||
		(boardSize[0][0] == boardSize[1][0] && boardSize[1][0] == boardSize[2][0] && boardSize[0][0] != "") ||
		(boardSize[0][1] == boardSize[1][1] && boardSize[1][1] == boardSize[2][1] && boardSize[0][1] != "") ||
		(boardSize[0][2] == boardSize[1][2] && boardSize[1][2] == boardSize[2][2] && boardSize[0][2] != "") ||
		(boardSize[0][0] == boardSize[1][1] && boardSize[1][1] == boardSize[2][2] && boardSize[0][0] != "") ||
		(boardSize[0][2] == boardSize[1][1] && boardSize[1][1] == boardSize[2][0] && boardSize[0][2] != "") {
		fmt.Printf("We have the winner\n")
		return true
	}

	return false
}

func printBoard(boardSize [3][3]string) {
	for i := 0; i < len(boardSize); i++ {
		fmt.Println(boardSize[i])
	}
}
