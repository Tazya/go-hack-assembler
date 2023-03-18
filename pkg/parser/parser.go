package parser

import (
	"github.com/pkg/errors"
	"github.com/tazya/go-hack-assembler/pkg/instruction"
	"github.com/tazya/go-hack-assembler/pkg/symbol_table"
	"github.com/tazya/go-hack-assembler/pkg/utils"
	"strings"
)

type Parser struct {
	symbolTable *symbol_table.SymbolTable
}

func New(st *symbol_table.SymbolTable) *Parser {
	return &Parser{
		symbolTable: st,
	}
}

func (p *Parser) ParseCode(codeLines []string) ([]instruction.Instruction, error) {
	lines, err := p.parseLabels(codeLines)

	if err != nil {
		return nil, err
	}

	instructions, err := p.parseInstructions(lines)

	return instructions, err
}

func (p *Parser) parseLabels(codeLines []string) ([]string, error) {
	codeLinesWithoutLabels := make([]string, 0, len(codeLines))
	instructionIndex := 0

	for _, line := range codeLines {
		preparedLine := utils.RemoveComments(line)

		if preparedLine == "" {
			codeLinesWithoutLabels = append(codeLinesWithoutLabels, preparedLine)

			continue
		}

		if !strings.HasPrefix(preparedLine, "(") {
			codeLinesWithoutLabels = append(codeLinesWithoutLabels, preparedLine)
			instructionIndex++

			continue
		}

		labelName := strings.Trim(preparedLine, "()")

		if !p.symbolTable.Has(labelName) {
			err := p.symbolTable.NewLabel(labelName, instructionIndex)

			if err != nil {
				return codeLinesWithoutLabels, err
			}
		}

		codeLinesWithoutLabels = append(codeLinesWithoutLabels, "")
	}

	return codeLinesWithoutLabels, nil
}

func (p *Parser) parseInstructions(codeLines []string) ([]instruction.Instruction, error) {
	var instructions []instruction.Instruction

	for lineNumber, codeLine := range codeLines {
		i, err := p.parseLine(codeLine)

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

func (p *Parser) parseLine(codeLine string) (instruction.Instruction, error) {
	var i instruction.Instruction
	var err error

	if codeLine == "" {
		return nil, nil
	}

	if strings.HasPrefix(codeLine, "@") {
		value := strings.TrimPrefix(codeLine, "@")

		if utils.IsNumeric(value) {
			i, err = instruction.NewInstructionA(codeLine)

			return i, err
		}

		if p.symbolTable.Has(value) {
			i, err = p.symbolTable.Get(strings.TrimPrefix(codeLine, "@"))
		} else {
			return p.symbolTable.NewVariable(value)
		}
	} else {
		i, err = instruction.NewInstructionC(codeLine)
	}

	return i, err
}
