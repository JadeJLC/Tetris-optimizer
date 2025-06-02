package tetris

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func PrintSquare(n int, finalSquare square) {
	for i := 0; i < n; i++ { // Outer loop for rows
		for j := 0; j < n; j++ { // Inner loop for columns
			fmt.Print(finalSquare[i][j])
		}
		fmt.Println() // New line after each row
	}
}
