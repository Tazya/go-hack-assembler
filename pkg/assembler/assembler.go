package assembler

import (
	"github.com/tazya/go-hack-assembler/pkg/instruction"
)

func Assemble(instructions []instruction.Instruction) string {
	result := ""
	lastIndex := len(instructions) - 1

	for i, inst := range instructions {
		result = result + inst.Assemble()

		if i != lastIndex {
			result = result + "\n"
		}
	}

	return result
}
