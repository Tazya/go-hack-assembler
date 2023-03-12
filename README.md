# Golang Assembler for [Nand To Tetris](https://www.nand2tetris.org/)

The assembler compile Hack assembly language files to CPU instructions in binary code.

"Build a Modern Computer from First Principles: From Nand to Tetris" is available on [Coursera](https://www.coursera.org/learn/build-a-computer)

## Usage

1. Install dependencies:
```
go mod vendor
```

2. Build:
```
go build -mod=vendor -o go-hack-assembler main.go
```

3. Use:

```
 ./go-hack-assembler -fileIn="./input/add.asm" -fileOut="./output/add" 
```

## Examples

Computes R0 = 2 + 3  (R0 refers to RAM[0])
```
@2      -> 0000000000000010
D=A     -> 1110100110000000
@3      -> 0000000000000011
D=D+A   -> 1110100000010000
@0      -> 0000000000000000
M=D     -> 1110010001100000
```



