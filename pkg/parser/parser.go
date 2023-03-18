package parser

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/tazya/go-hack-assembler/pkg/instruction"
	"github.com/tazya/go-hack-assembler/pkg/symbol"
	"github.com/tazya/go-hack-assembler/pkg/utils"
	"strings"
)

func ParseCode(codeLines []string) ([]instruction.Instruction, error) {
	var instructions []instruction.Instruction

	lines, err := passRegisterLabels(codeLines)

	if err != nil {
		return instructions, err
	}

	for lineNumber, codeLine := range lines {
		i, err := parseLine(codeLine)

		if err != nil {
			err = errors.Wrapf(err, "Error on line %d", lineNumber+1)
			return instructions, err
		}

		if i != nil {
			instructions = append(instructions, i)
		}
	}

	return instructions, nil
}

func parseLine(codeLine string) (instruction.Instruction, error) {
	preparedLine := strings.Trim(codeLine, " ")
	var i instruction.Instruction
	var err error

	if preparedLine == "" {
		return nil, nil
	}

	if strings.HasPrefix(preparedLine, "//") {
		return nil, nil
	}

	if strings.HasPrefix(preparedLine, "@") {
		if symbol.Has(strings.TrimPrefix(preparedLine, "@")) {
			i, err = symbol.Get(strings.TrimPrefix(preparedLine, "@"))
		} else {
			i, err = instruction.NewInstructionA(preparedLine)
		}
	} else {
		i, err = instruction.NewInstructionC(preparedLine)
	}

	if err != nil {
		return nil, err
	}

	return i, nil
}

func passRegisterLabels(codeLines []string) ([]string, error) {
	codeLinesWithoutLabels := make([]string, len(codeLines))
	currentLine := 1

	for _, line := range codeLines {
		preparedLine := utils.RemoveComments(line)

		if preparedLine == "" {
			continue
		}

		if !strings.HasPrefix(preparedLine, "(") {
			codeLinesWithoutLabels = append(codeLinesWithoutLabels, preparedLine)
			currentLine++

			continue
		}

		labelName := strings.Trim(preparedLine, "()")

		if !symbol.Has(labelName) {
			instructionMnemonic := fmt.Sprintf("@%d", currentLine)
			a, err := instruction.NewInstructionA(instructionMnemonic)

			if err != nil {
				return codeLinesWithoutLabels, err
			}

			codeLinesWithoutLabels = append(codeLinesWithoutLabels, "")

			symbol.Set(labelName, a)
		}
	}

	return codeLinesWithoutLabels, nil
}
