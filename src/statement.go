package main

type StatementType int

const (
	Memory StatementType = iota
	Arithmetic
)

type Statement interface {
	GetType() StatementType
}
