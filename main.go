package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

var board = [][]string{
	[]string{"_", "_", "_"},
	[]string{"_", "_", "_"},
	[]string{"_", "_", "_"},
}

func main() {
	fmt.Println("Welcome to Tic-tac-go !")
	fmt.Println("2 players are gonna send coordinates where they want to place some X, or O character")
	fmt.Println("When 1 player align 3 identical symbols, game is over and he win !")
	fmt.Println("If no one succeed to aling 3 symbols and board is full, it's a draw !")
	fmt.Println("")

	displayBoard()

	var playerSymbol string

	turnNumber := 0
	for !isGameOver() && !isBoardFull() {
		if turnNumber%2==0 {
			playerSymbol = "X"
		} else {
			playerSymbol = "O"
		}
		x, y := inputGamer(playerSymbol)

		board[x][y] = playerSymbol

		displayBoard()
		turnNumber++
	}
}

func displayBoard() {
	for i:=0; i<len(board); i++ {
		fmt.Println(strings.Join(board[i],"|"))
	}
}

func inputGamer(playerSymbol string) (x , y int){

	x, y = askCoordinates(playerSymbol)

	for !isCoordinatesValid(x, y) {
		fmt.Println("Coordinates are invalid, retry")
		x, y = askCoordinates(playerSymbol)
	}
	return
}

func askCoordinates(playerSymbol string) (x, y int) { 
	fmt.Println("Player "+playerSymbol+" turn")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("X value: ")
	xStringValue, _ := reader.ReadString('\n')

	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Y value: ")
	yStringValue, _ := reader.ReadString('\n')

	x, _ = strconv.Atoi(strings.TrimSuffix(xStringValue, "\n"))
	y, _ = strconv.Atoi(strings.TrimSuffix(yStringValue, "\n"))

	x--
	y--
	return
}

func isGameOver() bool {
	return isHorizontalLineDone() || isVerticalLineDone() || isDiagonalDone()
}

func isHorizontalLineDone() bool {
	for i:=0; i<len(board); i++ {
		if board[0][i] != "_" && board[0][i] == board[1][i] && board[0][i] == board[2][i] {
			fmt.Println("Player "+board[0][i]+" wins !")
			return true
		}
	}
	return false
}

func isVerticalLineDone() bool {
	for i:=0; i<len(board[0]); i++ {
		if board[i][0] != "_" && board[i][0] == board[i][1] && board[i][0] == board[i][2] {
			fmt.Println("Player "+board[i][0]+" wins !")
			return true
		}
	}
	return false
}

func isDiagonalDone() bool {
	if(board[1][1]!="_" && ((board[0][0] == board[1][1] && board[0][0] == board[2][2]) || (board[0][2] == board[1][1] && board[0][2] == board[2][0]))) {
		fmt.Println("Player "+board[1][1]+" wins !")
		return true
	}
	return false
}

func isBoardFull() (result bool) {
	result = true
        for i:=0; i<len(board); i++ {
		result = result && !strings.ContainsAny(strings.Join(board[i],""),"_")
	}
	return
}

func isCoordinatesValid(x, y int) bool {
	return !isCoordinatesOutOfBound(x,y) && isCoordinatesEmpty(x,y)
}

func isCoordinatesOutOfBound(x, y int) bool {
	return x< 0 || x > len(board) || y<0 || y > len(board[0])
}

func isCoordinatesEmpty(x,y int) bool {
	return board[x][y] == "_"
}

