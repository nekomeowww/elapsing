package utils

import (
	"golang.org/x/text/width"
)

func StringWidth(s string) int {
	size := 0
	for _, runeValue := range s {
		p := width.LookupRune(runeValue)
		if p.Kind() == width.EastAsianWide {
			size += 2
			continue
		}
		if p.Kind() == width.EastAsianNarrow {
			size += 1
			continue
		}
	}

	return size
}

func StringsMaxLength(ss []string) int {
	max := 0
	for _, s := range ss {
		if StringWidth(s) > max {
			max = StringWidth(s) + 1
		}
	}

	return max
}

func StringPadStart(s string, length int) string {
	if StringWidth(s) >= length {
		return s
	}

	paddingLength := length - int(StringWidth(s))
	padding := ""
	for i := 0; i < paddingLength; i++ {
		padding += " "
	}

	return padding + s
}

func StringPadEnd(s string, length int) string {
	if StringWidth(s) >= length {
		return s
	}

	paddingLength := length - int(StringWidth(s))
	padding := ""
	for i := 0; i < paddingLength; i++ {
		padding += " "
	}

	return s + padding
}
