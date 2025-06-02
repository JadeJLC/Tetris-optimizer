package tetris

import (
	"os"
)

type square [][]rune
type tetromino [4][4]rune

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

func Optimizer(file *os.File) {
	tetrominos := Tetromino(file)

	n := CountPieces(file)
	finalSquare := MakeSquare(n)

	PrintSquare(n, finalSquare)
}
