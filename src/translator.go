package main

import "strconv"

var segmentMap map[string]string = map[string]string{
	"argument": "ARG",
	"local":    "LCL",
	"this":     "this",
	"that":     "that",
}

func TranslateMemoryStatement(stmt MemoryStatement) []string {
	if stmt.Command == "push" {
		return translatePushStatement(stmt)
	}
	return translatePopStatement(stmt)
}

func translatePopStatement(stmt MemoryStatement) []string {
	return nil
}

func translatePushStatement(stmt MemoryStatement) []string {
	rs := make([]string, 9)
	if stmt.Segment == "local" || stmt.Segment == "argument" || stmt.Segment == "this" || stmt.Segment == "that" {
		rs = append(rs, "@"+segmentMap[stmt.Segment])
		rs = append(rs, "D=M")
		rs = append(rs, "@"+stmt.Index)
		rs = append(rs, "A=D+A")
		rs = append(rs, "D=M")
		rs = append(rs, "@SP")
		rs = append(rs, "M=M+1")
		rs = append(rs, "A=M-1")
		rs = append(rs, "M=D")
	} else if stmt.Segment == "constant" {
		rs = append(rs, "@"+stmt.Index)
		rs = append(rs, "D=A")
		rs = append(rs, "@SP")
		rs = append(rs, "M=M+1")
		rs = append(rs, "A=M-1")
		rs = append(rs, "M=D")
	} else if stmt.Segment == "static" {
		rs = append(rs, "@"+HandledFile+"."+stmt.Index)
		rs = append(rs, "D=M")
		rs = append(rs, "@SP")
		rs = append(rs, "M=M+1")
		rs = append(rs, "A=M-1")
		rs = append(rs, "M=D")
	} else if stmt.Segment == "temp" {
		stmtIndex, _ := strconv.Atoi(stmt.Index)
		index := strconv.Itoa(stmtIndex + 5)
		rs = append(rs, "@"+index)
		rs = append(rs, "D=M")
		rs = append(rs, "@SP")
		rs = append(rs, "M=M+1")
		rs = append(rs, "A=M-1")
		rs = append(rs, "M=D")
	} else if stmt.Segment == "pointer" {
		switch stmt.Index {
		case "0":
			rs = append(rs, "@this")
			rs = append(rs, "D=M")
			rs = append(rs, "@SP")
			rs = append(rs, "M=M+1")
			rs = append(rs, "A=M-1")
			rs = append(rs, "M=D")
		case "1":
			rs = append(rs, "@that")
			rs = append(rs, "D=M")
			rs = append(rs, "@SP")
			rs = append(rs, "M=M+1")
			rs = append(rs, "A=M-1")
			rs = append(rs, "M=D")
		}
	}
	return rs
}
