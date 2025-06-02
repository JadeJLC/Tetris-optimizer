package main

import (
	"os"
	tetris "tetris-optimizer/functions"
)

type tetromino [4][4]rune
type square [][]rune

func main() {
	if len(os.Args) != 2 {
		tetris.NoFile()
	}

	// Ouverture du fichier
	file, err := os.Open(os.Args[1])
	tetris.Error(err)
	defer file.Close()

	tetris.Optimizer(file)
}
