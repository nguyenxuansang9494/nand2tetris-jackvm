package main

func TranslateArithmeticStatement(stmt ArithmeticStatement) []string {
	return nil
}

func translateUnaryOperation(stmt ArithmeticStatement) []string {
	rs := make([]string, 3)
	if stmt.Command == "not" {
		rs = append(rs, "@SP")
		rs = append(rs, "A=M-1")
		rs = append(rs, "M=!M")
	} else if stmt.Command == "neg" {
		rs = append(rs, "@SP")
		rs = append(rs, "A=M-1")
		rs = append(rs, "M=-M")
	}
	return rs
}

func translateBinaryOperation(stmt ArithmeticStatement) []string {
	rs := make([]string, 6)
	rs = append(rs, "@SP")
	rs = append(rs, "AM=M-1")
	rs = append(rs, "D=M")
	rs = append(rs, "@SP")
	rs = append(rs, "A=M-1")
	if stmt.Command == "add" {
		rs = append(rs, "M=D+M")
	} else if stmt.Command == "sub" {
		rs = append(rs, "M=M-D")
	} else if stmt.Command == "and" {
		rs = append(rs, "M=D&M")
	} else if stmt.Command == "or" {
		rs = append(rs, "M=D|M")
	}

	return rs
}
