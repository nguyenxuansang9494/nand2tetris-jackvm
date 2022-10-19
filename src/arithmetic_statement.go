package main

type ArithmeticStatement struct {
	Command string
}

func (stmt ArithmeticStatement) GetType() StatementType {
	return ArithmeticType
}
