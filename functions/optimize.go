package tetris

import (
	"os"
)

type square [][]rune
type tetromino [4][4]rune

func Optimizer(file *os.File) {
	tetrominos := Tetromino(file)

	n := CountPieces(file)
	finalSquare := MakeSquare(n)

	PrintSquare(n, finalSquare)
}
