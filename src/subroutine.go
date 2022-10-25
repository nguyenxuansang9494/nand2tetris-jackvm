package main

func SetMinusOne() []string {
	rs := make([]string, 4)
	rs = append(rs, "(SET_MINUS_ONE)")
	rs = append(rs, "@SP")
	rs = append(rs, "A=M-1")
	rs = append(rs, "M=-1")
	return rs
}

func SetZero() []string {
	rs := make([]string, 4)
	rs = append(rs, "(SET_ZERO)")
	rs = append(rs, "@SP")
	rs = append(rs, "A=M-1")
	rs = append(rs, "M=0")
	return rs
}
