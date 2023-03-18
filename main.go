package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/tazya/go-hack-assembler/pkg/assembler"
	"github.com/tazya/go-hack-assembler/pkg/parser"
	"github.com/tazya/go-hack-assembler/pkg/symbol_table"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var inputFilepath string
	var outputDirectory string

	flag.StringVar(&inputFilepath, "fileIn", "", "a string var")
	flag.StringVar(&outputDirectory, "dirOut", "", "a string var")
	flag.Parse()

	if inputFilepath == "" || outputDirectory == "" {
		fmt.Println("Usage: go-hack-assembler -fileIn=\"path/to/input.asm\" dirOut=\"path/to/output\"")
		return
	}

	codeLines := readFile(inputFilepath)

	st := symbol_table.New()
	p := parser.New(st)
	instructions, err := p.ParseCode(codeLines)

	if err != nil {
		fmt.Println(err)
		return
	}

	fileNameWithoutExt, _, _ := strings.Cut(filepath.Base(inputFilepath), ".")

	outputPath := outputDirectory + fileNameWithoutExt + ".hack"

	writeFile(outputPath, assembler.Assemble(instructions))

	fmt.Println("Successfully assembled! See your file:", outputPath)
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
