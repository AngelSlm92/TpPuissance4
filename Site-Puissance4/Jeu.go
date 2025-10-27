package main

const rows = 6
const cols = 7

var board [rows][cols]int
var current = 1
var winner = 0
var draw = false

func reset() {
	board = [rows][cols]int{}
	current = 1
	winner = 0
	draw = false
}

func drop(col int, player int) bool {
	if col < 0 || col >= cols {
		return false
	}
	for r := rows - 1; r >= 0; r-- {
		if board[r][col] == 0 {
			board[r][col] = player
			return true
		}
	}
	return false
}

func checkWin() int {

	for r := 0; r < rows; r++ {
		for c := 0; c <= cols-4; c++ {
			p := board[r][c]
			if p != 0 && p == board[r][c+1] && p == board[r][c+2] && p == board[r][c+3] {
				return p
			}
		}
	}

	for c := 0; c < cols; c++ {
		for r := 0; r <= rows-4; r++ {
			p := board[r][c]
			if p != 0 && p == board[r+1][c] && p == board[r+2][c] && p == board[r+3][c] {
				return p
			}
		}
	}

	for r := 0; r <= rows-4; r++ {
		for c := 0; c <= cols-4; c++ {
			p := board[r][c]
			if p != 0 && p == board[r+1][c+1] && p == board[r+2][c+2] && p == board[r+3][c+3] {
				return p
			}
		}
	}
	for r := 3; r < rows; r++ {
		for c := 0; c <= cols-4; c++ {
			p := board[r][c]
			if p != 0 && p == board[r-1][c+1] && p == board[r-2][c+2] && p == board[r-3][c+3] {
				return p
			}
		}
	}
	return 0
}

func isFull() bool {
	for c := 0; c < cols; c++ {
		if board[0][c] == 0 {
			return false
		}
	}
	return true
}
