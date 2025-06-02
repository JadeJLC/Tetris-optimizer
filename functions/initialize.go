package tetris

import (
	"bufio"
	"os"
	"strings"
)

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
