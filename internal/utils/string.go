package utils

import (
	runewidth "github.com/mattn/go-runewidth"
)

func StringsMaxLength(ss []string) int {
	max := 0
	for _, s := range ss {
		stringWidth := runewidth.StringWidth(s)
		if stringWidth > max {
			max = stringWidth + 1
		}
	}

	return max
}

func StringPadStart(s string, length int) string {
	stringWidth := runewidth.StringWidth(s)
	if stringWidth >= length {
		return s
	}

	paddingLength := length - int(stringWidth)
	padding := ""
	for i := 0; i < paddingLength; i++ {
		padding += " "
	}

	return padding + s
}

func StringPadEnd(s string, length int) string {
	stringWidth := runewidth.StringWidth(s)
	if stringWidth >= length {
		return s
	}

	paddingLength := length - int(stringWidth)
	padding := ""
	for i := 0; i < paddingLength; i++ {
		padding += " "
	}

	return s + padding
}
