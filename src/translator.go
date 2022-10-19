package main

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
	}
	return rs
}
