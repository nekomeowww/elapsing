package utils

import (
	"runtime"
	"strings"

	"github.com/samber/lo"
)

func FunctionNameOfCaller(skip int) string {
	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		return ""
	}

	fullFunctionName := runtime.FuncForPC(pc).Name()
	split := strings.Split(fullFunctionName, "/")
	return lo.Ternary(
		len(split) == 0,     // if
		"",                  // then
		split[len(split)-1], // else
	)
}
