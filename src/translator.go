package main

func TranslateMemoryStatement(stmt MemoryStatement) []string {
	if stmt.GetType() != MemoryType {
		panic("invalid argument")
	}

	return nil
}

func translatePushStatement(stmt MemoryStatement) []string {
	rs := make([]string, 9)
	if stmt.Segment == "LCL" || stmt.Segment == "Arg" || stmt.Segment == "this" || stmt.Segment == "that" {
		rs = append(rs, "@"+stmt.Segment)
		rs = append(rs, "D=M")
		rs = append(rs, "@"+stmt.Index)
		rs = append(rs, "A=D+A")
		rs = append(rs, "D=M")
		rs = append(rs, "@SP")
		rs = append(rs, "M=M+1")
		rs = append(rs, "A=M-1")
		rs = append(rs, "M=D")
	}
	return rs
}
