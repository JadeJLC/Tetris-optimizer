package tetris_test

import (
	"testing"
	tetris "tetris-optimizer/functions"
)

func makeTetrominoFromStrings(lines []string) tetris.Tetro {
	var t tetris.Tetro
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			t[i][j] = rune(lines[i][j])
		}
	}
	return t
}

func TestCheckGeometry_Valid(t *testing.T) {
	lines := []string{
		".#..",
		"##..",
		"#...",
		"....",
	}
	tet := makeTetrominoFromStrings(lines)
	if !tetris.CheckGeometry(tet) {
		t.Error("CheckGeometry failed for a valid tetromino")
	}
}

func TestCheckGeometry_Invalid(t *testing.T) {
	lines := []string{
		"#...",
		".#..",
		"..#.",
		"...#",
	}
	tet := makeTetrominoFromStrings(lines)
	if tetris.CheckGeometry(tet) {
		t.Error("CheckGeometry returned true for invalid tetromino")
	}
}

func TestMoveTetromino(t *testing.T) {
	lines := []string{
		"..#.",
		"..#.",
		"..#.",
		"..#.",
	}
	tet := makeTetrominoFromStrings(lines)

	moved := tetris.MoveTetromino(tet)

	// After moving, the '#' should be at column 0 (moved left)
	if moved[0][0] != '#' {
		t.Errorf("Expected moved tetromino to have '#' at (0,0), got %c", moved[0][0])
	}
}
