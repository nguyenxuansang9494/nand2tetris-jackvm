package main

import (
	"bufio"
	"os"
)

func ReadFile(path string) *bufio.Scanner {
	file, err := os.Open(path)
	if err != nil {
		panic("failed to open file")
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	return fileScanner
}

func CreateFile(path string) *os.File {
	file, err := os.Create(path)
	if err != nil {
		panic("failed to create file")
	}
	return file
}
