package main

import (
	"strconv"
	"strings"
)

func Parse(line string) Statement {
	lexers := strings.Split(line, " ")
	if len(lexers) == 3 && (lexers[0] == "pop" || lexers[0] == "push") {
		index, _ := strconv.Atoi(lexers[2])
		return MemoryStatement{
			Command: lexers[0],
			Segment: lexers[1],
			Index:   index,
		}
	} else if len(lexers) == 1 {
		return ArithmeticStatement{Command: lexers[0]}
	}
	return ArithmeticStatement{}
}
