package main

import "fmt"

func main() {
	const (
		FIRST_GLYPH     = 33  // '!'
		LAST_GLYPH      = 126 // '~'
		GLYPHS_PER_LINE = 72
	)

	cursor := FIRST_GLYPH

	var i int
	for {
		i = cursor
		for range GLYPHS_PER_LINE {
			fmt.Print(string(rune(i)))
			if i == LAST_GLYPH {
				i = FIRST_GLYPH
			} else {
				i++
			}
		}
		fmt.Println()
		if cursor == LAST_GLYPH {
			cursor = FIRST_GLYPH
		} else {
			cursor++
		}
	}
}
