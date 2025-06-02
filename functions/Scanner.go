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

	if row == 4 {
		tetrominoes = append(tetrominoes, currentTetromino)
	} else if row > 0 {
		NotValid()
	}

	for i := 0; i < len(tetrominoes); i++ {
		tetrominoes[i] = MoveTetromino(tetrominoes[i])
	}

	return tetrominoes
}

func MoveTetromino(t tetromino) tetromino {
	minRow := 4
	minCol := 4

	// Cherche l'emplacement de la première case de la pièce
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			if t[r][c] == '#' {
				if r < minRow {
					minRow = r
				}
				if c < minCol {
					minCol = c
				}
			}
		}
	}
	if minRow == 4 && minCol == 4 {
		return t
	}

	// Crée un tetromino vide
	var newTetro tetromino
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			newTetro[r][c] = '.'
		}
	}

	// Replace la pièce dans le tetromino vide
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			if t[r][c] == '#' {
				newRow := r - minRow
				newCol := c - minCol
				newTetro[newRow][newCol] = '#'
			}
		}
	}

	return newTetro
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
