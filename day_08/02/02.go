package main

import (
	"bufio"
	"fmt"
	"os"
)

type OpCode string

const (
	Nop OpCode = "nop"
	Acc        = "acc"
	Jmp        = "jmp"
)

type Instruction struct {
	operation OpCode
	argument  int
	processed bool
}

func loadInstructionsFromFile(path string) (instructions []Instruction, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var instruction Instruction
		_, err := fmt.Sscanf(scanner.Text(), "%s %d", &instruction.operation, &instruction.argument)
		if err != nil {
			return nil, err
		}

		instructions = append(instructions, instruction)
	}
	return instructions, scanner.Err()
}

type Registers struct {
	ip  int
	acc int
}

func runProgram(registers *Registers, program []Instruction) bool {
	for {
		if registers.ip == len(program) {
			return true
		}

		if registers.ip < 0 || registers.ip > len(program) {
			return false
		}

		inst := &program[registers.ip]

		if inst.processed {
			return false
		}

		ipDelta := 1

		switch inst.operation {
		case Nop:
			break
		case Acc:
			registers.acc += inst.argument
			break
		case Jmp:
			ipDelta = inst.argument
			break
		}

		inst.processed = true
		registers.ip += ipDelta
	}
}

func main() {
	// Load the file
	program, err := loadInstructionsFromFile("../input.txt")
	if err != nil {
		fmt.Print("File loading failed!", err)
		os.Exit(1)
	}

	for idx, inst := range program {

		modifiedProgram := make([]Instruction, len(program))
		copy(modifiedProgram, program)

		switch inst.operation {
		case Jmp:
			modifiedProgram[idx].operation = Nop
			break
		case Nop:
			modifiedProgram[idx].operation = Jmp
			break
		case Acc:
			continue
		}

		registers := Registers{}
		success := runProgram(&registers, modifiedProgram)

		if success {
			fmt.Printf("idx: %d IP: %d ACC: %d\n", idx, registers.ip, registers.acc)
			break
		}
	}
}
