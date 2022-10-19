package main

type MemoryStatement struct {
	Command string
	Segment string
	Index   int
}

func (stmt MemoryStatement) GetType() StatementType {
	return MemoryType
}
