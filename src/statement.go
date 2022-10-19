package main

type StatementType int

const (
	MemoryType StatementType = iota
	ArithmeticType
)

type Statement interface {
	GetType() StatementType
}
