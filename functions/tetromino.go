package tetris

import (
	"bufio"
	"os"
	"strings"
)

func CheckGeometry(t Tetro) bool {
	visited := [4][4]bool{}
	queue := [][2]int{{0, 0}}

	if t[0][0] == '#' {
		visited[0][0] = true
	} else {
		for i := 1; i < 4; i++ {
			if t[i][0] == '#' {
				visited[i][0] = true
				queue = [][2]int{{i, 0}}
				break
			} else if t[0][i] == '#' {
				visited[0][i] = true
				queue = [][2]int{{i, 0}}
				break
			} else {
				continue
			}
		}
	}
	counter := 1

	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // right, down, left, up

	for i := 0; i < len(queue); i++ {
		x, y := queue[i][0], queue[i][1]
		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			if nx >= 0 && ny >= 0 && nx < 4 && ny < 4 &&
				!visited[nx][ny] && t[nx][ny] == '#' {
				visited[nx][ny] = true
				queue = append(queue, [2]int{nx, ny})
				counter++
			}
		}
	}

	return counter == 4
}

func Tetromino(file *os.File) []Tetro {
	var tetrominoes []Tetro

	scanner := bufio.NewScanner(file)
	var currentTetromino Tetro
	row := 0
	tetrominoCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			if row == 4 {
				tetrominoes = append(tetrominoes, currentTetromino)
				currentTetromino = Tetro{}
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

		hashnumber := 0

		for col, char := range line {
			if char == '#' {
				hashnumber += 1
			}
			if hashnumber > 4 {
				NotValid()
			}
			currentTetromino[row][col] = char
		}
		row++
	}

	if row == 4 {
		tetrominoes = append(tetrominoes, currentTetromino)
	} else {
		NotValid()
	}

	for i := 0; i < len(tetrominoes); i++ {
		tetrominoes[i] = MoveTetromino(tetrominoes[i])
	}

	return tetrominoes
}

func MoveTetromino(t Tetro) Tetro {
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
	var newTetro Tetro

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

	if !CheckGeometry(newTetro) {
		NotValid()
	}

	return newTetro
}

func Lettering(group []Tetro) []Tetro {
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
