package instruction

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/tazya/go-hack-assembler/pkg/utils"
	"strings"
)

var destRepresentations = map[string]string{
	"M":   "001",
	"D":   "010",
	"DM":  "011",
	"A":   "100",
	"AM":  "101",
	"AD":  "110",
	"ADM": "111",
}

var compRepresentations = map[string]string{
	"0":   "0101010",
	"1":   "0111111",
	"-1":  "0111010",
	"D":   "0001100",
	"A":   "0110000",
	"M":   "1110000",
	"!D":  "0001101",
	"!A":  "0110001",
	"!M":  "1110001",
	"-D":  "0001111",
	"-A":  "0110011",
	"-M":  "1110011",
	"D+1": "0110011",
	"A+1": "0110111",
	"M+1": "1110111",
	"D-1": "0001110",
	"A-1": "0110010",
	"M-1": "1110010",
	"D+A": "0000010",
	"D+M": "1000010",
	"D-A": "0010011",
	"D-M": "1010011",
	"A-D": "0000111",
	"M-D": "1000111",
	"D&A": "0000000",
	"D&M": "1000000",
	"D|A": "0010101",
	"D|M": "1010101",
}

var jumpRepresentations = map[string]string{
	"JGT": "001",
	"JEQ": "010",
	"JGE": "011",
	"JLT": "100",
	"JNE": "101",
	"JLE": "110",
	"JMP": "111",
}

type C struct {
	Dest string
	Comp string
	Jump string
}

func (c *C) Assemble() string {
	return fmt.Sprintf(
		"%s11%s%s%s",
		c.getOpcode(),
		c.representDest(),
		c.representComp(),
		c.representJump(),
	)
}

func (c *C) getOpcode() string {
	return "1"
}

func NewInstructionC(codeLine string) (*C, error) {
	c := &C{}
	preparedLine := utils.RemoveComments(strings.Replace(codeLine, " ", "", -1))

	if len(preparedLine) == 0 {
		return c, errors.New("syntax error. empty C instruction")
	}

	var err error

	c.Dest, err = c.parseDestination(preparedLine)
	if err != nil {
		return nil, err
	}

	c.Comp, err = c.parseComputation(preparedLine)
	if err != nil {
		return nil, err
	}

	c.Jump, err = c.parseJump(preparedLine)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *C) parseDestination(codeLine string) (string, error) {
	dest, _, hasDest := strings.Cut(codeLine, "=")

	if !hasDest {
		return "", nil
	}

	_, isExists := destRepresentations[dest]

	if !isExists {
		return "", errors.New(fmt.Sprintf("parsing error. destination %s does not exists", dest))
	}

	return dest, nil
}

func (c *C) parseComputation(codeLine string) (string, error) {
	comp := ""
	dest, rest, hasDest := strings.Cut(codeLine, "=")

	if !hasDest {
		comp, _, _ = strings.Cut(dest, ";")
	} else {
		comp, _, _ = strings.Cut(rest, ";")
	}

	if comp == "" {
		return "", errors.New("parsing error. computation is required")
	}

	_, isExists := compRepresentations[comp]

	if !isExists {
		return "", errors.New(fmt.Sprintf("parsing error. computation %s does not exists", comp))
	}

	return comp, nil
}

func (c *C) parseJump(codeLine string) (string, error) {
	_, jump, _ := strings.Cut(codeLine, ";")

	if jump == "" {
		return "", nil
	}

	_, isExists := jumpRepresentations[jump]

	if !isExists {
		return "", errors.New(fmt.Sprintf("parsing error. jump %s does not exists", jump))
	}

	return jump, nil
}

func (c *C) representDest() string {
	if c.Dest == "" {
		return "000"
	}

	binary, _ := destRepresentations[c.Dest]

	return binary
}

func (c *C) representComp() string {
	binary, _ := compRepresentations[c.Comp]

	return binary
}

func (c *C) representJump() string {
	if c.Jump == "" {
		return "000"
	}

	binary, _ := jumpRepresentations[c.Comp]

	return binary
}
