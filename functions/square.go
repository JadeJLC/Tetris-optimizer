package tetris

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CountPieces(file *os.File) int {
	file.Seek(0, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	lines := []string{}
	lineNum := 0
	for scanner.Scan() {
		text := scanner.Text()
		lineNum++
		trimmed := strings.TrimSpace(text)
		if trimmed != "" {
			if len(trimmed) != 4 {
				NotValid()
			}
			lines = append(lines, trimmed)
		}
	}

	if len(lines) == 0 {
		NotValid()
	}

	if len(lines)%4 != 0 {
		NotValid()
	}

	return len(lines) / 4
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

func PrintSquare(n int, finalSquare square) {
	for i := 0; i < n; i++ { // Outer loop for rows
		for j := 0; j < n; j++ { // Inner loop for columns
			if finalSquare[i][j] == 32 {
				finalSquare[i][j] = '.'
			}
			fmt.Printf("%c", finalSquare[i][j])
		}
		fmt.Println() // New line after each row
	}
}
