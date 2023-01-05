package utils

import (
	"fmt"
	"unicode/utf8"
)

func StringsMaxLength(ss []string) int {
	max := 0
	for _, s := range ss {
		if utf8.RuneCount([]byte(s)) > max {
			max = utf8.RuneCount([]byte(s))
		}
	}

	return max
}

func StringPadStart(s string, length int) string {
	return fmt.Sprintf("%"+fmt.Sprintf("%d", length)+"s", s)
}

func StringPadEnd(s string, length int) string {
	return fmt.Sprintf("%-"+fmt.Sprintf("%d", length)+"s", s)
}
