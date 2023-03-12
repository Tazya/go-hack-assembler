package instruction

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/tazya/go-hack-assembler/pkg/utils"
	"strconv"
	"strings"
)

type A struct {
	Value int
}

func (a *A) Assemble() string {
	return fmt.Sprintf("%s%015b", a.getOpcode(), a.Value)
}

func (a *A) getOpcode() string {
	return "0"
}

func NewInstructionA(codeLine string) (*A, error) {
	i := &A{}
	preparedLine := utils.RemoveComments(strings.TrimPrefix(codeLine, "@"))

	if len(preparedLine) == 0 {
		return i, errors.New("syntax error. empty A instruction")
	}

	instructionValue, err := strconv.Atoi(preparedLine)

	if err != nil {
		return i, errors.New("syntax error. the A instruction must be integer")
	}

	if instructionValue > max15bitValue {
		return i, errors.New("parsing error. the A instruction value must be less than " + strconv.Itoa(max15bitValue))
	}

	i.Value = instructionValue

	return i, nil
}
