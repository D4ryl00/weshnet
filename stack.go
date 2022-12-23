package weshnet

import (
	"fmt"
	"runtime"
	"strings"
)

func getStackFrameInfo(frame int) (string, bool) {
	pc, file, line, ok := runtime.Caller(frame)
	if !ok {
		return "", false
	}

	rawFuncElems := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	function := rawFuncElems[len(rawFuncElems)-1]

	return fmt.Sprintf("file: %s, line: %d, function: %s", file, line, function), true
}

func printCleanStack() {
	stack := []string{}
	for i := 2; ; i++ {
		info, ok := getStackFrameInfo(i)
		if !ok {
			break
		}
		stack = append([]string{info}, stack...)
	}

	for _, frame := range stack {
		fmt.Println("TOTOTOTO   ", frame)
	}
	fmt.Println("TOTOTOTO")
}
