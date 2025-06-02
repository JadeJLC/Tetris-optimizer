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

func Tetromino(file *os.File) []tetromino {
	var tetrominoes []tetromino

	scanner := bufio.NewScanner(file)
	var currentTetromino tetromino
	row := 0
	tetrominoCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			if row == 4 {
				tetrominoes = append(tetrominoes, currentTetromino)
				currentTetromino = tetromino{}
				row = 0
				tetrominoCount++
			} else if row > 0 {
				NotValid()
			}
			continue
		}

		if len(line) != 4 || row > 4 {
			NotValid()
		}

		for col, char := range line {
			currentTetromino[row][col] = char
		}
		row++
	}

	// After the loop, check if the last tetromino was complete
	if row == 4 {
		tetrominoes = append(tetrominoes, currentTetromino)
	} else if row > 0 {
		NotValid()
	}

	return tetrominoes
}

func Lettering(group []tetromino) []tetromino {
	for i := 0; i < len(group); i++ {
		for _, line := range group[i] {
			for _, char := range line {
				if char == '#' {
					char = rune(65 + i)
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
