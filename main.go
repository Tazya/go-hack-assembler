package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/tazya/go-hack-assembler/pkg/assembler"
	"github.com/tazya/go-hack-assembler/pkg/parser"
	"os"
)

func main() {
	var inputFilepath string
	var outputFilepath string

	flag.StringVar(&inputFilepath, "fileIn", "", "a string var")
	flag.StringVar(&outputFilepath, "fileOut", "", "a string var")
	flag.Parse()

	if inputFilepath == "" || outputFilepath == "" {
		fmt.Println("Usage: go-hack-assembler -input=\"path/to/input.asm\" output=\"path/to/output\"")
		return
	}

	codeLines := readFile(inputFilepath)
	instructions, err := parser.ParseCode(codeLines)

	if err != nil {
		fmt.Println(err)
		return
	}

	writeFile(outputFilepath, assembler.Assemble(instructions))
}

func readFile(filepath string) []string {
	var fileLines []string

	f, err := os.Open(filepath)

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	f.Close()

	return fileLines
}

func writeFile(filepath string, assembledCode string) {
	f, err := os.Create(filepath)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err = f.WriteString(assembledCode)

	if err != nil {
		panic(err)
	}
}
