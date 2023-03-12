package parser

import (
	"github.com/pkg/errors"
	"github.com/tazya/go-hack-assembler/pkg/instruction"
	"strings"
)

func ParseCode(codeLines []string) ([]instruction.Instruction, error) {
	var instructions []instruction.Instruction

	for lineNumber, codeLine := range codeLines {
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
		i, err = instruction.NewInstructionA(preparedLine)
	} else {
		i, err = instruction.NewInstructionC(preparedLine)
	}

	if err != nil {
		return nil, err
	}

	return i, nil
}
