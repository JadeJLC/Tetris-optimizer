package tetris

import (
	"fmt"
	"os"
)

func NoFile() {
	fmt.Println("Error : Program needs a file to run.")
	os.Exit(1)
}

func NotValid() {
	fmt.Println("Error : Invalid format. Check the files and try again.")
	os.Exit(1)
}

func Error(e error) {
	if e != nil {
		fmt.Printf("Error : Something went wrong. Check the files and try again. %v\n", e)
		os.Exit(1)
	}
}
