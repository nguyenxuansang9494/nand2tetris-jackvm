package main

type MemoryStatement struct {
	Command string
	Segment string
	Index   string
}

func (stmt MemoryStatement) GetType() StatementType {
	return MemoryType
}
