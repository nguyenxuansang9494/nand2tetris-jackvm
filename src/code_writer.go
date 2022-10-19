package main

import "os"

var (
	openedFile *os.File = nil
	filePath   string   = ""
)

func WriteLine(path string, line string) {
	if path != filePath {
		if openedFile != nil {
			openedFile.Close()
		}
		filePath = path
		openedFile = CreateFile(path)
	}
	openedFile.WriteString(line)
}
