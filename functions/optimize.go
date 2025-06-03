package tetris

import (
	"os"
)

type square [][]rune
type tetromino [4][4]rune

func CheckPiece(t tetromino, s square, x, y int) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if t[i][j] != '.' {
				sx := x + i
				sy := y + j
				if sx >= len(s) || sy >= len(s) || s[sx][sy] != ' ' {
					return false
				}
			}
		}
	}
	return true
}

func PlacePiece(t tetromino, s square, x, y int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if t[i][j] != '.' {
				s[x+i][y+j] = t[i][j]
			}
		}
	}
}

func RemovePiece(t tetromino, s square, x, y int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if t[i][j] != '.' {
				s[x+i][y+j] = ' '
			}
		}
	}
}

func PlaceThemAll(tetrominos []tetromino, s square, n int) bool {
	if len(tetrominos) == 0 {
		return true
	}

	t := tetrominos[0]
	size := len(s)

	for x := 0; x <= size-n; x++ {
		for y := 0; y <= size-n; y++ {
			if CheckPiece(t, s, x, y) {
				PlacePiece(t, s, x, y)
				if PlaceThemAll(tetrominos[1:], s, n) {
					return true
				}
				RemovePiece(t, s, x, y)
			}
		}
	}

	return false
}

func Optimizer(file *os.File) {
	tetrominos := Tetromino(file)
	tetrominos = Lettering(tetrominos)

	n := CountPieces(file)
	size := 1
	for size*size < n*4 {
		size++
	}

	var finalSquare square
	for {
		finalSquare = MakeSquare(size)
		if PlaceThemAll(tetrominos, finalSquare, n) {
			break
		}
		size++
	}

	PrintSquare(size, finalSquare)
}
