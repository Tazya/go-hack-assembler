package symbol

import (
	"github.com/tazya/go-hack-assembler/pkg/instruction"
	"strconv"
)

var symbols = map[string]string{
	"R0":     "@0",
	"R1":     "@1",
	"R2":     "@2",
	"R3":     "@3",
	"R4":     "@4",
	"R5":     "@5",
	"R6":     "@6",
	"R7":     "@7",
	"R8":     "@8",
	"R9":     "@9",
	"R10":    "@10",
	"R11":    "@11",
	"R12":    "@12",
	"R13":    "@13",
	"R14":    "@14",
	"R15":    "@15",
	"SCREEN": "@16384",
	"KBD":    "@24576",
	"SP":     "@0",
	"LCL":    "@1",
	"ARG":    "@2",
	"THIS":   "@3",
	"THAT":   "@4",
}

var variableInc = 16

func Has(label string) bool {
	_, isExists := symbols[label]

	return isExists
}

func Get(label string) (*instruction.A, error) {
	mnemonic, isExists := symbols[label]

	if !isExists {
		return nil, nil
	}

	i, err := instruction.NewInstructionA(mnemonic)

	if err != nil {
		return nil, err
	}

	return i, nil
}

func Set(label string, a *instruction.A) {
	symbols[label] = a.GetMnemonic()
}

func NewVariable(label string) (*instruction.A, error) {
	a, err := instruction.NewInstructionA("@" + strconv.Itoa(variableInc))

	if err != nil {
		return nil, err
	}

	symbols[label] = a.GetMnemonic()

	variableInc++

	return a, nil
}
