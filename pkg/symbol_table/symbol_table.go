package symbol_table

import (
	"fmt"
	"github.com/tazya/go-hack-assembler/pkg/instruction"
	"strconv"
)

var initSymbols = map[string]string{
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

type SymbolTable struct {
	symbols     map[string]string
	variableInc int
}

func New() *SymbolTable {
	st := &SymbolTable{
		symbols:     make(map[string]string, len(initSymbols)),
		variableInc: 16,
	}

	for k, v := range initSymbols {
		st.symbols[k] = v
	}

	return st
}

func (st *SymbolTable) Has(label string) bool {
	_, isExists := st.symbols[label]

	return isExists
}

func (st *SymbolTable) Get(label string) (*instruction.A, error) {
	mnemonic, isExists := st.symbols[label]

	if !isExists {
		return nil, nil
	}

	i, err := instruction.NewInstructionA(mnemonic)

	if err != nil {
		return nil, err
	}

	return i, nil
}

func (st *SymbolTable) Set(label string, a *instruction.A) {
	st.symbols[label] = a.GetMnemonic()
}

func (st *SymbolTable) NewLabel(label string, lineNumber int) error {
	instructionMnemonic := fmt.Sprintf("@%d", lineNumber)
	a, err := instruction.NewInstructionA(instructionMnemonic)

	if err != nil {
		return err
	}

	st.Set(label, a)

	return nil
}

func (st *SymbolTable) NewVariable(label string) (*instruction.A, error) {
	a, err := instruction.NewInstructionA("@" + strconv.Itoa(st.variableInc))

	if err != nil {
		return nil, err
	}

	st.symbols[label] = a.GetMnemonic()

	st.variableInc++

	return a, nil
}
