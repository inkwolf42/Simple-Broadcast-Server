package client

import (
	"fmt"
	"os"

	"golang.org/x/term"
	"inkwolf.com/broadcast/internal/utils"
)

var messages Stack[string]

func getSize() (int, int) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	utils.ExitIfErr(err)
	return width, height
}

func ClearTerminal() {
	fmt.Print("\033[2J")
}

func MoveCursorTo(x, y int) {
	fmt.Printf("\033[%d;%dH", x, y)
}

func drawRect(startRow, startCol, width, height int) {
	const (
		topLeft     = "┌"
		topRight    = "┐"
		bottomLeft  = "└"
		bottomRight = "┘"
		horiz       = "─"
		vert        = "│"
	)

	if width < 2 || height < 2 {
		return
	}
	MoveCursorTo(startRow, startCol)
	fmt.Print(topLeft)
	for i := 0; i < width-2; i++ {
		fmt.Print(horiz)
	}
	fmt.Print(topRight)

	for r := 1; r < height-1; r++ {
		MoveCursorTo(startRow+r, startCol)
		fmt.Print(vert)

		MoveCursorTo(startRow+r, startCol+width-1)
		fmt.Print(vert)
	}

	for r := 1; r < height-1; r++ {
		MoveCursorTo(startRow+r, startCol+1)
		for i := 1; i < width-2; i++ {
			fmt.Print(" ")
		}
	}

	MoveCursorTo(startRow+height-1, startCol)
	fmt.Print(bottomLeft)
	for i := 0; i < width-2; i++ {
		fmt.Print(horiz)
	}
	fmt.Print(bottomRight)
}

func ReDraw() {
	ClearTerminal()

	MoveCursorTo(0, 0)

	messages.EachBackwards(func(item string) {
		fmt.Print(item)
	})

	println("")
	println("")
	println("")
	println("")

	w, h := getSize()

	drawRect(h-2, 1, w, 3)

	MoveCursorTo(h-1, 3)

	fmt.Print(" > ")
}
