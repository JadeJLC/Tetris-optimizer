package main

import (
	"fmt"
	"os"
	tetris "tetris-optimizer/functions"
)

type tetromino [4][4]rune
type square [][]rune

func main() {
	if len(os.Args) != 2 {
		tetris.NoFile()
	}

	file, err := os.Open(os.Args[1])
	tetris.Error(err)
	tetris.Scanner(file)
	defer file.Close()
	finalSquare := tetris.MakeSquare(tetris.CountPieces(file))

	fmt.Println(finalSquare)
}
