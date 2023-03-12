package instruction

const max15bitValue = 36767

type Instruction interface {
	Assemble() string
}
