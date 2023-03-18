package parser

import (
	"bufio"
	"github.com/tazya/go-hack-assembler/pkg/assembler"
	"github.com/tazya/go-hack-assembler/pkg/symbol_table"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestParseCode(t *testing.T) {
	tests := []struct {
		name             string
		codeFilePath     string
		expectedFilePath string
		wantErr          bool
	}{
		{
			name:             "Successfully assembled: add.asm",
			codeFilePath:     "tests/fixtures/add.asm",
			expectedFilePath: "tests/fixtures/add.hack",
			wantErr:          false,
		},
		{
			name:         "Error while assembly: add.asm",
			codeFilePath: "tests/fixtures/add_with_error.asm",
			wantErr:      true,
		},
		{
			name:             "Successfully assembled: Max.asm",
			codeFilePath:     "tests/fixtures/Max.asm",
			expectedFilePath: "tests/fixtures/Max.hack",
			wantErr:          false,
		},
		{
			name:         "Error while assembly: Max.asm",
			codeFilePath: "tests/fixtures/Max_with_error.asm",
			wantErr:      true,
		},
		{
			name:             "Successfully assembled: Rect.asm",
			codeFilePath:     "tests/fixtures/Rect.asm",
			expectedFilePath: "tests/fixtures/Rect.hack",
			wantErr:          false,
		},
		{
			name:             "Successfully assembled: Pong.asm",
			codeFilePath:     "tests/fixtures/Pong.asm",
			expectedFilePath: "tests/fixtures/Pong.hack",
			wantErr:          false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := symbol_table.New()
			p := New(st)
			codeLines := readLinesFromFile(tt.codeFilePath)
			instructions, err := p.ParseCode(codeLines)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			expected := readTextFromFile(tt.expectedFilePath)
			assembled := assembler.Assemble(instructions)

			if assembled != expected {
				t.Errorf("ParseCode() = %v, want from file %v", assembled, tt.expectedFilePath)
			}
		})
	}
}

func readLinesFromFile(relativePath string) []string {
	var fileLines []string

	rootPath, _ := filepath.Abs("./../../")
	f, err := os.Open(rootPath + "/" + relativePath)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines
}

func readTextFromFile(relativePath string) string {
	rootPath, _ := filepath.Abs("./../../")
	file, err := os.Open(rootPath + "/" + relativePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	bytes, _ := io.ReadAll(file)

	return string(bytes)
}
