package tetris

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type square [][]rune
type tetromino [4][4]rune

func CountPieces(file *os.File) int {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	n := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue // Skip empty lines
		}
		n += 1
	}

	if n%4 != 0 || n == 0 {
		NotValid()
	}
	n /= 4
	return n
}

func MakeSquare(n int) square {
	square := make([][]rune, n)

	for i := range square {
		square[i] = make([]rune, n)
		for j := range square[i] {
			square[i][j] = ' '
		}
	}
	return square
}

func Lettering(group []tetromino) []tetromino {
	for i := 0; i < len(group); i++ {
		for rowIdx := 0; rowIdx < 4; rowIdx++ {
			for colIdx := 0; colIdx < 4; colIdx++ {
				if group[i][rowIdx][colIdx] == '#' {
					group[i][rowIdx][colIdx] = rune(65 + i)
				}
			}
		}
	}
	return group
}

func PrintSquare(n int, finalSquare square) {
	for i := 0; i < n; i++ { // Outer loop for rows
		for j := 0; j < n; j++ { // Inner loop for columns
			fmt.Print(finalSquare[i][j])
		}
		fmt.Println() // New line after each row
	}
}

func Optimizer(file *os.File) {
	tetrominos := Tetromino(file)

	n := CountPieces(file)
	finalSquare := MakeSquare(n)

	PrintSquare(n, finalSquare)
}
